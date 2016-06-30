package spordifakt

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
)

var baseURL = "https://www.spordiregister.ee/opendata/kov?EHAK="

type HResponse struct{
	EHAK Sehak `json:"ehak"`
	Alad []Ala `json:"harrastajad:spordiala"`
}

type Sehak struct{
	Kood string		`json:"EHAKKood"`
	Maakond []string 	`json:"maakond:maakond"`
	Kov []string		`json:"kov:kov"`
	Asutus []string		`json:"asustus:asustus"`
}

type Ala struct{
	Kood string	`json:"spordiala_kood"`
	Ala string	`json:"spordiala"`
	Kokku string	`json:"harrastajad"`
	Noori string	`json:"noored"`
}

func load(ehak string)*HResponse{
	log.Debug("Loading data for " + ehak)
	res, err := http.Get(baseURL + ehak)
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	r := new(HResponse)
	err = json.Unmarshal(body, &r)
	if err != nil{
		log.Error("Invalid response from data source")
	}
	r.EHAK.Kood = ehak
	return r
}