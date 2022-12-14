package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

func main() {
	r := gin.Default()
	m := melody.New()

	r.Use(static.Serve("/", static.LocalFile("./frontend", true)))

	r.GET("/chat", func(c *gin.Context) {
		err := m.HandleRequest(c.Writer, c.Request)
		if err != nil {
			return
		}
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		err := m.Broadcast(msg)
		if err != nil {
			return
		}
	})

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
