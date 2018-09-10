package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
	//"github.com/Depado/ginprom"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"github.com/valinurovam/safequeue"
)

// The base coffee machine has two different types of beans as a base.

type UserError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type BaseCallBack struct {
	UserRequestID    string     `json:"requestId"`
	Success          bool       `json:"success,omitempty"`
	Error            *UserError `json:",omitempty"`
	CallBackUrl      string     `json:"callbackurl,omitempty"`
	ShellCallBackUrl string     `json:"shellcallbackurl,omitempty"`
	Active           bool       `json:"jobactive,omitempty"`
	Running          bool       `json:"running,omitempty"`
}

type CupParms struct {
	StartBrewTime string `json:"startbrewtime,omitempty"`
	CupSize       int    `json:"CupSize,omitempty"`
	CupBean       int    `json:"CupBean,omitempty"`
	CupStrength   int    `json:"CupStrength,omitempty"`
}

// An Array of Cup Parameters!
type Cups []CupParms

type MessageTmp struct {
	BaseCallBack
	CupParms
}
type QueueRequestParms struct {
	BaseCallBack
	Cups
}
type QueueStatusResponse struct {
	BaseCallBack
	Cups
	CurrentlyBrewing string `json:"currentlybrewing"`
	FinishedBrewing  string `json:"finishedBrewing",omitempty`
	RemainingToBrew  string `json:"remainingtobew"`
}
type Token struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

const SIZE = 4096

var q = safequeue.NewSafeQueue(SIZE)
var lock sync.Mutex

func main() {
	r := SetupRouter()
	// Start server
	r.Run()
}
func SetupRouter() *gin.Engine {
	// Initialize router
	r := gin.Default()
	// Get the requestID from the HEADER
	r.Use(RequestId())

	/*
			   Turning off Prometheus because it is breaking coffeemaker_test.go

		// Turn on exporting of metrics via /metrics
				p := ginprom.New(
					ginprom.Engine(r),
					ginprom.Subsystem("gin"),
					ginprom.Path("/metrics"),
				)
				r.Use(p.Instrument())
	*/

	// Routes
	r.GET("/QueueStatus/:requestId", QueueStatus)
	r.POST("/BrewCup", BrewCup)
	r.POST("/QueueRequest", QueueRequest)
	r.POST("/QueuePause", QueuePause)
	r.POST("/QueueCancel", QueueCancel)
	r.POST("/QueueStart", QueueStart)
	return r
}

func BrewCup(c *gin.Context) {
	var msg MessageTmp
	msg.Success = true
	checkuserRequestID := c.MustGet("RequestId").(string)
	msg.UserRequestID = checkuserRequestID
	if err := c.ShouldBindJSON(&msg); err != nil {
		msg.Success = false
		msg.Error.Code = 1
		msg.Error.Message = "Please provide a JSON string describing the type of coffee you would like"
		c.JSON(http.StatusBadRequest, msg)
		return
	}

	// Add to queue
	q.Push(msg)
	c.JSON(http.StatusOK, msg)
}
func QueueStatus(c *gin.Context) {
	var msg QueueRequestParms
	id := c.Param("requestId")
	msg.UserRequestID = id
	queueLength := q.Length()
	for item := uint64(0); item < queueLength; item++ {
		msg := q.Pop()
		user := msg.(QueueRequestParms).UserRequestID
		if user == id {
			q.Push(msg)
			break
		}
		q.Push(msg)
	}

	msg.Success = true
	c.JSON(http.StatusOK, msg)
}
func QueueRequest(c *gin.Context) {
	var msg QueueRequestParms
	checkuserRequestID := c.MustGet("RequestId").(string)
	msg.UserRequestID = checkuserRequestID
	msg.Success = true
	if err := c.ShouldBindJSON(&msg); err != nil {
		msg.Success = false
		msg.Error.Code = 2
		msg.Error.Message = "Please provide a JSON string describing the type of coffee you would like"
		c.JSON(http.StatusBadRequest, msg)
		return
	}
	msg.Active = false
	msg.CallBackUrl = "localhost:8080/QueueStatus?requestId=" + checkuserRequestID
	msg.ShellCallBackUrl = "curl -s localhost:8080/QueueStatus/" + checkuserRequestID + " | jq -C ."
	// Add to queue
	q.Push(msg)
	fmt.Println("Something in the coffee queue need to start a job")
	c.JSON(http.StatusOK, msg)
}

func QueuePause(c *gin.Context) {
	var b QueueStatusResponse
	id := c.Param("requestId")
	b.UserRequestID = id
	b.Success = true
	b.Active = false
	c.IndentedJSON(http.StatusOK, b)
}
func QueueCancel(c *gin.Context) {
	var b QueueStatusResponse
	id := c.Param("requestId")
	b.UserRequestID = id
	b.Success = true
	b.Active = false
	c.IndentedJSON(http.StatusOK, b)
}
func QueueStart(c *gin.Context) {
	var b QueueStatusResponse
	id := c.Param("requestId")
	b.UserRequestID = id
	b.Success = true
	b.Active = true
	// Start making coffee
	go MakeCoffee()
	c.IndentedJSON(http.StatusOK, b)
}

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check for incoming header, use it if exists
		requestID := c.Request.Header.Get("X-Request-Id")

		// Create request id with UUID4
		if requestID == "" {
			uuid4, _ := uuid.NewV4()
			requestID = uuid4.String()
		}

		// Put in header
		c.Set("RequestId", requestID)

		// Set X-Request-Id header
		c.Writer.Header().Set("X-Request-Id", requestID)
		c.Next()
	}
}

//This does not start until after the first QueueStart Happens because before that there is NO coffee to make
func MakeCoffee() {
	lock.Lock()
	defer lock.Unlock()

	fmt.Println("Start making coffee")
	for {

		time.Sleep(time.Second * 200)
		fmt.Println("Still making coffee")
	}
}
