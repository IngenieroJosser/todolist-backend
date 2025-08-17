package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "todolist-backend/models"
    "todolist-backend/repository"

    "github.com/gorilla/mux"
)

type UserHandler struct {
  Repo *repository.UserRepository
}

func NewUserHandler(repo *repository.UserRepository) *UserHandler {
  return &UserHandler{Repo: repo}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
  var user models.User
  if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
      http.Error(w, err.Error(), http.StatusBadRequest)
      return
  }
  if err := h.Repo.Create(&user); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
  }
  json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
  users, err := h.Repo.FindAll()
  if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
  }
  json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
  id, _ := strconv.Atoi(mux.Vars(r)["id"])
  user, err := h.Repo.FindByID(uint(id))
  if err != nil {
      http.Error(w, err.Error(), http.StatusNotFound)
      return
  }
  json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
  id, _ := strconv.Atoi(mux.Vars(r)["id"])
  user, err := h.Repo.FindByID(uint(id))
  if err != nil {
      http.Error(w, "Usuario no encontrado", http.StatusNotFound)
      return
  }
  if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
      http.Error(w, err.Error(), http.StatusBadRequest)
      return
  }
  if err := h.Repo.Update(&user); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
  }
  json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(mux.Vars(r)["id"])
    if err := h.Repo.Delete(uint(id)); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}
