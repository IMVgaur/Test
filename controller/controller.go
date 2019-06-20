package controller

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//Handlers ...
func Handlers() http.Handler {
	server := mux.NewRouter()
	server.HandleFunc("/emp", AddEmployee).Methods("POST")
	server.HandleFunc("/emps", GetEmployees).Methods("GET")
	server.HandleFunc("/emp/update/{id}", UpdateEmp).Methods("UPDATE")
	server.HandleFunc("/emp/delete/{id}", RemoveEmp).Methods("DELETE")
	server.HandleFunc("/emp/{id}", GetEmployee).Methods("GET")
	return server
}

//AddEmployee ...
func AddEmployee(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside AddEmployee")
}

//GetEmployee ...
func GetEmployee(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside GetEmployee")

}

//GetEmployees ...
func GetEmployees(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside GetEmployees")

}

//UpdateEmp ...
func UpdateEmp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside UpdateEmp")

}

//RemoveEmp ...
func RemoveEmp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside RemoveEmp")

}
