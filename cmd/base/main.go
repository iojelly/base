package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iojelly/base/pkg/logger"
	"net/http"
	"time"
)

func main() {
	router := gin.New()

	router.Use(logger.JSONLogger())

	router.Use(gin.Recovery())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
