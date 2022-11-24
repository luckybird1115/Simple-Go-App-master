package entity

//student info entity with there data type and json name
type StudentInfo struct {
	ID       int   `gorm:"SERIAL PRIMARY KEY" json:"id"`
	FullName string `gorm:"type:varchar(255);not null" json:"full_name"`
	Email    string `gorm:"type:varchar(255);not null; UNIQUE" json:"email"`
	Image    string `gorm:"type:varchar(255);not null" json:"image"`
}
