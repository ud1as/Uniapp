package model

type CreateProfessorRequest struct {
	Name    string `json:"name" db:"name" example:"Ben"`
	Surname string `json:"surname" db:"surname" example:"Tyler"`
	Email   string `json:"email" db:"email" example:"ben.tyler@uni.edu.kz"`
	Degree  string `json:"degree" db:"degree" example:"Masters in Computer Science"`
}

type CreateProfessorResponse struct {
	Id int `json:"id,omitempty" db:"id" example:"1" swaggerignore:"true"`
}
