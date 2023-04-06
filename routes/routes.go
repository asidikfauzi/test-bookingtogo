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
	r.HandleFunc("/", controllers.GetNationalities).Methods("GET")
	r.HandleFunc("/", controllers.CreateNationality).Methods("POST")

	// start server
	fmt.Println("RUN", utils.GetEnv("PORT"))
	http.ListenAndServe(":"+utils.GetEnv("PORT"), r)
}
