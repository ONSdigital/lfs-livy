package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

type respData struct {
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"Content"`
}
type resp []respData

func main() {

	router := mux.NewRouter()

	// Test Routes
	router.HandleFunc("/test", testHit).Methods(http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodOptions)
	router.HandleFunc("/test/json", jsonResponse).Methods(http.MethodGet)

	router.HandleFunc("/test/livy/session", livySessionTest).Methods(http.MethodGet)
	router.HandleFunc("/test/livy/batch", livyBatchTest).Methods(http.MethodGet)
	router.HandleFunc("/test/livy/batch/no/{no}", livyBatchGetTest).Methods(http.MethodGet)

	router.HandleFunc("/test/livy/batch/post", livyBatchPostTest).Methods(http.MethodGet)

	// Routes
	router.HandleFunc("/livy/batch/post", livyBatchPostLFS).Methods(http.MethodGet)

	http.ListenAndServe(":8999", router)
}

func livyBatchPostLFS(w http.ResponseWriter, r *http.Request) {
	fmt.Println("LFS Batches Posty")
	fmt.Fprintln(w, "LFS Batches Posty")

	//className := "uk.gov.ons.lfs.LFSMonthly"
	jsonJarStr := []byte(`{
		"className" :  "uk.gov.ons.lfs.LFSMonthly",
		"file"  : "/Users/andrewurquhart/Documents/Repositories/GitHub/lfs-monthly/target/scala-2.11/lfs-monthly-assembly-1.0.jar",
		"executorMemory": "20g",
        "args": [2000],
    	"proxyUser": "andyyyyyyyyy",
	    "queue": "default",
		"conf": {"DB_USER": "lfs", "DB_PASSWORD": "lfs", "DB_URI": "jdbc:postgresql://localhost:5432/lfs"}

	}`)

	postLivy(w, "batches", jsonJarStr)
}

func livyBatchPostTest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Livy Batches Posty")
	fmt.Fprintln(w, "Livy Batches Posty")

	jsonJarStr := []byte(`{
		"className" :  "com.cloudera.sparkwordcount.SparkWordCount",
		"file"  : "/Users/andrewurquhart/Documents/Repositories/GitHub/simplesparkapp/target/sparkwordcount-0.0.1-SNAPSHOT.jar",
		"executorMemory": "20g",
        "args": [2000],
    	"proxyUser": "andyyyyyyyyy",
	    "queue": "default"


	}`)

	postLivy(w, "batches", jsonJarStr)
}

func livyBatchGetTest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Livy Batches Getty")
	fmt.Fprintln(w, "Livy Batches Getty")

	vars := mux.Vars(r)
	no := vars["no"]

	getLivy(w, "batches/"+no+"/log")
}

func livyBatchTest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Livy Batches Getty")
	fmt.Fprintln(w, "Livy Batches Getty")

	getLivy(w, "batches")
}

func livySessionTest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fuckn Livy Kick Off")
	fmt.Fprintln(w, "Fuckn Livy Kick Off")

	getLivy(w, "sessions")
}

func jsonResponse(w http.ResponseWriter, r *http.Request) {
	articles := resp{
		respData{Title: "Json Response", Desc: "A Json Description", Content: "Some Json Content"},
	}
	fmt.Println("Endpoint Hit: Jsonnnnnn")
	json.NewEncoder(w).Encode(articles)
}

func testHit(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Oi Oi Savaloy")
	fmt.Println("Endpoint Hit: testHit")

}

func getLivy(w http.ResponseWriter, path string) {
	url := "http://localhost:8998/" + path
	var client http.Client
	resp, _ := client.Get(url)

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Fprintln(w, string(body))
}

func postLivy(w http.ResponseWriter, path string, json []byte) {
	url := "http://localhost:8998/" + path
	var client http.Client

	resp, err := client.Post(url, "Content-Type: application/json", bytes.NewBuffer(json))

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Fprintln(w, string(body))
	fmt.Println(string(body))
	fmt.Fprintln(w, err)
}
