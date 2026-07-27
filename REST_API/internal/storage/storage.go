package storage

import "github.com/Abdullah-Builds/Go_Projects/internal/types"


type Storage interface {
	CreateStudent(name string, age int64, email string) (int64, error)
	GetStudentByID(id int) (*types.Student, error)
	GetStudents(limit int, cursor int) ([]types.Student, *int, error)
}
