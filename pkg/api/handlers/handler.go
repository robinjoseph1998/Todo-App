package handlers

import (
	"Todo/pkg/api/models"
	use "Todo/pkg/api/usecase/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	usecase use.UseCaseInterface
}

func NewHandler(usecase use.UseCaseInterface) *Handler {

	return &Handler{
		usecase: usecase,
	}
}

func (hh *Handler) TestFunction(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{":::::": "WELCOME TO MY TODO APP"})
}

func (hh *Handler) AddName(c *gin.Context) {
	var request models.Test
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	SavedName, err := hh.usecase.ExecuteAddName(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Name Saved": SavedName})

}
