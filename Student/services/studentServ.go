package services

import (
	"github.com/Yuideg/firstApp/Student"
	"github.com/Yuideg/firstApp/entity"
)

// GormServiceImpl implements Student.StudentRepository interface
type GormServiceImpl struct {
	sturepo Student.StudentRepository
}

// NewGorServiceImpl will create new NewGorServiceImpl object
func NewGorServiceImpl(st Student.StudentRepository) *GormServiceImpl {
	return &GormServiceImpl{sturepo: st}
}

// Students returns list of all Student that exist in the database
func (rs *GormServiceImpl) Students() ([]entity.StudentInfo, []error) {

	news, err := rs.sturepo.Students()

	if err != nil {
		return nil, err
	}

	return news, nil
}

// RegisterStudent registers  new students  information
func (rs *GormServiceImpl) RegisterStudent(st entity.StudentInfo) (*entity.StudentInfo, []error) {

	r,err := rs.sturepo.RegisterStudent(st)

	if err != nil {
		return nil,err
	}

	return r,nil
}

//Student returns  single student info identified by id
func (rs *GormServiceImpl)Student(id int) (*entity.StudentInfo, []error) {

	r, err := rs.sturepo.Student(int(id))

	if err != nil {
		return r, err
	}

	return r, nil
}

// UpdateStudentInfor updates a StudentsInfo with new data
func (rs *GormServiceImpl) UpdateStudentInfor(st *entity.StudentInfo) (*entity.StudentInfo, []error) {
	r, err := rs.sturepo.UpdateStudentInfor(st)
	if err != nil {
		return r, err
	}
	return r, nil
}
// DeleteStudent delete a students by its id
func (rs *GormServiceImpl) DeleteStudent(id int) (*entity.StudentInfo, []error) {
	r,err := rs.sturepo.DeleteStudent(int(id))
	if err != nil {
		return r,err
	}
	return r,nil
}
