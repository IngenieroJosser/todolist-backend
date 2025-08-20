package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"todolist-backend/models"
	"todolist-backend/repository"

	"github.com/gorilla/mux"
)

type ProjectHandler struct {
	RepoProject *repository.ProjectRepository
}

func NewProjectHandler(repoProject *repository.ProjectRepository) *ProjectHandler {
	return &ProjectHandler{RepoProject: repoProject}
}

func (h *ProjectHandler) CreateProject(w http.ResponseWriter, r *http.Request) {
	var project models.Project
	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.RepoProject.CreateProject(&project); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(project)
}

func (h *ProjectHandler) GetAllProjects(w http.ResponseWriter, r *http.Request) {
	project, err := h.RepoProject.FindAllProject()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(project)
}

func (h *ProjectHandler) GetProjectById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
  user, err := h.RepoProject.FindProjectByID(uint(id))
  if err != nil {
      http.Error(w, err.Error(), http.StatusNotFound)
      return
  }
  json.NewEncoder(w).Encode(user)
}

func (h *ProjectHandler) UpdateProject(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	project, err := h.RepoProject.FindProjectByID(uint(id))
  if err != nil {
      http.Error(w, "Projecto no encontrado", http.StatusNotFound)
      return
  }
  if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
      http.Error(w, err.Error(), http.StatusBadRequest)
      return
  }
  if err := h.RepoProject.UpdateProject(&project); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
  }
  json.NewEncoder(w).Encode(project)
}

func (h *ProjectHandler) DeleteProject(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if err := h.RepoProject.DeleteProject(uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
