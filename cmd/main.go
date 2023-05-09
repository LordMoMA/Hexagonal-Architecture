package main

import (
	"fmt"
	"log"
	"os"

	"github.com/LordMoMA/Hexagonal-Architecture/internal/adapters/cache"
	"github.com/LordMoMA/Hexagonal-Architecture/internal/adapters/handler"
	"github.com/LordMoMA/Hexagonal-Architecture/internal/adapters/repository"
	"github.com/LordMoMA/Hexagonal-Architecture/internal/core/services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	msgService     *services.MessengerService
	userService    *services.UserService
	paymentService *services.PaymentService
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

	store := repository.NewDB(db)
	redisCache, err := cache.NewRedisCache("localhost:6379", "")
	if err != nil {
		panic(err)
	}

	msgService = services.NewMessengerService(store)
	userService = services.NewUserService(store, redisCache)
	paymentService = services.NewPaymentService(store)

	InitRoutes()

}

func InitRoutes() {
	router := gin.Default()
	// router.Static("/", "internal/web/public")

	v1 := router.Group("/v1")

	messageHandler := handler.NewMessageHandler(*msgService)
	v1.GET("/messages/:id", messageHandler.ReadMessage)
	v1.GET("/messages", messageHandler.ReadMessages)
	v1.POST("/messages", messageHandler.CreateMessage)
	v1.PUT("/messages/:id", messageHandler.UpdateMessage)
	v1.DELETE("/messages/:id", messageHandler.DeleteMessage)

	userHandler := handler.NewUserHandler(*userService)
	v1.GET("/users/:id", userHandler.ReadUser)
	v1.GET("/users", userHandler.ReadUsers)
	v1.POST("/users", userHandler.CreateUser)
	v1.PUT("/users", userHandler.UpdateUser)
	v1.DELETE("/users", userHandler.DeleteUser)

	v1.POST("/login", userHandler.LoginUser)
	v1.POST("/membership/webhooks", userHandler.UpdateMembershipStatus)

	v2 := router.Group("/v2")
	paymentHandler := handler.NewPaymentHandler(*paymentService)
	v2.POST("/create-checkout-session", paymentHandler.CreateCheckoutSession)
	// v2.POST("/payments", paymentHandler.CreatePayment)

	go func() {
		if err := router.Run(":5000"); err != nil {
			log.Fatalf("failed to run messages and users service: %v", err)
		}
	}()

	if err := router.Run(":4242"); err != nil {
		log.Fatalf("failed to run payments service: %v", err)
	}

	// router.Run(":4242")
}
