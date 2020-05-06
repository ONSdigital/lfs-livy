package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/run/monthly", testHit).Methods(http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodOptions)
	http.ListenAndServe(":8999", router)

}

func testHit(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Oi Oi Savaloy")
}