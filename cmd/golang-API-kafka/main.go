package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Printf("GO API STARTING")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			fmt.Fprintln(w, "Hello world")
		}
	})
	http.HandleFunc("/now", func(w http.ResponseWriter, r *http.Request) {
		// Obter horario atual
		currentTime := time.Now()
		// Ajustar o horario para o fuso horário de Brasilia
		brasiliaTime := currentTime.In(time.FixedZone("Brasília", -3*60*60))
		// Retornar o horario no formato brasileiro
		fmt.Fprint(w, brasiliaTime.Format("2006-01-02 15:04:05"))
	})

	http.ListenAndServe(":8080", nil)
}
