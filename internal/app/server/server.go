package server

import (
	v1 "business/internal/app/router"
	"business/internal/di"
	"business/internal/library/mysql"
	"log"

	"github.com/gin-gonic/gin"
)

func Run() {
	g := gin.Default()

	conn, err := mysql.New()
	if err != nil {
		log.Fatalf("failed to initialize mysql: %v", err)
	}

	container := di.BuildContainer(conn)

	router := v1.NewRouter(g, container)
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}
}
