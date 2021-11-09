package server

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"todoapp-golang/handler"
)

type Server struct {
	chi.Router
}

func SetupRoutes() *Server {
	router := chi.NewRouter()
	router.Route("/to-do-app", func(api chi.Router) {
		api.Post("/add-user",handler.AddUser)
		api.Post("/add-task",handler.AddTask)
		api.Post("/delete-task",handler.DeleteTask)
		api.Get("/get-all-task",handler.GetTask)

	})
	return &Server{router}
}

func (svc *Server) Run(port string) error {
	return http.ListenAndServe(port, svc)
}


