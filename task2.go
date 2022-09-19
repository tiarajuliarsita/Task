package main

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
	Code    int    `json:"id"`
	Status  string `json:"status"`
	Message string `json:"message"`
}


var (
	notes []Note
)

func sendResponse(w http.ResponseWriter, code int, message string, status string{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	var response Response
	response.Code = code
	response.Status = status
	response.Message = message	


	noteByte, err := json.Marshal(response)
	if err != nil {
		log.Fatal("ada error", err)
	}

	w.Write(noteByte)
}


func main(){
	http.HandleFunc("/notes", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			defer r.Body.Close()
			dByte, err := io.ReadAll(r.Body)
			if err != nil {
				sendResponse(w, http.StatusInternalServerError, "Terjadi Masalah", nil)
				return
			

			var note Note
			index := len(notes)
			json.Unmarshal(dByte, &note)

			note.Index = index

			notes = append(notes, note)
			sendResponse(w, http.CodeCreated, 200, "")
			sendResponse(w, http.StatusCreated, "succes", "")
			sendResponse(w, http.MessageCreated, "Note baru Berhasil diTambahkan", "")
			return
		}
		// if r.Method == http.MethodDelete {
		// 	id := r.URL.Query().Get("id")

		// 	idT, _ := strconv.Atoi(id)
		// 	remove(notes, idT)
		// 	sendResponse(w, http.StatusOK, "Data Berhasil dihapus", "")
		// 	return

		// }
		}
		
	}









// func TambahData(){

// }

// func UbahData(){

// }

// func HapusData(){

// }

// }
