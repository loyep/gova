package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func bootstrap(router *gin.Engine) {
	router.LoadHTMLGlob("views/**/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home/index.html", gin.H{})
	})
	router.StaticFS("/static", http.Dir("./public/static"))

	port := ":8000"
	fmt.Printf("127.0.0.1%s\n", port)

	srv := &http.Server{
		Addr:    port,
		Handler: router,
	}
	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

}

func main() {
	router := gin.Default()
	bootstrap(router)

}
