package main

import (
	"encoding/json"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"os"
)

type Holiday struct {
	Name      string `json:"name"`
	DateStart string `json:"dateStart"`
	DateEnd   string `json:"dateEnd"`
	TotalDays string `json:"totalDays"`
}

type Group struct {
	Group   string    `json:"group"`
	State   []string  `json:"state"`
	Session string    `json:"session"`
	Holiday []Holiday `json:"holidays"`
}

func loadGroups() ([]Group, error) {
	var groups []Group
	data, err := os.ReadFile("holidays.json") // use os.ReadFile instead of ioutil.ReadFile
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &groups)
	if err != nil {
		return nil, err
	}
	return groups, nil
}

func holidayHandler(w http.ResponseWriter, r *http.Request) {
	groups, err := loadGroups()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(groups)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	http.Handle("/holidays", handlers.CORS()(http.HandlerFunc(holidayHandler)))
	log.Print("Server running at :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
