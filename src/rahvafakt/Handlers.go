package rahvafakt

import (
	"net/http"
	"faktoid"
	"encoding/json"
//	"strings"
)

func sendHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}


func GetFaktoid(w http.ResponseWriter, r *http.Request) {
	f := faktoid.Faktoid{
		Language: "EST",
		Content: "Juhuslik fakt",
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(f); err != nil {
		log.Error("Error encoding the faktoid")
		panic(err)
	}

}
