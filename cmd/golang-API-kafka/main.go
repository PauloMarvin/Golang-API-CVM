package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Data struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	// Verifique se a método da requisição é POST
	if r.Method != http.MethodPost {
		http.Error(w, "<Método não permitido>", http.StatusMethodNotAllowed)
		fmt.Print("<Método não permitido>")
		return
	}

	// Verifique se o cabeçalho "Content-Type" é "application/json"
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		http.Error(w, "<Tipo de conteúdo não suporte>", http.StatusUnsupportedMediaType)
		return
	}

	// Crie um novo objeto Data para armazenar os dados do corpo da requisição
	data := Data{}

	// Desencode o corpo da requisição em um objeto Data
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "<Erro ao ler dados de entrada>", http.StatusBadRequest)
		return
	}

	// Processamento do dado para a cada requisição retorna um json com o datetime.now
	json_data := fmt.Sprintf("{Nome:%v, Age:%v, Datetime:%v}\n", data.Name, data.Age, time.Now().Format(time.RFC822))
	fmt.Printf(json_data)

	// Retorne uma resposta apropriada
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("<Dados recebidos com sucesso!>\n"))
}

func main() {
	fmt.Print("<GO API STARTING>\n")
	// rota para post
	http.HandleFunc("/data", apiHandler)

	http.ListenAndServe(":8000", nil)
}
