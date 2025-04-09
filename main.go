package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Define o header como JSON
		w.Header().Set("Content-Type", "application/json")

		// Monta a resposta
		response := map[string]string{
			"data": "hello world",
		}

		// Converte e envia
		json.NewEncoder(w).Encode(response)
	})

	// Inicia o servidor
	http.ListenAndServe(":3333", nil)
}
