package rahvafakt

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"faktoid"
)

var thisF faktoid.Fakt
func sendHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}


func GetFaktoid(w http.ResponseWriter, r *http.Request) {
	f := thisF.GetOne()
	returnFaktoid(f, w)
}

func GetFilteredFaktoid(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	f := thisF.GetOneFiltered(vars["filter"])
	returnFaktoid(f, w)
}

func returnFaktoid(f *faktoid.Faktoid, w http.ResponseWriter){
	w.WriteHeader(http.StatusOK)
	if f != nil{
		if err := json.NewEncoder(w).Encode(f); err != nil {
			log.Error("Error encoding the faktoid")
			panic(err)
		}		
	}

}

func InitFakt(someF faktoid.Fakt){
	thisF = someF
}