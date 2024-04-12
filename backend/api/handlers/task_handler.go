package handlers

import (
	"encoding/json"
	"log"

	"net/http"
	"time"
	"to-do-list/backend/pkg/task"

	"github.com/gorilla/mux"
)

// TaskHandler handles HTTP requests related to tasks.
type TaskHandler struct {
	taskService  *task.Service
}

// NewTaskHandler creates a new TaskHandler.
func NewTaskHandler(taskService *task.Service) *TaskHandler {
	return &TaskHandler{
		taskService:  taskService,
	}
}

// HandleTasks handles requests for creating and listing tasks.
// this is for when you don't have the ID
func (h *TaskHandler) HandleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.createTask(w, r)
	case http.MethodGet:
		h.listTasks(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// HandleTask handles requests for retrieving, updating, and deleting a task.
// this is for the one with the ID
func (h *TaskHandler) HandleTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	switch r.Method {
	case http.MethodGet:
		h.getTask(w, r, id)
	case http.MethodPut:
		h.updateTask(w, r, id)
	case http.MethodDelete:
		h.deleteTask(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *TaskHandler) createTask(w http.ResponseWriter, r *http.Request) {

	//log.Printf("Request Method: %s, URL: %s\n", r.Method, r.URL.String())

	var req struct {
		Title    string `json:"title"`
		Deadline string `json:"deadline"`
	}

	//log.Printf("Request Body: %+v\n", r)

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	deadline, err := time.Parse(time.RFC3339, req.Deadline)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task, err := h.taskService.CreateTask(req.Title, deadline)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)

}

func (h *TaskHandler) listTasks(w http.ResponseWriter, r *http.Request) {
	tasks := h.taskService.ListTasks()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) getTask(w http.ResponseWriter, r *http.Request, id string) {
	task, err := h.taskService.GetTask(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) updateTask(w http.ResponseWriter, r *http.Request, id string) {
	var req struct {
		Completed bool `json:"completed"`
	}

	log.Printf("%s", id)

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task, err := h.taskService.GetTask(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// call database update task
	err = h.taskService.MarkTaskComplete(id, req.Completed)
	// although the method name is 'MarkTaskComplete', it is used to update the task

	task.Completed = req.Completed
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) deleteTask(w http.ResponseWriter, r *http.Request, id string) {
	err := h.taskService.DeleteTask(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	
	w.WriteHeader(http.StatusOK)

}
