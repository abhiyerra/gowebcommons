package render

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func RenderJson(w http.ResponseWriter, page interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	switch page.(type) {
	case []byte:
		w.Write(page.([]byte))
	default:
		b, err := json.Marshal(page)
		if err != nil {
			log.Println("error:", err)
			fmt.Fprintf(w, "")
		}

		w.Write(b)
	}
}
