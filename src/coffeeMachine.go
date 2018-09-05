package main

import (
	"fmt"
	"github.com/Depado/ginprom"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

// The base coffee machine has two different types of beans as a base.

type BaseCallBack struct {
	UserRequestID string `json:"userrequestId"`
	Success       bool   `json:"success"`
	Error         struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

type CupParms struct {
	CupSize     int `validate:"required,max=3"`
	CupBean     int `validate:"required,min=1,max=2"`
	CupStrength int `validate:"required,min=1,max=5"`
}

var validate *validator.Validate

func main() {

	// Initialize router
	r := gin.Default()
	// Turn exporting of metrics via /metrics
	r.Use(RequestId())
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
	var msg BaseCallBack
	var cups CupParms
	validate = validator.New()
	validateCupParms()
	//fmt.Println(cups)
	checkuserRequestID := c.MustGet("RequestId").(string)
	msg.UserRequestID = checkuserRequestID
	msg.Success = true
	c.IndentedJSON(http.StatusOK, msg)
}
func QueueRequest(c *gin.Context) {
	var msg BaseCallBack
	var cups CupParms
	validate = validator.New()
	validateCupParms()
	fmt.Println(cups)
	checkuserRequestID := c.MustGet("RequestId").(string)
	msg.UserRequestID = checkuserRequestID
	msg.Success = true
	c.IndentedJSON(http.StatusOK, msg)
}
func QueueStatus(c *gin.Context) {
	var msg BaseCallBack
	var cups CupParms
	validate = validator.New()
	validateCupParms()
	fmt.Println(cups)
	checkuserRequestID := c.MustGet("RequestId").(string)
	msg.UserRequestID = checkuserRequestID
	msg.Success = true
	c.IndentedJSON(http.StatusOK, msg)
}
func QueuePause(c *gin.Context) {
	var msg BaseCallBack
	var cups CupParms
	validate = validator.New()
	validateCupParms()
	fmt.Println(cups)
	checkuserRequestID := c.MustGet("RequestId").(string)
	msg.UserRequestID = checkuserRequestID
	msg.Success = true
	c.IndentedJSON(http.StatusOK, msg)
}
func QueueCancel(c *gin.Context) {
	var msg BaseCallBack
	var cups CupParms
	validate = validator.New()
	validateCupParms()
	fmt.Println(cups)
	checkuserRequestID := c.MustGet("RequestId").(string)
	msg.UserRequestID = checkuserRequestID
	msg.Success = true
	c.IndentedJSON(http.StatusOK, msg)
}
func QueueStart(c *gin.Context) {
	var msg BaseCallBack
	var cups CupParms
	validate = validator.New()
	validateCupParms()
	fmt.Println(cups)
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

func validateCupParms() {

	cupparms := &CupParms{
		CupSize:     1,
		CupBean:     1,
		CupStrength: 4,
	}

	errs := validate.Var(cupparms, "required")

	if errs != nil {
		fmt.Println(errs)
		return
	}

}
