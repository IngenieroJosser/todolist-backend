package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"todolist-backend/models"
	"todolist-backend/repository"

	"github.com/gorilla/mux"
)

type TaskHandler struct {
	RepoTask *repository.TaskRepository
}

func NewTaskHandler(repoTask *repository.TaskRepository) *TaskHandler {
	return &TaskHandler{RepoTask: repoTask}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} 
	if err := h.RepoTask.CreateTask(&task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.RepoTask.FindAllTask()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) GetTaskById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
  user, err := h.RepoTask.FindTaskByID(uint(id))
  if err != nil {
      http.Error(w, err.Error(), http.StatusNotFound)
      return
  }
  json.NewEncoder(w).Encode(user)
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	task, err := h.RepoTask.FindTaskByID(uint(id))
  if err != nil {
      http.Error(w, "Usuario no encontrado", http.StatusNotFound)
      return
  }
  if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
      http.Error(w, err.Error(), http.StatusBadRequest)
      return
  }
  if err := h.RepoTask.UpdateTask(&task); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
  }
  json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if err := h.RepoTask.DeleteTask(uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
