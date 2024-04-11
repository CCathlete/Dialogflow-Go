package main

import (
	"Golangbot/pkg/dialogflowUtils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var dfp dialogflowUtils.DialogflowProcessor

func main() {
	dfp.Init("prototypechatbpt-d4db0", "prototypechatbot-d4db0-0fea9ca7f682"+
		".json", "en", "America/Montevideo")
	http.HandleFunc("/", RequestHandler)
	fmt.Println("Started Listening...")
	http.ListenAndServe(":5000", nil)
}

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// POST method, recieves json to parse.
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		}
		type inboundMessage struct {
			Message string
		}
		var m inboundMessage
		err = json.Unmarshal(body, &m)
		if err != nil {
			panic(err)
		}

		// Use NLP.
		response := dfp.ProcessNLP(m.Message, "testUser")
		fmt.Printf("%#v", response)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
