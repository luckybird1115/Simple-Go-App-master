package Student

import (
	"github.com/Yuideg/firstApp/entity"
)

// StudentRepository specifies  students repositroy
type StudentRepository interface {
	Students() ([]entity.StudentInfo, []error)
	Student(id int) (*entity.StudentInfo, []error)
	UpdateStudentInfor(*entity.StudentInfo) (*entity.StudentInfo, []error)
	DeleteStudent(id int) (*entity.StudentInfo, []error)
	RegisterStudent(news entity.StudentInfo) (*entity.StudentInfo, []error)
}
