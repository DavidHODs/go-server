package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/gorilla/mux"
)

// middleware prints to the terminal the URL accessed as well as the edited response from the URLs
func middleware (next http.Handler) http.Handler {
	return http.HandlerFunc(func (res http.ResponseWriter, req *http.Request) {
		fmt.Printf("You accessed %v", req.URL)
		fmt.Println()

		rec := httptest.NewRecorder()
		next.ServeHTTP(rec, req)

		recRes := rec.Result()
		resBody, _ := ioutil.ReadAll(recRes.Body)

		fmt.Printf("Middleware added this response to: %v", string(resBody))
		fmt.Println()

		next.ServeHTTP(res, req)
	}) 
}

func info(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Solving API observability and documentation via automation and AI\n")
}

func client(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "The API Toolkit golang client is an open source sdk used to integrate golang web services with APIToolkit. It monitors incoming traffic, gathers the requests and sends the request to the apitoolkit servers.\n")
}

func server(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Definitely not source. Receives info from our open source API Toolkit client before sending to gcp pubsub.\n")
}

func routes() http.Handler {
	router := mux.NewRouter()
	router.Use(middleware)
	router.HandleFunc("/apitoolkit/info", info).Methods("GET")
	router.HandleFunc("/apitoolkit/client", client).Methods("GET")
	router.HandleFunc("/apitoolkit/server", server).Methods("GET")

	return router
}

func main() {
	srv := &http.Server{
        Addr: ":8000",
        Handler: routes(),
    }

	fmt.Println("Listening on port 8000")

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}