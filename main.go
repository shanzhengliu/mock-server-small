package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Route struct {
	Path         string `json:"path"`
	Method       string `json:"method"`
	ResponseBody string `json:"responseBody"`
	ContextType  string `json:"contextType"`
	Status       int    `json:"status"`
}

func main() {

	data, err := ioutil.ReadFile("routes.json")
	if err != nil {
		fmt.Println("file not exist", err)
		return
	}

	var routes []Route
	router := mux.NewRouter()
	err = json.Unmarshal(data, &routes)
	if err != nil {
		fmt.Println("json failed", err)
		return
	}

	for _, route := range routes {
		handler := createHandler(route)
		router.HandleFunc(route.Path, handler)
	}

	cor := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: false,
		AllowedHeaders:   []string{"*"},
	})

	corHandler := cor.Handler(router)

	fmt.Println("listening :8080")
	http.ListenAndServe(":8080", corHandler)
}

func createHandler(route Route) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == route.Method {
			w.Header().Set("Content-Type", route.ContextType)
			w.WriteHeader(route.Status)
			fmt.Fprint(w, route.ResponseBody)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
