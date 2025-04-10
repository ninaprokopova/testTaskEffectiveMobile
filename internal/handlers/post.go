package handlers

import (
	"net/http"
	"testEffMobile/internal/dto"
	database "testEffMobile/packages/database/models"

	"github.com/gin-gonic/gin"
)

// CreatePerson godoc
// @Summary Create new person
// @Description Create new person with provided data
// @Tags persons
// @Accept  json
// @Produce  json
// @Param input body dto.CreatePersonRequest true "Person data"
// @Success 201 {object} dto.CreatePersonResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/person [post]
func (h *PersonHandler) CreatePerson(c *gin.Context) {

	h.logger.Debug("Start Handler CreatePerson")
	var input dto.CreatePersonRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		h.logger.Warn("Invalid JSON format", "error", err)
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(
			dto.CodeBadRequest,
			"Invalid JSON format",
		))
		return
	}

	h.logger.Info("Creating new person", "name", input.Name, "surname", input.Surname)

	gender, _ := h.enricher.GetGender(c, input.Name)
	age, _ := h.enricher.GetAge(c, input.Name)
	nationality, _ := h.enricher.GetNationality(c, input.Name)

	user := database.User{
		Name:       input.Name,
		Surname:    input.Surname,
		Patronymic: input.Patronymic,
		Gender:     gender,
		Age:        age,
		Nation:     nationality,
	}

	result := h.DB.Model(&database.User{}).Create(&user)

	if result.Error != nil {
		h.logger.Error("Database error", "error", result.Error)
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse(
			dto.CodeInternalError,
			"Failed to create person, database problem",
		))
		return
	}

	h.logger.Info("Person created successfully", "id", user.ID)
	c.JSON(http.StatusCreated, dto.CreatePersonResponse{
		ID: user.ID,
	})
}
