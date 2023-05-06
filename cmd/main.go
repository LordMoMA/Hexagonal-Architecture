package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/LordMoMA/Hexagonal-Architecture/internal/adapters/handler"
	"github.com/LordMoMA/Hexagonal-Architecture/internal/core/services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var (
   repo        = flag.String("db", "postgres", "Database for storing messages")
   // httpHandler *handler.HTTPHandler
   svc         *services.MessengerService
)

func main() {
   err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open("postgres", conn)
	if err != nil {
		panic(err)
	}

   defer db.Close()

   // newdb := repository.NewDB(db)


   // flag.Parse()

   // fmt.Printf("Application running using %s\n", *repo)
   // switch *repo {
   // case "redis":
   //     store := repository.NewMessengerRedisRepository()
   //     svc = services.NewMessengerService(store)
   // default:
   //     store := repository.NewMessengerPostgresRepository()
   //     svc = services.NewMessengerService(store)
   // }

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