package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type respData struct {
	Title string `json:"Title"`
	Desc string `json:"Desc"`
	Content string `json:"Content"`
}
type resp []respData

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/test", testHit).Methods(http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodOptions)
	router.HandleFunc("/test/json", jsonResponse).Methods(http.MethodGet)
	http.ListenAndServe(":8999", router)

}

func testHit(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Oi Oi Savaloy")
	fmt.Println("Endpoint Hit: testHit")

}

func jsonResponse(w http.ResponseWriter, r *http.Request) {
	articles := resp{
		respData{Title: "Json Response", Desc: "A Json Description", Content: "Some Json Content"},
	}
	fmt.Println("Endpoint Hit: Jsonnnnnn")
	json.NewEncoder(w).Encode(articles)
}