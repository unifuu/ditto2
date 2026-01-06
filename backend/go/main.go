package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	handler "github.com/unifuu/ditto2/backend/go/api"
	mongo "github.com/unifuu/ditto2/backend/go/db/mongo"
	"github.com/unifuu/ditto2/backend/go/db/redis"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("No .env file found")
	}

	// Databases
	redis.Cli = redis.NewRedisClient()
	mongo.Init()

	// Router
	router := gin.Default()
	router.Use(cors.Default())
	handler.Init(router)
	router.Run(":8080")
}
