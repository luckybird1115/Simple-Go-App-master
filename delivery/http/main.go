package main

import (
	_ "database/sql"
	"fmt"
	"github.com/Yuideg/firstApp/Student/repository"
	"github.com/Yuideg/firstApp/Student/services"
	"github.com/Yuideg/firstApp/delivery/http/handler"
	"github.com/Yuideg/firstApp/entity"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"html/template"
	"net/http"
)
var templ *template.Template

var err error
func createTables(dbconn *gorm.DB) []error {
	dbconn.Debug().DropTableIfExists(&entity.StudentInfo{})
	errs := dbconn.Debug().CreateTable(&entity.StudentInfo{}).GetErrors()
	if errs != nil {
		return errs
	}
	return nil
}


func main() {
	dbconn, err := gorm.Open("postgres",
		"postgres://postgres:yideg2378@localhost/student?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer dbconn.Close()
	fmt.Println("connected")
	//createTables(dbconn)
	router:=mux.NewRouter()
	templ = template.Must(template.ParseGlob("../../ui/template/*"))
	//fs := http.FileServer(http.Dir("../../ui/assets"))
	//http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	fs := http.FileServer(http.Dir("../../ui/assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))

	Repor := repository.NewGormRepositoryImpl(dbconn)
	Srv := services.NewGorServiceImpl(Repor)
	handler := handler.NewStudentHandler(templ, Srv)
	//router := httprouter.New()
	router.HandleFunc("/",handler.Home)
	router.HandleFunc("/user/students/all",handler.StudentsHandler)
	router.HandleFunc("/user/student/{id}",handler.StudentHandler)
	router.HandleFunc("/user/students/delete/{id}",handler.StudentDeleteHandler)
	router.HandleFunc("/user/students/new",handler.RegisterStudent)
	router.HandleFunc("/user/students/update",handler.UpdateStudentInfoHandler)
	router.HandleFunc("/user/students/update/{id}",handler.Update)
	http.ListenAndServe(":8080",router)

}