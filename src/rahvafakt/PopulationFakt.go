package rahvafakt

import(
	"faktoid"
	"net/http"
	"encoding/json"
)
type PopulationFakt struct{
	
}

func (fakt *PopulationFakt)GetOne() *faktoid.Faktoid{
	f := faktoid.Faktoid{
		Language: "EST",
		Content: "Juhuslik fakt",
	}
	return &f
}


func (fakt *PopulationFakt)GetOneFiltered(filter string) *faktoid.Faktoid{
	f := faktoid.Faktoid{
		Language: "EST",
		Content: "Juhuslik fakt",
	}
	return &f
}

func (fakt *PopulationFakt)WriteData(w http.ResponseWriter){
	fs := map[string]interface{}{
		"Fakt1":"First",
		"Fakt2":"Second",
	}
		
	str, err := json.Marshal(fs)
	if err != nil{
		panic(err)
	}
	w.Write([]byte(str))
}