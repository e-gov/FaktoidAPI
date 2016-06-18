package rahvafakt

import (
	"encoding/json"
	"faktoid"
	"fmt"
	"math/rand"
	"net/http"
	"time"
	"EHAK"
)

// PopulationFakt is the implementation of the Faktoid interface
// That is able to parse the output of Estonian Board of Statistics
// output file and cross-reference it with the EHAK classificator
type PopulationFakt struct {
	ehak *[]string
	pop  *map[string]population
}

// The files to use. TODO: make these configurable in a standard fashion
var ehakF = "EHAK2015v1.txt"
var dataF = "RV0241_utf.csv"
var rnd *rand.Rand

// GetOne implements returning one random population fact
func (fakt *PopulationFakt) GetOne() *faktoid.Faktoid {
	var p population

	i := rnd.Intn(len(*fakt.pop))
	thatc := 0
	// AFAIK the simplest way to jump to a random spot of a map
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

// GetOneFiltered returns one factoid interpreting the filter as a EHAK code
func (fakt *PopulationFakt) GetOneFiltered(filter string) *faktoid.Faktoid {

	p := (*fakt.pop)[filter]
	u := EHAK.GetUnitByCode(p.EhakCode, fakt.ehak)

	return getFakt(u.Name, p.Men, p.Women)
}

// WriteData writes the entire population dataset to the writer
func (fakt *PopulationFakt) WriteData(w http.ResponseWriter) {
	str, err := json.Marshal(fakt.pop)
	if err != nil {
		panic(err)
	}
	w.Write([]byte(str))
}

// Init loads the EHAK classificator and the population file
// Also, a new source of randomness is created
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
