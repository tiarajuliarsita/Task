package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	// "strconv"
)

type Note struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}
type Request struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
type Response struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

var (
	notes []Note
)

// respon
func sendResponse(w http.ResponseWriter, code int, status string, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	var response Response
	response.Code = code
	response.Status = status
	response.Message = message

	dByte, _ := json.Marshal(response)

	w.Write(dByte)
}

// GetNote
func NOTES(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// id := r.URL.Query().Get("id")
		// if id != "" {
		// 	idT, _ := strconv.Atoi(id)
		// 	notes[idT]
		// }
		DNote, err := json.Marshal(notes)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(DNote)
		return
	}

	if r.Method == "POST" {
		var nt Note
		decodeJSON := json.NewDecoder(r.Body)
		if err := decodeJSON.Decode(&nt); err != nil {
			log.Fatal(err)
		}
		nt.Id = len(notes) + 1
		notes = append(notes, nt)

		DNote, _ := json.Marshal(nt) // ke byte
		w.Write(DNote)
		return
	}

	http.Error(w, "hayo  capek ya wkwkw?", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/notes", NOTES)
	fmt.Println("on proses")
	if err := http.ListenAndServe(":7000", nil); err != nil {
		log.Fatal(err)
	}
}

