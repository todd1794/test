package main

import (
	"fmt"
	"github.com/Depado/ginprom"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"github.com/valinurovam/safequeue"
	"net/http"
)

// The base coffee machine has two different types of beans as a base.

type UserError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type BaseCallBack struct {
	UserRequestID string     `json:"userrequestId"`
	Success       bool       `json:"success"`
	Error         *UserError `json:",omitempty"`
}

type CupParms struct {
	StartBrewTime string `json:"startbrewtime,omitempty"`
	CupSize       int    `json:"CupSize,omitempty"`
	CupBean       int    `json:"CupBean,omitempty"`
	CupStrength   int    `json:"CupStrength,omitempty"`
}

type Cups []CupParms

type MessageTmp struct {
	BaseCallBack
	CupParms
}
type QueueRequestParms struct {
	BaseCallBack
	Cups
}

const SIZE = 4096

var q = safequeue.NewSafeQueue(SIZE)

func main() {
	// Initialize router
	r := gin.Default()
	// Get the requestID from the HEADER
	r.Use(RequestId())
	// Turn exporting of metrics via /metrics
	p := ginprom.New(
		ginprom.Engine(r),
		ginprom.Subsystem("gin"),
		ginprom.Path("/metrics"),
	)
	r.Use(p.Instrument())

	// Routes
	r.GET("/QueueStatus", QueueStatus)
	r.POST("/BrewCup", BrewCup)
	r.POST("/QueueRequest", QueueRequest)
	r.POST("/QueuePause", QueuePause)
	r.POST("/QueueCancel", QueueCancel)
	r.POST("/QueueStart", QueueStart)

	// Start server
	r.Run()
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
	var msg MessageTmp
	queueLength := q.Length()
	for item := uint64(0); item < queueLength; item++ {
		pop := q.Pop()
		fmt.Printf("The value of pop is:=%+v\n", pop)
		q.Push(pop)
	}
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
	// Add to queue
	q.Push(msg)

	c.JSON(http.StatusOK, msg)
}

func QueuePause(c *gin.Context) {
	var msg MessageTmp
	checkuserRequestID := c.MustGet("RequestId").(string)
	msg.UserRequestID = checkuserRequestID
	msg.Success = true
	c.IndentedJSON(http.StatusOK, msg)
}
func QueueCancel(c *gin.Context) {
	var msg MessageTmp
	checkuserRequestID := c.MustGet("RequestId").(string)
	msg.UserRequestID = checkuserRequestID
	msg.Success = true
	c.IndentedJSON(http.StatusOK, msg)
}
func QueueStart(c *gin.Context) {
	var msg MessageTmp
	checkuserRequestID := c.MustGet("RequestId").(string)
	msg.UserRequestID = checkuserRequestID
	msg.Success = true
	c.IndentedJSON(http.StatusOK, msg)
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

		// Expose it for use in the application
		// email ok, move on
		c.Set("RequestId", requestID)

		// Set X-Request-Id header
		c.Writer.Header().Set("X-Request-Id", requestID)
		c.Next()
	}
}
