package v1

import (
	"business/internal/app/presentation"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func NewRouter(g *gin.Engine, container *dig.Container) *gin.Engine {
	var sampleController *presentation.SampleController
	if err := container.Invoke(func(controller *presentation.SampleController) {
		sampleController = controller
	}); err != nil {
		log.Fatalf("failed to resolve SampleController: %v", err)
	}

	g.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Sample Clean Architecture stack is running",
		})
	})

	v1 := g.Group("/v1")

	v1.GET("/samples", sampleController.ListSamples)
	v1.POST("/samples", sampleController.CreateSample)

	return g
}
