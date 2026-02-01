package main

import (
	"context"
	"fmt"
	"log"

	"github.com/emel-study/emel-study-server/modules/server_module"
	"github.com/emel-study/emel-study-server/modules/study_module"
	"github.com/emel-study/emel-study-server/pkg"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	log.Printf("Starting emel-study-server with version: %s", pkg.Version)
	config, err := pkg.LoadConfig("conf/local.yml")
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	app := fiber.New()
	app.Use(cors.New())

	ctx := context.Background()
	if _, err := server_module.NewServerModule(ctx, app); err != nil {
		log.Fatalf("server module: %v", err)
	}
	if _, err := study_module.NewStudyModule(ctx, app); err != nil {
		log.Fatalf("study module: %v", err)
	}

	address := fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port)
	if err := app.Listen(address); err != nil {
		log.Fatalf("listen: %v", err)
	}
}
