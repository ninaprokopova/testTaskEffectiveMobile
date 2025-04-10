package dto

type CreatePersonRequest struct {
	Name       string `json:"name" binding:"required"`
	Surname    string `json:"surname" binding:"required"`
	Patronymic string `json:"patronymic,omitempty"`
}

// GetPeopleRequest represents query parameters for getting people
// @Description Parameters for filtering and pagination of people list
type GetPeopleRequest struct {
	Name        string `form:"name"`
	Surname     string `form:"surname"`
	Patronymic  string `form:"patronymic"`
	Age         int    `form:"age"`
	Gender      string `form:"gender"`
	Nationality string `form:"nationality"`
	Page        int    `form:"page,default=1" binding:"min=1"`
	Limit       int    `form:"limit,default=10" binding:"min=1,max=100"`
	SortBy      string `form:"sort_by"`
	SortDesc    bool   `form:"sort_desc"`
}

type UpdatePersonRequest struct {
	Id          int     `form:"id" binding:"required,min=1"`
	Name        *string `form:"name,omitempty"`
	Surname     *string `form:"surname,omitempty"`
	Patronymic  *string `form:"patronymic,omitempty"`
	Age         *int    `form:"age,omitempty" binding:"min=0"`
	Gender      *string `form:"gender,omitempty" validate:"oneof=male female"`
	Nationality *string `form:"nationality,omitempty"`
}
