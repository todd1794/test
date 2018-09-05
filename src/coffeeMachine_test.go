Generated Test_main
Generated TestBrewCup
Generated TestQueueRequest
Generated TestQueueStatus
Generated TestQueuePause
Generated TestQueueCancel
Generated TestQueueStart
Generated TestRequestId
Generated Test_validateCupParms
package main

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func TestBrewCup(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BrewCup(tt.args.c)
		})
	}
}

func TestQueueRequest(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QueueRequest(tt.args.c)
		})
	}
}

func TestQueueStatus(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QueueStatus(tt.args.c)
		})
	}
}

func TestQueuePause(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QueuePause(tt.args.c)
		})
	}
}

func TestQueueCancel(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QueueCancel(tt.args.c)
		})
	}
}

func TestQueueStart(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QueueStart(tt.args.c)
		})
	}
}

func TestRequestId(t *testing.T) {
	tests := []struct {
		name string
		want gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RequestId(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RequestId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateCupParms(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for range tests {
		t.Run(tt.name, func(t *testing.T) {
			validateCupParms()
		})
	}
}
