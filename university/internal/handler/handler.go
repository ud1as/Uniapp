package handler

import (
	"fmt"
	_ "github.com/Studio56School/university/docs"
	"github.com/Studio56School/university/internal/model"
	"github.com/Studio56School/university/internal/service"
	"github.com/Studio56School/university/internal/storage"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type Handler struct {
	repo *storage.Repo
	svc  service.IService
	log  *zap.Logger
}

type IHandler interface {
	GetStudents(c echo.Context) error
	GetStudentsById(c echo.Context) error
	CreateStudent(c echo.Context) error
	DeleteStudent(c echo.Context) error
	CreateProfessor(c echo.Context) error
	SignUp(c echo.Context) error
	SignIn(c echo.Context) error
	UserIdentity(next echo.HandlerFunc) echo.HandlerFunc
}

func NewHandler(svc service.IService, logger *zap.Logger) *Handler {
	return &Handler{log: logger, svc: svc}
}

// GetStudents godoc
// @Summary GetStudents Get all students
// @Description Get all students
// @Tags students
// @Accept json
// @Produce json
// @Success 200 {object} []model.Student
// @Router /students [get]
func (h *Handler) GetStudents(c echo.Context) error {

	fmt.Println(c.Get(userCtx))

	students, err := h.svc.AllStudentsService()
	if err != nil {
		h.log.Sugar().Error(err)
	}

	return c.JSON(http.StatusOK, students)

}

// @Summary		GetStudentsById
// @Description	Get student by id
// @Tags			students
// @ID				get-student
// @Accept			json
// @Produce		json
// @Success		200	{object}	model.Student
// @Router			/students/{id} [get]
func (h *Handler) GetStudentsById(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.log.Sugar().Error(err)
	}
	student, err := h.svc.StudentByID(id)

	if err != nil {
		h.log.Sugar().Error(err)
	}

	return c.JSON(http.StatusOK, student)
}

// @Summary		CreateStudent
// @Description	Create Student
// @Tags			students
// @ID				create-student
// @Accept			json
// @Produce		json
// @Param input body model.Student true "create account"
// @Success		200	{object}	model.Student
// @Router			/students/create [post]
func (h *Handler) CreateStudent(c echo.Context) error {
	var request model.Student
	err := c.Bind(&request)
	if err != nil {
		h.log.Sugar().Error(err)
	}

	student, err := h.svc.AddNewStudent(request)
	if err != nil {
		h.log.Sugar().Error(err)

	}

	return c.JSON(http.StatusOK, student)
}

// @Summary		DeleteStudent
// @Description	Delete Student
// @Tags			students
// @ID				delete-student
// @Accept			json
// @Produce		json
// @Success 200 {string} string "Successful deleted user with id"
// @Router			/students/{id} [delete]
func (h *Handler) DeleteStudent(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		h.log.Sugar().Error(err)
	}

	err = h.svc.DeleteStudentById(id)
	if err != nil {
		h.log.Sugar().Error(err)
	}

	defaultString := "Successful deleted user with id"

	return c.JSON(http.StatusOK, map[string]interface{}{
		defaultString: id,
	})
}

func (h *Handler) CreateProfessor(c echo.Context) error {
	var request *model.CreateProfessorRequest
	err := c.Bind(&request)
	if err != nil {
		h.log.Sugar().Error(err)
	}

	professor, err := h.svc.CreateNewProfessorService(request)
	if err != nil {
		h.log.Sugar().Error(err)
	}

	return c.JSON(http.StatusOK, professor)
}
