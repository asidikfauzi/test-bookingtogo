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
	r.HandleFunc("/nationality", controllers.GetNationalities).Methods("GET")
	r.HandleFunc("/nationality/{id}", controllers.GetNationalityById).Methods("GET")
	r.HandleFunc("/nationality", controllers.CreateNationality).Methods("POST")
	r.HandleFunc("/nationality/{id}", controllers.UpdateNationality).Methods("PUT")
	r.HandleFunc("/nationality/{id}", controllers.DeleteNationality).Methods("DELETE")

	// start server
	fmt.Println("RUN", utils.GetEnv("PORT"))
	http.ListenAndServe(":"+utils.GetEnv("PORT"), r)
}
