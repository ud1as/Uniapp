package model

// Student Group Scheduler

// swagger:model Student
type Student struct {
	Id      int    `json:"id,omitempty" db:"id" example:"1" swaggerignore:"true"`
	Name    string `json:"name" db:"name" example:"Madina"`
	Surname string `json:"surname" db:"surname" example:"Akhmetova"`
	Gender  string `json:"gender" db:"gender" example:"F"`
}

type Group struct {
	Id        int    `json:"id" db:"id" example:"id"`
	Groupname string `json:"groupname" db:"group_name" example:"group_name"`
}

type Schedule struct {
	Id        int    `json:"id" db:"id" example:"id"`
	GroupId   int    `json:"groupId" db:"group_id" example:"group_id"`
	Professor string `json:"professor" db:"professor" example:"professor"`
	Room      int    `json:"room" db:"room" example:"room"`
}
