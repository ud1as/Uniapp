package storage

import (
	"context"
	"fmt"
	"github.com/Studio56School/university/internal/config"
	"github.com/Studio56School/university/internal/model"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

func NewRepository(conf *config.Config, logger *zap.Logger) (*Repo, error) {
	pgDB, err := ConnectDB(conf)
	if err != nil {
		logger.Sugar().Error("Unable to connect")
		return nil, err
	}
	return &Repo{DB: pgDB}, nil
}

type Repo struct {
	DB *pgx.Conn
}

func (r *Repo) StudentByID(ctx context.Context, id int) (student model.Student, err error) {

	query := `select id, name, surname, gender from students where id = $1 `
	err = r.DB.QueryRow(ctx, query, id).Scan(&student.Id, &student.Name, &student.Surname, &student.Gender)

	if err != nil {
		//r.l.Sugar().Error(fmt.Sprintf("Не отработался запрос студентам по id: %s", err))
		return student, err
	}

	return student, err
}

func (r *Repo) AllStudents(ctx context.Context) (students []model.Student, err error) {

	students = make([]model.Student, 0)
	query := `select id, name, surname, gender from students`
	rows, err := r.DB.Query(ctx, query)
	if err != nil {
		//r.l.Sugar().Error(fmt.Sprintf("Не отработался запрос студентам по id: %s", err))
		return nil, err
	}
	var student model.Student

	for rows.Next() {
		err := rows.Scan(&student.Id, &student.Name, &student.Surname, &student.Gender)
		if err != nil {
			//r.l.Sugar().Error(fmt.Sprintf("Не отработался запрос студентам по id: %s", err))
			return nil, err
		}

		students = append(students, student)
	}

	defer rows.Close()
	return students, nil
}

func (r *Repo) AddNewStudent(ctx context.Context, student model.Student) (id int, err error) {

	query := `INSERT INTO public.students
	(name, surname, gender)
	VALUES ($1, $2, $3) RETURNING id`

	err = r.DB.QueryRow(ctx, query, student.Name, student.Surname, student.Gender).Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (r *Repo) UpdateStudent(ctx context.Context, student *model.Student) (err error) {

	tx, err := r.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	query := `UPDATE public.students
	SET name=$2, surname = $3, gender = $4 
	WHERE id = $1;`

	_, err = tx.Exec(ctx, query, student.Id, student.Name, student.Surname, student.Gender)

	if err != nil {
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			return fmt.Errorf("ERROR: transaction: %s", errTX)
		}

		return fmt.Errorf("error occurred while updating students info in users table: %v", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			return fmt.Errorf("ERROR: transaction error: %s", errTX)
		}
		return fmt.Errorf("error occurred while updating students info in users table: %v", err)
	}

	return err
}

func (r *Repo) DeleteStudentById(ctx context.Context, id int) error {

	query := `DELETE FROM students_by_group WHERE student_id = $1`
	query2 := `DELETE FROM students WHERE id = $1`

	rows, err := r.DB.Exec(ctx, query, id)
	rows, err = r.DB.Exec(ctx, query2, id)
	if err != nil {
		return fmt.Errorf("error occurred while deleting patient: %v", err)
	}
	if rows.RowsAffected() < 1 {
		return fmt.Errorf("error: no patient in db with such id %d", id)
	}
	return nil

}

func (r *Repo) CreateProfessor(professor *model.CreateProfessorRequest) (*model.CreateProfessorResponse, error) {

	var Id int
	query := `INSERT INTO public.professors
	(name, surname, email, degree)
	VALUES ($1, $2, $3, $4) RETURNING id`

	err := r.DB.QueryRow(context.Background(), query, professor.Name, professor.Surname, professor.Email, professor.Degree).Scan(&Id)

	if err != nil {
		return nil, fmt.Errorf("error occurred while creating professor in system: %v", err)
	}

	return &model.CreateProfessorResponse{
		Id: Id,
	}, nil
}

func (r *Repo) CreateUser(user model.User) (int, error) {

	var id int

	query := `INSERT INTO public.users
	(name, username, password)
	VALUES ($1, $2, $3) RETURNING id`

	err := r.DB.QueryRow(context.Background(), query, user.Name, user.Username, user.Password).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("error occurred while creating users in system: %v", err)
	}

	return id, nil
}

func (r *Repo) GetUser(username, password string) (model.User, error) {
	var user model.User

	query := `SELECT id, name FROM public.users
	WHERE username=$1 AND password=$2`

	row := r.DB.QueryRow(context.Background(), query, username, password)

	err := row.Scan(&user.Id, &user.Name)

	if err != nil {
		return model.User{}, fmt.Errorf("error occurred while creating users in system: %v", err)
	}

	return user, err
}
