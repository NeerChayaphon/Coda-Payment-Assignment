package game

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Ex: {"game":"Mobile Legends", "gamerID":"GYUTDTE", "points":20}
type RequestBody struct {
	Game    string `json:"game" binding:"required"`
	GamerID string `json:"gamerID" binding:"required"`
	Points  int    `json:"points" binding:"required"`
}

type gameInterface interface {
	Topup(c *gin.Context)
}

type gameHandler struct{}

func NewGameHandler() gameInterface {
	return &gameHandler{}
}

func getSlowResponse() (time.Duration, error) {
	slowResponseStr := os.Getenv("slow")
	if slowResponseStr != "" {
		slowResponse, err := strconv.Atoi(slowResponseStr)
		if err != nil {
			return 0, err
		} else {
			return time.Duration(slowResponse) * time.Second, nil
		}
	}

	return 0, nil
}

func (h *gameHandler) Topup(c *gin.Context) {

	// Simulate slow response
	slowResponse, err := getSlowResponse()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if slowResponse > 0 {
		time.Sleep(slowResponse)
	}

	var requestBody RequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, requestBody)
}
