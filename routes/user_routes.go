package routes

import (
  "todolist-backend/handlers"
  "github.com/gorilla/mux"
)

func SetupRoutes(userHandler *handlers.UserHandler) *mux.Router {
  r := mux.NewRouter()
  r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
  r.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
  r.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
  r.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
  r.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")
  return r
}
