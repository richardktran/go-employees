package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/richardktran/go-employees/configs/env"
	"github.com/richardktran/go-employees/internal/controller/employee"
	httpEmployeeHandler "github.com/richardktran/go-employees/internal/handler/http"
	"github.com/richardktran/go-employees/internal/repository/mysql"
)

func init() {
	env.Setup()
}

func main() {
	var port = flag.String("port", "8080", "port to listen on")
	flag.Parse()

	// repo := file.New("data/employees.txt")
	repo, err := mysql.New()
	if err != nil {
		panic(err)
	}

	ctrl := employee.New(repo)
	h := httpEmployeeHandler.New(ctrl)

	http.Handle("/employees", http.HandlerFunc(h.Handle))
	http.Handle("PUT /employees/{id}", http.HandlerFunc(h.UpdateEmployee))
	http.Handle("DELETE /employees/{id}", http.HandlerFunc(h.DeleteEmployee))

	log.Println("Listening on port", *port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", *port), nil); err != nil {
		panic(err)
	}
}
