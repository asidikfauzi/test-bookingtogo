package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"test-bookingtogo/controllers"
	"test-bookingtogo/lib/utils"
)

func NewRouter() {
	r := mux.NewRouter()

	//--------start nationality group router
	nationalityRouter := r.PathPrefix("/nationality").Subrouter()
	nationalityRouter.HandleFunc("", controllers.GetNationalities).Methods("GET")
	nationalityRouter.HandleFunc("/{id}", controllers.GetNationalityById).Methods("GET")
	nationalityRouter.HandleFunc("", controllers.CreateNationality).Methods("POST")
	nationalityRouter.HandleFunc("/{id}", controllers.UpdateNationality).Methods("PUT")
	nationalityRouter.HandleFunc("/{id}", controllers.DeleteNationality).Methods("DELETE")
	//--------end nationality group router

	// start server
	fmt.Println("RUN", utils.GetEnv("PORT"))
	http.ListenAndServe(":"+utils.GetEnv("PORT"), r)
}
