package hooks

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/docker/distribution/notifications"
)

// StartServer starts the HTTP notification server
func StartServer() {
	router := gin.Default()

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

		c.JSON(http.StatusOK, gin.H{
			"status": "received",
		})
	})

	router.Run(":5555")
}

/*
func handleRequest(w http.ResponseWriter, r *http.Request) {
    var ev notifications.Envelope

    bs, err := ioutil.ReadAll(r.Body)
    if err != nil {
        return
    }

    err = json.Unmarshal(bs, &ev)
    if err != nil {
        return
    }

    e := ev.Events[0]
    if e.Action == "push" && e.Target.Tag != "" {
        fmt.Printf("Pushed %s:%s\n", e.Target.Repository, e.Target.Tag)
        // OnPush()
    }
}
*/
