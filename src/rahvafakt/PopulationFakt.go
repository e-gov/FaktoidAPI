package rahvafakt

import (
	"ehak"
	"encoding/json"
	"faktoid"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type PopulationFakt struct {
	ehak *[]string
	pop  *map[string]Population
}

var ehakF = "EHAK2015v1.txt"
var dataF = "RV0241_utf.csv"
var rnd *rand.Rand

func (fakt *PopulationFakt) GetOne() *faktoid.Faktoid {
	var p Population

	i := rnd.Intn(len(*fakt.pop))
	thatc := 0
	for that := range *fakt.pop {
		if thatc == i {
			p = (*fakt.pop)[that]
			break
		}
		thatc++
	}
	u := EHAK.GetUnitByCode(p.EhakCode, fakt.ehak)

	return getFakt(u.Name, p.Men, p.Women)
}

func (fakt *PopulationFakt) GetOneFiltered(filter string) *faktoid.Faktoid {

	p := (*fakt.pop)[filter]
	u := EHAK.GetUnitByCode(p.EhakCode, fakt.ehak)

	return getFakt(u.Name, p.Men, p.Women)
}

func (fakt *PopulationFakt) WriteData(w http.ResponseWriter) {
	str, err := json.Marshal(fakt.pop)
	if err != nil {
		panic(err)
	}
	w.Write([]byte(str))
}

func (fakt *PopulationFakt) Init() {
	var err error

	if fakt.ehak, err = EHAK.Load(ehakF); err != nil {
		panic("Failed to load the EHAK dataset")
	}

	fakt.pop = LoadData(dataF, fakt.ehak)
	rnd = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
}

func getFakt(name string, m int, w int) *faktoid.Faktoid {
	f := faktoid.Faktoid{
		Language: "EST",
		Content:  fmt.Sprintf("Asulas %s elab statistika andmetel %d meest ja %d naist", name, m, w),
	}
	return &f
}
