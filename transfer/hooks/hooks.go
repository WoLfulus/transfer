package hooks

import (
	"fmt"
	"net/http"
	"os"

	"github.com/docker/distribution/notifications"
	"github.com/foomo/htpasswd"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func init() {
	setupPasswords()
	startServer()
}

func setupPasswords() {
	username := os.Getenv("REGISTRY_USERNAME")
	if username == "" {
		log.Fatal("Missing required environment variable: REGISTRY_USERNAME")
	}

	password := os.Getenv("REGISTRY_PASSWORD")
	if username == "" {
		log.Fatal("Missing required environment variable: REGISTRY_PASSWORD")
	}

	err := htpasswd.SetPassword("/data/auth/htpasswd", username, password, htpasswd.HashBCrypt)
	if err != nil {
		log.Fatal(err)
	}
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
	log.Info("image received %s", image)
}
