package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/richardktran/go-employees/internal/controller/employee"
	httpEmployeeHandler "github.com/richardktran/go-employees/internal/handler/http"
	"github.com/richardktran/go-employees/internal/repository/file"
)

func main() {
	var port = flag.String("port", "8080", "port to listen on")
	flag.Parse()

	repo := file.New("data/employees.txt")
	ctrl := employee.New(repo)
	h := httpEmployeeHandler.New(ctrl)

	http.Handle("/employees", http.HandlerFunc(h.Handle))

	log.Println("Listening on port", *port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", *port), nil); err != nil {
		panic(err)
	}
}
