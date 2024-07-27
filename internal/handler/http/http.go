package http

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/richardktran/go-employees/internal/controller/employee"
	"github.com/richardktran/go-employees/pkg/model"
)

type Handler struct {
	ctrl *employee.Controller
}

func New(ctrl *employee.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		h.GetEmployees(w, r)
		return
	}

	if r.Method == http.MethodPost {
		h.AddEmployee(w, r)
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

// GetEmployees returns all employees
func (h *Handler) GetEmployees(w http.ResponseWriter, r *http.Request) {
	var paging model.Paging
	queryString := r.URL.Query()

	paging.Page, _ = strconv.Atoi(queryString.Get("page"))
	paging.Limit, _ = strconv.Atoi(queryString.Get("limit"))

	search := queryString.Get("search")

	paging.Process()

	employees, err := h.ctrl.GetEmployees(r.Context(), search, &paging)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(model.EmployeeResponse{
		Status: "success",
		Data:   employees,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// AddEmployee adds a new employee
func (h *Handler) AddEmployee(w http.ResponseWriter, r *http.Request) {
	var employee model.EmployeeCreation
	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Adding employee: %+v", employee)
	if err := h.ctrl.AddEmployee(r.Context(), employee); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(employee); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
