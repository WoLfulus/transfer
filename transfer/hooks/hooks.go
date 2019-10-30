package hooks

import (
	"fmt"
	"net/http"

	"github.com/docker/distribution/notifications"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func init() {
	startServer()
}

func startServer() {
	router := gin.New()

	router.Use(gin.Recovery())

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "received",
		})
	})

	router.POST("/", func(c *gin.Context) {
		var e notifications.Envelope

		if err := c.ShouldBindJSON(&e); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"error":  err.Error(),
			})
			return
		}

		for i := 0; i < len(e.Events); i++ {
			ev := e.Events[i]
			if ev.Action == "push" && ev.Target.Tag != "" {
				handlePushEvent(fmt.Sprintf("%s:%s", ev.Target.Repository, ev.Target.Tag))
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "received",
		})
	})

	go router.Run(":5555")
}

// OnPush processes push notifications
func handlePushEvent(image string) {
	log.Info("image received ", image)
}
