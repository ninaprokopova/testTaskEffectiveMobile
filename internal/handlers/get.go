package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"testEffMobile/internal/dto"
	database "testEffMobile/packages/database/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetPersonHandler godoc
// @Summary Get person by ID
// @Description Get person details by person ID
// @Tags persons
// @Accept  json
// @Produce  json
// @Param id path int true "Person ID"
// @Success 200 {object} dto.GetPersonResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/person/{id} [get]
func (h *PersonHandler) GetPersonHandler(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(
			dto.CodeBadRequest, "ID must be positive integer"))
		return
	}

	user := database.User{}
	result := h.DB.Model(&database.User{}).Where("id = ?", id).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, dto.NewErrorResponse(
			dto.CodeNotFound, "There is no user with this id"))
		return
	}
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse(
			dto.CodeInternalError, "Internal server error"))
		return
	}

	c.JSON(http.StatusOK, dto.GetPersonResponse{
		ID:         user.ID,
		Name:       user.Name,
		Surname:    user.Surname,
		Patronymic: user.Patronymic,
		Gender:     user.Gender,
	})
}

// GetPeopleHandler godoc
// @Summary Get people by query parameters
// @Description Get filtered and paginated list of people
// @Tags people
// @Accept json
// @Produce json
// @Param name 		  query string false "Filter by name"
// @Param surname     query string false "Filter by surname"
// @Param patronymic  query string false "Filter by patronymic"
// @Param age 		  query int    false "Filter by age"
// @Param gender 	  query string false "Filter by gender"
// @Param nationality query string false "Filter by nationality"
// @Param page        query int    false "Page number" default(1) minimum(1)
// @Param limit       query int    false "Items per page" default(10) minimum(1) maximum(100)
// @Param sort_by query string false "Field to sort by (name, surname, age, gender, nationality)"
// @Param sort_desc query boolean false "Sort in descending order" default(false)
// @Success 200 {object} dto.GetPeopleResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/people [get]
func (h *PersonHandler) GetPeopleHandler(c *gin.Context) {
	h.logger.Debug("Starting GetPeople Handler")
	var req dto.GetPeopleRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		h.logger.Warn("Invalid query params")
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(
			dto.CodeBadRequest, fmt.Sprintf("Invalid query params: %v", err)))
		return
	}

	query := h.DB.Model(&database.User{})

	if req.Name != "" {
		query = query.Where("name = ?", req.Name)
	}
	if req.Surname != "" {
		query = query.Where("surname = ?", req.Surname)
	}
	if req.Patronymic != "" {
		query = query.Where("patronymic = ?", req.Patronymic)
	}
	if req.Age > 0 {
		query = query.Where("age = ?", req.Age)
	}
	if req.Gender != "" {
		query = query.Where("gender = ?", req.Gender)
	}
	if req.Nationality != "" {
		query = query.Where("nation = ?", req.Nationality)
	}

	if req.SortBy != "" {
		order := req.SortBy
		if req.SortDesc {
			order += " DESC"
		}
		query = query.Order(order)
	}

	offset := (req.Page - 1) * req.Limit
	var total int64
	query.Count(&total)

	var users []database.User
	if err := query.Offset(offset).Limit(req.Limit).Find(&users).Error; err != nil {
		h.logger.Warn("Database request error", "request", query)
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse(
			dto.CodeInternalError, "database request error"))
		return
	}

	response := make([]dto.Person, 0, len(users))
	for i := range users {
		response = append(response, dto.Person{
			ID:         users[i].ID,
			Name:       users[i].Name,
			Surname:    users[i].Surname,
			Patronymic: users[i].Patronymic,
			Gender:     users[i].Gender,
			Nation:     users[i].Nation,
		})
	}

	h.logger.Info("GetPeopleHandler completed successfully")
	c.JSON(http.StatusOK, dto.GetPeopleResponse{
		People: response,
		MetaData: dto.Meta{
			Total: total,
			Page:  req.Page,
			Limit: req.Limit,
		},
	})
}
