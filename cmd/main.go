package main

import (
	"log"
	"net/http"

	"github.com/richardktran/go-employees/internal/controller/employee"
	httpEmployeeHandler "github.com/richardktran/go-employees/internal/handler/http"
	"github.com/richardktran/go-employees/internal/repository/file"
)

func main() {
	repo := file.New("employees.txt")
	ctrl := employee.New(repo)
	h := httpEmployeeHandler.New(ctrl)

	http.Handle("/employees", http.HandlerFunc(h.Handle))

	log.Println("Listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
