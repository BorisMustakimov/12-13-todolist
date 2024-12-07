package hendlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/BorisMustakimov/12-13-todolist/service"
	"github.com/BorisMustakimov/12-13-todolist/task"
)

type TaskHandler struct {
	TaskService service.TaskService
}

func NewTaskHandler(taskService service.TaskService) *TaskHandler {
	return &TaskHandler{
		TaskService: taskService,
	}
}
func (h *TaskHandler) TaskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		id := r.URL.Query().Get("id")
		if id == "" {
			// Если id нет, выполняем добавление задачи
			h.AddTaskHandler(w, r)
		} else {
			// Если id есть, отмечаем задачу как выполненную
			h.DoneTaskHandler(w, r)
		}

	default:
		http.Error(w, `{"error":"Method not allowed"}`, http.StatusMethodNotAllowed)
	}
}

func (h *TaskHandler) AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task task.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, `{"error":"Failed to decode JSON"}`, http.StatusBadRequest)
		return
	}

	if task.Title == "" {
		http.Error(w, `{"error":"Empty title field"}`, http.StatusBadRequest)
		return
	}

	id, err := h.TaskService.AddTask(&task)
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(map[string]interface{}{"id": id})
}

func (h *TaskHandler) DoneTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, `{"error":"Missing task ID"}`, http.StatusBadRequest)
		return
	}

	now := time.Now()

	err := h.TaskService.TaskDone(id, now)
	if err != nil {
		if err.Error() == "task not found" {
			http.Error(w, `{"error":"Task not found"}`, http.StatusNotFound)
		} else {
			http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{})
}
