package kutsefakt

import(
	"faktoid"
	"net/http"
	logging "github.com/op/go-logging"
	"time"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"encoding/csv"
	"io"
	"strconv"
)

var log = logging.MustGetLogger("Kutsefakt")

var kutseF = "kutsed.csv"
var rnd *rand.Rand

type KutseFakt struct {
	kutsed [][]string
}

func (fakt *KutseFakt)GetOne() *faktoid.Faktoid{
	return getFakt(fakt.kutsed[rnd.Intn(len((*fakt).kutsed))])
}

func (fakt *KutseFakt)GetMeta() *faktoid.Meta{
	m := faktoid.Meta{
		Source: "http://www.kutsekoda.ee/et/kutseregister/kutsetunnistused/statistika",
		Updated: "29/06/2016",
	}
	return &m
}


func (fakt *KutseFakt)GetOneFiltered(f string) *faktoid.Faktoid{
	n := faktoid.Faktoid{
		Language: "EST",
		Content: "Andmed puuduvad",
	}
	return &n
}

func (fakt *KutseFakt) Init() {
	fakt.kutsed = load(kutseF)
	rnd = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
}


func (fakt *KutseFakt) WriteData(w http.ResponseWriter) {
	str, err := json.Marshal(fakt.kutsed)
	if err != nil {
		panic(err)
	}
	w.Write([]byte(str))
}

func getFakt(r []string) *faktoid.Faktoid{
	var f faktoid.Faktoid
	i, err := strconv.Atoi(r[1])
	if err != nil || i >1 {
		f = faktoid.Faktoid{
			Language: "EST",
			Content:  fmt.Sprintf(`Eestis on praegu %s ametimeest ja -naist tiitliga "%s"`,
				r[1],
				r[0]),
		}
	}else{
		f = faktoid.Faktoid{
			Language: "EST",
			Content:  fmt.Sprintf(`Eestis on praegu üks ametimees või -naine tiitliga "%s"`, r[0]),
		}

	}
	return &f
}

func load(fname string) [][]string{
	var d [][]string
	f, _ := os.Open(fname)
	defer f.Close()

	rdr := csv.NewReader(f)
	for{
		r, err := rdr.Read()
		if err == io.EOF {
			break
		}
		d = append(d, r)
	}
	return d
}