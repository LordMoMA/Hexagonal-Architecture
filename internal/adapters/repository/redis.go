package repository

import (
	"encoding/json"
	"os"

	"github.com/LordMoMA/Hexagonal-Architecture/internal/core/domain"
	"github.com/go-redis/redis/v7"
	"github.com/joho/godotenv"
)

type MessengerRedisRepository struct {
   client *redis.Client
}

func NewMessengerRedisRepository() *MessengerRedisRepository {
    err := godotenv.Load()
    if err != nil {
        panic(err)
    }
    addr := os.Getenv("REDIS_ADDR")

    client := redis.NewClient(&redis.Options{
        Addr:     addr,
        Password: "",
        DB:       0,
    })


   return &MessengerRedisRepository{
       client: client,
   }
}

func (r *MessengerRedisRepository) CreateMessage(message domain.Message) error {
   json, err := json.Marshal(message)
   if err != nil {
       return err
   }
   r.client.HSet("messages", message.ID, json)
   return nil
}

func (r *MessengerRedisRepository) ReadMessage(id string) (*domain.Message, error) {
   value, err := r.client.HGet("messages", id).Result()
   if err != nil {
       return nil, err
   }
   message := &domain.Message{}
   err = json.Unmarshal([]byte(value), message)
   if err != nil {
       return nil, err
   }
   return message, nil
}

func (r *MessengerRedisRepository) ReadMessages() ([]*domain.Message, error) {
   messages := []*domain.Message{}
   value, err := r.client.HGetAll("messages").Result()
   if err != nil {
       return nil, err
   }

   for _, val := range value {
       message := &domain.Message{}
       err = json.Unmarshal([]byte(val), message)
       if err != nil {
           return nil, err
       }
       messages = append(messages, message)
   }

   return messages, nil

}

// update message
func (r *MessengerRedisRepository) UpdateMessage(id string, message domain.Message) error {
    value, err := r.client.HGet("messages", id).Result()
    if err != nil {
        return err
    }
    message.ID = id
    json, err := json.Marshal(message)
    if err != nil {
        return err
    }
    r.client.HSet("messages", id, json)
    return nil
}

// delete message
func (r *MessengerRedisRepository) DeleteMessage(id string) error {

}
