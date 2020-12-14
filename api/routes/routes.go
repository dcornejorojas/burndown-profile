package routes

import "github.com/gorilla/mux"

//InitRoutes init all routes of the API

// func InitRoutes() *mux.Router {
// 	//fs := http.FileServer(http.Dir("./assets"))
// 	router := mux.NewRouter()
// 	router = SetUtilsRouter(router)
// 	router = SetUserRoutes(router)
// 	router = SetProfileRoutes(router)
// 	return router
// }

func InitRoutes() *mux.Router {
	//fs := http.FileServer(http.Dir("./assets"))
	router := mux.NewRouter()
	router = SetUtilsRouter(router)
	router = SetUserRoutes(router)
	router = SetProfileRoutes(router)
	return router
}