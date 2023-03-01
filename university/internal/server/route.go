package server

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func (s *Server) routeApiV1(r *echo.Echo) {

	// r.Use(s.handlers.university.UserIdentity)

	auth := r.Group("/api/v1/auth")
	{
		auth.POST("/sign-up", s.handlers.university.SignUp)
		auth.POST("/sign-in", s.handlers.university.SignIn)
	}
	// хочу чтоб при переходе на данные эндпойнты постман жаловался на отсутствия header если его нет
	// чтоб перейти нужно передавать bearer token in authorization
	// ошибка Cannot use 's.handlers.university.UserIdentity' (type func() echo.HandlerFunc) as the type MiddlewareFunc

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/students", s.handlers.university.GetStudents, s.handlers.university.UserIdentity)
		apiv1.GET("/students/:id", s.handlers.university.GetStudentsById)
		apiv1.POST("/students/create", s.handlers.university.CreateStudent)
		apiv1.DELETE("/students/:id", s.handlers.university.DeleteStudent)
		apiv1.POST("/professor/create", s.handlers.university.CreateProfessor)
	}

}

func (s *Server) routeSwagger(r *echo.Echo) {
	r.GET("/swagger/*", echoSwagger.WrapHandler)
}
