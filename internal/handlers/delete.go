package handlers

import (
	"net/http"
	"strconv"
	"testEffMobile/internal/dto"
	database "testEffMobile/packages/database/models"

	"github.com/gin-gonic/gin"
)

// DeletePerson godoc
// @Summary Delete person by ID
// @Description Delete person from database by person ID
// @Tags persons
// @Param id path int integer "Person ID" minimum(1)
// @Success 204 "No Content"
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/person/{id} [delete]
func (h *PersonHandler) DeletePerson(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	h.logger.Info("Starting DeletePerson handler", "id", id)

	if err != nil || id <= 0 {
		h.logger.Warn("Invalid ID parmetr", "id", id)
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(
			dto.CodeBadRequest, "ID must be positive integer"))
		return
	}

	h.logger.Debug("Attempting to delete user from databse", "id", id)

	result := h.DB.Where("id = ?", id).Delete(&database.User{})

	if result.RowsAffected == 0 {
		h.logger.Warn("User not found", "id", id)
		c.JSON(http.StatusNotFound, dto.NewErrorResponse(
			dto.CodeNotFound, "There is no user with this id"))
		return
	}
	if result.Error != nil {
		h.logger.Error("Database error when deleting user", "id", id)
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse(
			dto.CodeInternalError, "Internal server error"))
		return
	}

	h.logger.Info("User successfully deleted", "id", id)
	c.Status(http.StatusNoContent)
}
