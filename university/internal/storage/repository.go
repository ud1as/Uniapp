package storage

import (
	"context"
	"github.com/Studio56School/university/internal/model"
)

type IRepository interface {
	AllStudents(ctx context.Context) (students []model.Student, err error)
	StudentByID(ctx context.Context, id int) (student model.Student, err error)
	DeleteStudentById(ctx context.Context, id int) error
	UpdateStudent(ctx context.Context, student *model.Student) (err error)
	AddNewStudent(ctx context.Context, student model.Student) (id int, err error)
	CreateProfessor(professor *model.CreateProfessorRequest) (*model.CreateProfessorResponse, error)
	CreateUser(user model.User) (int, error)
}
