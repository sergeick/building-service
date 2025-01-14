package handlers

import (
	"building-service/internal/models"
	"building-service/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BuildingHandler struct {
	repo *repository.BuildingRepository
}

func NewBuildingHandler(repo *repository.BuildingRepository) *BuildingHandler {
	return &BuildingHandler{repo: repo}
}

// @Summary Create a new building
// @Description Create a new building
// @Tags buildings
// @Accept json
// @Produce json
// @Param building body models.Building true "Building object"
// @Success 201 {object} models.Building
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /buildings [post]
func (h *BuildingHandler) Create(c *gin.Context) {
	var building models.Building
	if err := c.ShouldBindJSON(&building); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.repo.Create(&building); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, building)
}

// @Summary List buildings
// @Description Get list of buildings
// @Tags buildings
// @Accept json
// @Produce json
// @Param city query string false "City"
// @Param year query int false "Year built"
// @Param floors query int false "Floors count"
// @Success 200 {array} models.Building
// @Failure 500 {object} models.ErrorResponse
// @Router /buildings [get]
func (h *BuildingHandler) List(c *gin.Context) {
	city := c.Query("city")
	year := c.DefaultQuery("year", "")
	floors := c.DefaultQuery("floors", "")
	buildings, err := h.repo.List(city, year, floors)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, buildings)
}
