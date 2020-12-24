package controllers

import (
	"profile/api/middlewares"
)

func (s *Server) InitRoutes() {
	//fs := http.FileServer(http.Dir("./assets"))
	//Home Route
	s.Router.HandleFunc("/", s.HomeEndpoint).Methods("GET")

	//Profile Routes
	s.Router.HandleFunc("/", middlewares.SetAuthMiddleware(s.CreateProfile)).Methods("POST")
	s.Router.HandleFunc("/list", middlewares.SetAuthMiddleware(s.ListProfiles)).Methods("POST")
	s.Router.HandleFunc("/image", middlewares.SetAuthMiddleware(s.GetAvatars)).Methods("GET")
	s.Router.HandleFunc("/{idProfile}", middlewares.SetAuthMiddleware(s.UpdateProfile)).Methods("PUT")
	s.Router.HandleFunc("/{idProfile}", middlewares.SetAuthMiddleware(s.GetProfile)).Methods("GET")
	s.Router.HandleFunc("/{idProfile}", middlewares.SetAuthMiddleware(s.DeleteProfile)).Methods("DELETE")

	//User Routes
	s.Router.HandleFunc("/user/login", s.Login).Methods("POST")
	s.Router.HandleFunc("/user/logout", s.Logout).Methods("POST")
	s.Router.HandleFunc("/user", middlewares.SetAuthMiddleware(s.CreateUser)).Methods("POST")

}
