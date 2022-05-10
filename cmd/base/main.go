package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iojelly/base/pkg/configs"
	"net/http"
	"time"
)

func main() {
	configs.Conf2Values()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	s := &http.Server{
		Addr:           ":" + configs.AppSetting.Port,
		Handler:        router,
		ReadTimeout:    configs.AppSetting.ReadTimeout * time.Second,
		WriteTimeout:   configs.AppSetting.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
