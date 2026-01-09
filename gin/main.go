package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	handler "github.com/unifuu/ditto2/gin/api"
	mongo "github.com/unifuu/ditto2/gin/db/mongo"
	"github.com/unifuu/ditto2/gin/db/redis"
	"github.com/unifuu/ditto2/gin/seed"
)

func main() {
	// Load .env file
	_ = godotenv.Load()
	log.Println(".env file loaded (or skipped if not found)")

	// Databases
	redis.Cli = redis.NewRedisClient()
	mongo.Init()

	// Seed database with initial data
	seed.SeedData()

	// Router
	router := gin.Default()
	router.Use(cors.Default())
	handler.Init(router)
	router.Run(":8080")
}
