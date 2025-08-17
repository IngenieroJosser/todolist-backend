package routes

import (
  "todolist-backend/handlers"
  "github.com/gorilla/mux"
)

func SetupRoutes(
  userHandler *handlers.UserHandler,
  taskHandler *handlers.TaskHandler,
  ) *mux.Router {
  r := mux.NewRouter()

  // Rutas de usuarios
  r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
  r.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
  r.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
  r.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
  r.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

  // Rutas de tareas
  r.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")
  r.HandleFunc("/tasks", taskHandler.GetAllTasks).Methods("GET")
  r.HandleFunc("/tasks/{id}", taskHandler.GetTaskById).Methods("GET")
  r.HandleFunc("/tasks/{id}", taskHandler.UpdateTask).Methods("PUT")
  r.HandleFunc("/tasks/{id}", taskHandler.DeleteTask).Methods("DELETE")

  return r
}