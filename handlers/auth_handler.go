package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todolist-backend/models"
	"todolist-backend/repository"
)

type AuthHandler struct {
	Repo *repository.UserRepository
}

func NewAuthHandler(repo *repository.UserRepository) *AuthHandler {
	return &AuthHandler{Repo: repo}
}

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	type credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var creds credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.Repo.FindByEmail(creds.Email)
	if err != nil {
		if err == repository.ErrUserNotFound {
			http.Error(w, "Credenciales invalidas", http.StatusUnauthorized)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if err := user.CheckPassword(creds.Password); err != nil {
		http.Error(w, "Credenciales invalidas", http.StatusUnauthorized)
		return
	}

	// Convertir ID a string correctamente
	userIDStr := strconv.FormatUint(uint64(user.ID), 10)

	response := map[string]string{
		"message": "Inicio de sesión exitoso",
		"user_id": userIDStr,
		// En producción, agregar token JWT aquí
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Handler para registro de usuarios
func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.Password == "" {
		http.Error(w, "La contraseña es requerida", http.StatusBadRequest)
		return
	}

	if err := h.Repo.Create(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Limpiar contraseña en respuesta
	user.Password = ""
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}