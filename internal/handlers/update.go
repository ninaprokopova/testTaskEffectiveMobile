package handlers

import (
	"fmt"
	"net/http"
	"testEffMobile/internal/dto"
	database "testEffMobile/packages/database/models"

	"github.com/gin-gonic/gin"
)

// UpdatePerson godoc
// @Summary Update person by ID
// @Description Update person with provided data
// @Tags persons
// @Accept  json
// @Produce  json
// @Param input body dto.UpdatePersonRequest true "Person data"
// @Success 200 {object} dto.UpdatePersonResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/person [patch]
func (h *PersonHandler) UpdatePerson(c *gin.Context) {

	h.logger.Debug("Start UpdatePerson Handler")
	var input dto.UpdatePersonRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		h.logger.Warn("Invalid JSON format")
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(
			dto.CodeBadRequest,
			"Invalid JSON format",
		))
		return
	}

	if input.Id <= 0 {
		h.logger.Warn("Invalid ID. ID must be positive", "id", input.Id)
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(
			dto.CodeBadRequest,
			"Invalid ID. ID must be positive",
		))
		return
	}

	updates := database.User{}

	if input.Name != nil {
		updates.Name = *input.Name
	}

	if input.Surname != nil {
		updates.Surname = *input.Surname
	}

	if input.Patronymic != nil {
		updates.Patronymic = *input.Patronymic
	}

	if input.Age != nil {
		if *input.Age < 0 {
			h.logger.Warn("Invalid Age. Age must be positive", "Age", *input.Age)
			c.JSON(http.StatusBadRequest, dto.NewErrorResponse(
				dto.CodeBadRequest,
				"Age must be >= 0",
			))
			return
		}
		updates.Age = *input.Age
	}

	if input.Gender != nil {
		if *input.Gender != "male" && *input.Gender != "female" {
			h.logger.Warn("Invalid Gender. Gender must be male or female", "Gender", *input.Age)
			c.JSON(http.StatusBadRequest, dto.NewErrorResponse(
				dto.CodeBadRequest,
				"Gender must be male or female",
			))
			return
		}
		updates.Gender = *input.Gender
	}

	if input.Nationality != nil {
		updates.Nation = *input.Nationality
	}

	h.logger.Debug("Starting update user data in database", "id", input.Id)

	result := h.DB.Model(&database.User{}).Where("id = ?", input.Id).Updates(updates)

	if result.Error != nil {
		h.logger.Warn("Failed to update person, database problem")
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse(
			dto.CodeInternalError,
			"Failed to update person, database problem",
		))
		return
	}

	if result.RowsAffected == 0 {
		h.logger.Warn("Person doesn't found in database", "id", input.Id)
		c.JSON(http.StatusNotFound, dto.NewErrorResponse(
			dto.CodeNotFound,
			fmt.Sprintf("Person with ID = %v not found", input.Id),
		))
		return
	}

	h.logger.Info("Successful update of user data", "id", input.Id)
	c.JSON(http.StatusOK, dto.UpdatePersonResponse{
		ID:      uint(input.Id),
		Code:    dto.CodeStatusOK,
		Message: "Successful update of user data",
	})
}
