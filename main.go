package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"os"
	"reneat-microservice-user/config"
	"reneat-microservice-user/database"
	timeHelper "reneat-microservice-user/helpers/time"
	"reneat-microservice-user/routes"
	"reneat-microservice-user/services/logService"
	"sync"
)

var (
	engine *gin.Engine
	cfg    *viper.Viper
)

func init() {
	engine = gin.New()
	engine.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: []string{"/health"},
	}))
	engine.Use(gin.Logger())

	logService.NewLogrus()
}

func main() {
	enviroment := flag.String("e", os.Getenv("APP_ENV"), "")

	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}

	flag.Parse()

	config.Init(*enviroment)
	cfg := config.GetConfig()

	_, err := database.Init()
	if err == nil {
		fmt.Println("\n Database connected!")
	} else {
		fmt.Println("Fatal error database connection", err)
	}

	ok := timeHelper.SetTimeZone(cfg.GetString("server.timezone"))
	if ok != nil {
		fmt.Println("Fatal error timezone", ok)
	}

	port := cfg.GetString("server.port")
	print("port ", port)

	var wg sync.WaitGroup
	wg.Add(1) // Adding one task to wait for - the server

	go func() {
		log.Println("Starting server...")
		defer wg.Done() // Mark this task as done when the func returns
		StartRest(port)
		log.Println("Server terminated.")
	}()

	wg.Wait() // Block here until wg.Done() is called
}

func StartRest(port string) {
	routes.RouteInit(engine)

	if err := engine.Run(":" + port); err != nil {
		log.Fatalln(err)
	}

	if err := engine.Run(":" + port); err != nil {
		log.Fatalf("Failed to start the server on port %s: %v", port, err)
	}
}
