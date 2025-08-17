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
	RepoUser *repository.UserRepository
}

func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	return &UserHandler{RepoUser: repo}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.Password == "" {
		http.Error(w, "Password is required", http.StatusBadRequest)
		return
	}

	if err := h.RepoUser.Create(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Limpiar contraseña antes de responder
	user.Password = ""
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.RepoUser.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	// Limpiar contraseñas
	for i := range users {
		users[i].Password = ""
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	user, err := h.RepoUser.FindByID(uint(id))
	if err != nil {
		if err == repository.ErrUserNotFound {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	
	user.Password = ""
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	user, err := h.RepoUser.FindByID(uint(id))
	if err != nil {
		if err == repository.ErrUserNotFound {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Guardar contraseña actual
	currentPassword := user.Password
	
	var updateData struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Age      int    `json:"age"`
	}
	
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	// Actualizar solo los campos permitidos
	user.Name = updateData.Name
	user.Email = updateData.Email
	user.Age = updateData.Age
	
	// Manejar actualización de contraseña
	if updateData.Password != "" {
		user.Password = updateData.Password
	} else {
		user.Password = currentPassword
	}
	
	if err := h.RepoUser.Update(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	user.Password = ""
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if err := h.RepoUser.Delete(uint(id)); err != nil {
		if err == repository.ErrUserNotFound {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusNoContent)
}