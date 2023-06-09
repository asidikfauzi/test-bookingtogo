package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
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

	//--------start customers group router
	customerRouter := r.PathPrefix("/customer").Subrouter()
	customerRouter.HandleFunc("", controllers.GetCustomers).Methods("GET")
	customerRouter.HandleFunc("/{id}", controllers.GetCustomerById).Methods("GET")
	customerRouter.HandleFunc("", controllers.CreateCustomer).Methods("POST")
	customerRouter.HandleFunc("/{id}", controllers.UpdateCustomer).Methods("PUT")
	customerRouter.HandleFunc("/{id}", controllers.DeleteCustomer).Methods("DELETE")
	//--------end nationality group router

	//--------start family list group router
	familyListRouter := r.PathPrefix("/family-list").Subrouter()
	familyListRouter.HandleFunc("", controllers.GetFamilyLists).Methods("GET")
	familyListRouter.HandleFunc("/{id}", controllers.GetFamilyListById).Methods("GET")
	familyListRouter.HandleFunc("", controllers.CreateFamilyList).Methods("POST")
	familyListRouter.HandleFunc("/{id}", controllers.UpdateFamilyList).Methods("PUT")
	familyListRouter.HandleFunc("/{id}", controllers.DeleteFamilyList).Methods("DELETE")
	//--------end family list group router

	// start server
	fmt.Println("RUN", utils.GetEnv("PORT"))
	err := http.ListenAndServe(":"+utils.GetEnv("PORT"), r)
	if err != nil {
		log.Fatal(err)
	}
}
