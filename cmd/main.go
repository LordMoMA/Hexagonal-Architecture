package main

import (
	"flag"
	"fmt"

	"github.com/LordMoMA/Hexagonal-Architecture/internal/adapters/handler"
	"github.com/LordMoMA/Hexagonal-Architecture/internal/adapters/repository"
	"github.com/LordMoMA/Hexagonal-Architecture/internal/core/services"
	"github.com/gin-gonic/gin"
)

var (
   repo        = flag.String("db", "postgres", "Database for storing messages")
   // httpHandler *handler.HTTPHandler
   svc         *services.MessengerService
)

func main() {
   flag.Parse()

   fmt.Printf("Application running using %s\n", *repo)
   switch *repo {
   case "redis":
      //  store := repository.NewMessengerRedisRepository()
      //  svc = services.NewMessengerService(store)
   default:
       store := repository.NewMessengerPostgresRepository()
       svc = services.NewMessengerService(store)
   }

   InitRoutes()

}

func InitRoutes() {
   router := gin.Default()
   handler := handler.NewHTTPHandler(*svc)
   router.GET("/messages/:id", handler.ReadMessage)
   router.GET("/messages", handler.ReadMessages)
   router.POST("/messages", handler.CreateMessage)
   router.PUT("/messages/:id", handler.UpdateMessage)
   router.DELETE("/messages/:id", handler.DeleteMessage)
   router.Run(":5000")
}