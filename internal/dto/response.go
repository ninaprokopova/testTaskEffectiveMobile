package dto

const (
	CodeStatusOK = "STATUS_CODE_OK"
)

type GetPersonResponse struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic,omitempty"`
	Age        int    `json:"age"`
	Gender     string `json:"gender"`
	Nation     string `json:"nation"`
}

type CreatePersonResponse struct {
	ID uint `json:"id"`
}

type UpdatePersonResponse struct {
	ID      uint   `json:"id"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type GetPeopleResponse struct {
	People   []Person `json:"people"`
	MetaData Meta     `json:"meta"`
}

type Person struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Age        int    `json:"age"`
	Gender     string `json:"gender"`
	Nation     string `json:"nation"`
}

type Meta struct {
	Total int64 `json:"total"`
	Page  int   `json:"page"`
	Limit int   `json:"limit"`
}
