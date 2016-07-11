package arifakt

import (
	"EHAK"
	"faktoid"
	"math/rand"
	"net/http"

	logging "github.com/op/go-logging"
	"os"
	"encoding/csv"
	"io"
	"fmt"
	"time"
)

var log = logging.MustGetLogger("SpordiSvc")

// The files to use. TODO: make these configurable in a standard fashion
var ehakF = "EHAK2015v1.txt"
var ariF = "ettevotja_rekvisiidid_2016-07-06.csv"
var rnd *rand.Rand

type AriFakt struct {
	ehak *[]string
	data [][]string
	count map[string]int
}

func (fakt *AriFakt) GetOne() *faktoid.Faktoid {
	var i int
	n := rnd.Intn(len((*fakt).count))

	for ehakC, count := range (*fakt).count{

		if i == n{
			return getFakt(count, ehakC, (*fakt).ehak)
		}
		i++
	}
	return nil
}

func (fakt *AriFakt) GetMeta() *faktoid.Meta {
	m := faktoid.Meta{
		Source:  "http://avaandmed.rik.ee/andmed/ARIREGISTER/ariregister_csv.zip",
		Updated: "11/07/2016",
	}
	return &m
}

func (fakt *AriFakt) GetOneFiltered(f string) *faktoid.Faktoid {
	n := faktoid.Faktoid{
		Language: "EST",
		Content:  "Andmed puuduvad",
	}
	return &n
}

func (fakt *AriFakt) Init() {
	var err error
	log.Debug("Loading the EHAK dataset")
	if fakt.ehak, err = EHAK.Load(ehakF); err != nil {
		panic("Failed to load the EHAK dataset")
	}
	log.Debug("Done loading EHAK")

	log.Debug("Loading business data")
	fakt.data = load(ariF)
	log.Debug("Counting instances")
	fakt.count = count(fakt.data)
	log.Debug("Done loading")

	rnd = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
}

func (fakt *AriFakt) WriteData(w http.ResponseWriter) {
}

func load(fname string) [][]string{
	var d [][]string
	f, _ := os.Open(fname)

	defer f.Close()

	rdr := csv.NewReader(f)
	rdr.Comma = ';'
	for{
		r, err := rdr.Read()
		if err == io.EOF {
			break
		}
		d = append(d, r)
	}
	return d
}

func count(source [][]string) map[string]int{

	destination := make(map[string]int)
	for _, l := range source{
		status := l [3]
		ehak := l[7]
		if status == "R"{
			destination[ehak] = destination[ehak] + 1
		}
	}

	return destination
}

func getFakt(c int, code string, ehak *[]string) *faktoid.Faktoid{
	var f faktoid.Faktoid
	var countS  string

	if c == 1{
		countS = "on 1 ettevõte"
	}else{
		if c == 0{
			countS = "ettevõtlust nagu polekski"
		}else{
			countS = fmt.Sprintf("on %d ettevõtet", c)
		}
	}

	u := EHAK.GetUnitByCode(code, ehak)
	if u == nil{
		f = faktoid.Faktoid{
			Language: "EST",
			Content:  "Asumis EHAK koodiga " + code + " " + countS,
		}

	}else{
		f = faktoid.Faktoid{
			Language: "EST",
			Content: EHAK.Elative(EHAK.GetUnitByCode(code, ehak).Name) + " " + countS,
		}
	}
	return &f
}