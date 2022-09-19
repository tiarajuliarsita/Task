package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Pokemon struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	BaseExpereience int    `json:"base_experience"`
	Height          int    `json:"height"`
	IsDefault       bool   `json:"is_default"`
	Order           int    `json:"order"`
	Weight          int    `json:"weight"`
}

func main() {
	http.HandleFunc("/pokemon", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		resp, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + id)
		if err != nil {
			log.Fatal("kamu gagal", err)
		}
		defer resp.Body.Close()

		pokemon := Pokemon{}
		respByte, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal("wah ada eror ges", err)
		}

		err = json.Unmarshal(respByte, &pokemon)
		if err != nil {
			log.Fatal("erorr again", err)
		}

		fmt.Println(pokemon)
		pokemonbyte, err := json.Marshal(pokemon)
		if err != nil {
			log.Fatal("erorr again", err)
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write(pokemonbyte)
	})

	if err := http.ListenAndServe(":5000", nil); err != nil {
		log.Fatal("erorr again", err)
	}

}
