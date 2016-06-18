package rahvafakt

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"faktoid"
)

var thisF faktoid.Fakt

// The handlers are all abstract. TODO: refactor moving standard code to the faktoid package

// GetFaktoid handles the case of returning one random factoid
func GetFaktoid(w http.ResponseWriter, r *http.Request) {
	f := thisF.GetOne()
	returnFaktoid(f, w)
}

// GetFilteredFaktoid implements returning of one filtered factoid
func GetFilteredFaktoid(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	f := thisF.GetOneFiltered(vars["filter"])
	returnFaktoid(f, w)
}

// GetData writes the raw dataset to the output stream as a json string
func GetData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	thisF.WriteData(w)
}

func returnFaktoid(f *faktoid.Faktoid, w http.ResponseWriter){
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if f != nil{
		if err := json.NewEncoder(w).Encode(f); err != nil {
			log.Error("Error encoding the faktoid")
			panic(err)
		}		
	}

}

// InitFakt initializes the supplied faktoid and keeps it for handlers to use
func InitFakt(someF faktoid.Fakt){
	thisF = someF
	thisF.Init()
}