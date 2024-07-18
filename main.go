package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	rest "github.com/scblur869/neo4j-api/rest"
	handler "github.com/scblur869/neo4j-api/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	key := handler.SetEncryptionKeyEnv()
	handler.CreateKeyFile("key.txt", key)
	os.Setenv("AESKEY", key)
	fmt.Println("key :", os.Getenv("AESKEY"))
}

func main() {

	//lets create some tables if they DONT EXIST!!!
	// allowedHost := os.Getenv("ALLOWED")
	appAddr := "0.0.0.0:" + os.Getenv("PORT")

	// setting up the handler for reciever functions
	neo := new(handler.NeoHandler)
	neo.Ctx = context.Background()
	neo.Config = handler.ParseConfiguration()
	driver, err := neo.Config.NewDriver()
	if err != nil {
		log.Fatal(err)
	}
	neo.Driver = driver
	// gin.SetMode(gin.ReleaseMode)
	// routes
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost"},
		AllowMethods:     []string{"POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Access-Control-Allow-Headers", "Access-Control-Allow-Origin", "Origin", "Accept", "X-Requested-With", "Content-Type", "Authorization", "Access-Control-Request-Method", "Access-Control-Request-Headers"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.GET("/api/v1/health", rest.HealthCheck)            // tested
	r.POST("/api/db/purge", neo.PurgeDbData)             // tested
	r.POST("/api/v1/addNode", neo.CreateNode)            // tested
	r.POST("/api/v1/addRelation", neo.CreatRelationship) // tested
	r.POST("/api/v1/deleteNode", neo.DeleteNode)         // tested
	r.POST("/api/v1/deleteRelation", neo.DeleteRelation) // tested
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Not Found"})
	})
	//server config

	srv := &http.Server{
		Addr:    appAddr,
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	//Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
