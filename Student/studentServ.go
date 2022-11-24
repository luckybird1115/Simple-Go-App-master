package Student

import (
	"github.com/Yuideg/firstApp/entity"
)

// StudentServices specifies  Student services
type StudentServices interface {
	Students() ([]entity.StudentInfo, []error)
	Student(id int) (*entity.StudentInfo, []error)
	UpdateStudentInfor(*entity.StudentInfo) (*entity.StudentInfo, []error)
	DeleteStudent(id int) (*entity.StudentInfo, []error)
	RegisterStudent(news entity.StudentInfo) (*entity.StudentInfo, []error)
}
