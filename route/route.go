package route

import (
	"context"
	"datingAppBE/controllers"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router struct {
	engine *gin.Engine
}

func GetRouter() Router {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:4200"}

	engine := gin.Default()
	engine.Use(cors.New(corsConfig))

	router := Router{
		engine: engine,
	}

	return router
}

func (router *Router) RegisterGET(relativePath string, handlers ...gin.HandlerFunc) {
	router.engine.GET(relativePath, handlers...)
}

func (router *Router) RegisterGETWithAuth(relativePath string, handlers ...gin.HandlerFunc) {
	router.engine.GET(relativePath, controllers.AuthCheck(handlers...))
}

func (router *Router) RegisterPOST(relativePath string, handlers ...gin.HandlerFunc) {
	router.engine.POST(relativePath, controllers.AuthCheck(handlers...))
}

func (router *Router) RegisterPUT(relativePath string, handlers ...gin.HandlerFunc) {
	router.engine.PUT(relativePath, controllers.AuthCheck(handlers...))
}

func (router *Router) RegisterPATCH(relativePath string, handlers ...gin.HandlerFunc) {
	router.engine.PATCH(relativePath, controllers.AuthCheck(handlers...))
}

func (router *Router) RegisterDELETE(relativePath string, handlers ...gin.HandlerFunc) {
	router.engine.DELETE(relativePath, controllers.AuthCheck(handlers...))
}

func (router *Router) RegisterController(controller Controller) {
	controller.Register(router)
}

func (router *Router) Start(addr string) {
	srv := &http.Server{
		Addr:    addr,
		Handler: router.engine,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 32)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
}
