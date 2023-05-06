package main

import (
	"fmt"
	"os"

	"github.com/LordMoMA/Hexagonal-Architecture/internal/adapters/handler"
	"github.com/LordMoMA/Hexagonal-Architecture/internal/adapters/repository"
	"github.com/LordMoMA/Hexagonal-Architecture/internal/core/services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
   // repo        = flag.String("db", "postgres", "Database for storing messages")
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

   // defer db.Close()

   store := repository.NewDB(db)
   svc = services.NewMessengerService(store)

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
   v1 := router.Group("/v1")
   handler := handler.NewMessageHandler(*svc)
   v1.GET("/messages/:id", handler.ReadMessage)
   v1.GET("/messages", handler.ReadMessages)
   v1.POST("/messages", handler.CreateMessage)
   v1.PUT("/messages/:id", handler.UpdateMessage)
   v1.DELETE("/messages/:id", handler.DeleteMessage)
   router.Run(":5000")
}