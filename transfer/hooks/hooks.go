package hooks

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

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
	htpasswdPath := os.Getenv("HTPASSWD_PATH")
	if htpasswdPath == "" {
		htpasswdPath = "/data/auth/htpasswd"
	}

	username := os.Getenv("MANAGEMENT_USERNAME")
	if username == "" {
		username = "transfer"
	}

	password := os.Getenv("MANAGEMENT_PASSWORD")
	if password == "" {
		log.Errorln("MANAGEMENT_PASSWORD is not set. A random password will be used.")
		src := rand.NewSource(time.Now().UnixNano())
		password = string(1000000000 + src.Int63()%8000000000)
	}

	err := htpasswd.SetPassword(htpasswdPath, username, password, htpasswd.HashBCrypt)
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
