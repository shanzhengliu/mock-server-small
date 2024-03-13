package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
	err = json.Unmarshal(data, &routes)
	if err != nil {
		fmt.Println("json failed", err)
		return
	}

	for _, route := range routes {
		handler := createHandler(route)
		http.HandleFunc(route.Path, handler)
	}

	fmt.Println("listening :8080")
	http.ListenAndServe(":8080", nil)
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
