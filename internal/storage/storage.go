package storage

import "github.com/Aniket-Kumar-Paul/go-students-api/internal/types"

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error) // (id, error)
	GetStudentById(id int64) (types.Student, error)
}