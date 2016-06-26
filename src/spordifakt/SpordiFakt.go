package spordifakt

import(
	"faktoid"
	"net/http"
	logging "github.com/op/go-logging"
	"EHAK"
	"strings"
	"time"
	"encoding/json"
	"fmt"
	"math/rand"
)

var log = logging.MustGetLogger("SpordiSvc")

// The files to use. TODO: make these configurable in a standard fashion
var ehakF = "EHAK2015v1.txt"
var rnd *rand.Rand

type SpordiFakt struct {
	ehak *[]string
	dta []HResponse
	timestamp time.Time
}

func (fakt *SpordiFakt)GetOne() *faktoid.Faktoid{
	fmt.Println(len((*fakt).dta))
	return getFakt(&(*fakt).dta[rnd.Intn(len((*fakt).dta))])
}

func (fakt *SpordiFakt)GetMeta() *faktoid.Meta{
	m := faktoid.Meta{
		Source: baseURL,
		Updated: (*fakt).timestamp.Format(time.ANSIC),
	}
	return &m
}


func (fakt *SpordiFakt)GetOneFiltered(f string) *faktoid.Faktoid{
	for _, d := range fakt.dta{
		if d.EHAK.Kood == f{
			return getFakt(&d)
		}
	}

	n := faktoid.Faktoid{
		Language: "EST",
		Content: "Andmed puuduvad",
	}
	return &n
}

func (fakt *SpordiFakt) Init() {
	var err error
	log.Debug("Loading the EHAK dataset")
	if fakt.ehak, err = EHAK.Load(ehakF); err != nil {
		panic("Failed to load the EHAK dataset")
	}
	log.Debug("Done loading EHAK")

	for i, line := range *fakt.ehak {
		log.Debugf("Loading %d of %d", i, len(*fakt.ehak))
		v := strings.Split(line, "\t")
		d := load(v[0])
		if len((*d).Alad) > 0{
			fakt.dta = append(fakt.dta, *d)
		}
		if i == 50{
			break
		}
	}
	fakt.timestamp = time.Now()
	rnd = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
}


func (fakt *SpordiFakt) WriteData(w http.ResponseWriter) {
	str, err := json.Marshal(fakt.dta)
	if err != nil {
		panic(err)
	}
	w.Write([]byte(str))
}

func getFakt(r *HResponse) *faktoid.Faktoid{
	i := rnd.Intn(len((*r).Alad))
	f := faktoid.Faktoid{
		Language: "EST",
		Content:  fmt.Sprintf("%s asulas %s tegeleb alaga %s %s inimest, neist %s noored",
			(*r).EHAK.Maakond[0],
			(*r).EHAK.Kov[0] + ", " + (*r).EHAK.Asutus[0],
			(*r).Alad[i].Ala,
			(*r).Alad[i].Kokku,
			(*r).Alad[i].Noori),
	}
	return &f
}