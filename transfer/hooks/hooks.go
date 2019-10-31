package hooks

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/docker/distribution/notifications"
	"github.com/foomo/htpasswd"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/wolfulus/transfer/transfer/tags"
)

var (
	managementUsername string = "transfer"
	managementPassword string = "transfer"
)

func init() {
	setupPasswords()
	startServer()
}

func setupPasswords() {
	htpasswdPath := os.Getenv("TRANSFER_HTPASSWD_PATH")
	if htpasswdPath == "" {
		htpasswdPath = "/data/auth/htpasswd"
	}

	managementUsername = os.Getenv("TRANSFER_MANAGEMENT_USERNAME")
	if managementUsername == "" {
		managementUsername = "transfer"
	}

	managementPassword = os.Getenv("TRANSFER_MANAGEMENT_PASSWORD")
	if managementPassword == "" {
		log.Errorln("TRANSFER_MANAGEMENT_PASSWORD is not set. A random password will be used.")
		src := rand.NewSource(time.Now().UnixNano())
		managementPassword = strconv.Itoa(1000000000 + int(src.Int63())%8000000000)
	}

	err := htpasswd.SetPassword(htpasswdPath, managementUsername, managementPassword, htpasswd.HashBCrypt)
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
				handlePushEvent(ev.Target.Repository, ev.Target.Tag)
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "received",
		})
	})

	go router.Run(":5555")
}

// OnPush processes push notifications
func handlePushEvent(image string, tag string) {
	portString := os.Getenv("TRANSFER_SERVICE_PORT")
	if portString == "" {
		portString = "5000"
	}

	port, err := strconv.Atoi(portString)
	if err != nil {
		log.Errorln("failed to convert registry port from env: TRANSFER_SERVICE_PORT=", port)
		return
	}

	registry := fmt.Sprintf("localhost:%d", port)
	registryImage := fmt.Sprintf("%s/%s:%s", registry, image, tag)
	localImage, err := tags.Decode(tag)
	if err != nil {
		log.Errorln("Failed to decode tag: ", tag)
		return
	}

	log.Infof("received image [%s] which resolves to [%s]\n", registryImage, localImage)

	cmd := exec.Command(
		"/bin/transfer", "restore",
		fmt.Sprintf("--registry=%s", registry),
		fmt.Sprintf("--username=%s", managementUsername),
		fmt.Sprintf("--password=%s", managementPassword),
		registryImage,
		localImage,
	)

	cmd.Stderr = log.StandardLogger().Writer()
	cmd.Stdout = log.StandardLogger().Writer()

	err = cmd.Run()
	if err != nil {
		log.Error("transfer command failed", err)
	}
}
