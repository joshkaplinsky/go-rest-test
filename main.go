package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joshkaplinsky/go-rest-test/pkg/routes"
	"log"
	"os"
	"os/signal"
)

func main() {
	app := fiber.New()
	routes.PublicRoutes(app)

	StartServer(app)
}

func StartServer(a *fiber.App) {
	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := a.Shutdown(); err != nil {
			log.Printf("error shutting down server: %v", err)
		}

		close(idleConnsClosed)
	}()

	if err := a.Listen(GetServerHost()); err != nil {
		log.Printf("error running server: %v", err)
	}

	<-idleConnsClosed
}

func GetServerHost() string {
	return fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"))
}
