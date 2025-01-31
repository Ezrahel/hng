package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type Detail struct {
	Email      string `json:"email"`
	DateTime   string `json:"current_datetime"`
	GithubLink string `json:"github_url"`
}

func User(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		return
	}
	w.Header().Set("Content-Type", "Application/Json")
	w.WriteHeader(http.StatusOK)
	recentTime := time.Now().UTC()
	recent := recentTime.Format(time.RFC3339)
	details := Detail{
		Email:      "adelakinisrael024@gmail.com",
		DateTime:   recent,
		GithubLink: "https://github.com/ezrahel/hng",
	}
	response, err := json.Marshal(details)
	if err != nil {
		http.Error(w, "Error parsing data", http.StatusBadRequest)
	}
	w.Write(response)

}
func main() {
	http.HandleFunc("/", User)

	http.ListenAndServe(":8000", nil)
}
