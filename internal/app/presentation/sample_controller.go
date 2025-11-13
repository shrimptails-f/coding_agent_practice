package presentation

import (
	"business/internal/sample/application"
	"business/internal/sample/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SampleController exposes HTTP handlers that talk to the sample use case.
type SampleController struct {
	useCase application.UseCase
}

func NewSampleController(useCase application.UseCase) *SampleController {
	return &SampleController{useCase: useCase}
}

// ListSamples returns every stored sample. It intentionally keeps the handler tiny.
func (c *SampleController) ListSamples(ctx *gin.Context) {
	samples, err := c.useCase.ListSamples()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": samples})
}

// CreateSample persists a new sample entry.
func (c *SampleController) CreateSample(ctx *gin.Context) {
	var req createSampleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	_, err := c.useCase.CreateSample(req.toInput())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusCreated)
}

type createSampleRequest struct {
	Title string `json:"title"`
}

func (r createSampleRequest) toInput() domain.CreateSampleInput {
	return domain.CreateSampleInput{Title: r.Title}
}
