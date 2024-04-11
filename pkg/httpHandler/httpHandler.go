package httpHandler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

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
		fmt.Println(m.Message)
	}
}
