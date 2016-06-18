package rahvafakt

import (
	"encoding/csv"
	"io"
	"os"
	"regexp"
	"strconv"
	"EHAK"
)

type population struct {
	EhakCode string `json:"ehak"`
	Men      int    `json:"men"`
	Women    int    `json:"women"`
}

func readPopLine(rdr *csv.Reader) *population {
	var r []string
	var p population

	// Ignore the first line, it is a sum of the next two
	r, _ = rdr.Read()

	r, _ = rdr.Read()
	p.Men, _ = strconv.Atoi(r[3])

	r, _ = rdr.Read()
	p.Women, _ = strconv.Atoi(r[3])

	return &p
}

// LoadData loads the population data from the filename supplied and merges it with the EHAK dataset
// TODO: make it call EHAK parsing directly, the knowledge is only useful in here
func LoadData(fname string, ehak *[]string) *map[string]population {
	var ps map[string]population
	var content bool
	var p *population
	var v *Stack

	if _, fErr := os.Stat(fname); os.IsNotExist(fErr) {
		return &ps
	}

	f, _ := os.Open(fname)
	defer f.Close()

	rdr := csv.NewReader(f)
	rdr.Comma = '\t'

	rdr.FieldsPerRecord = -1
	content = false
	for {
		r, err := rdr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		if r[0] == "2016" {
			content = true
			break
		}
	}

	pDots := 0
	v = new(Stack)
	dots := regexp.MustCompile("^(\\.*)")
	ps = make(map[string]population)
	if content {
		for {
			r, err := rdr.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
			}

			// we are done with the file
			if len(r) < 2 {
				break
			}

			s := dots.ReplaceAllString(r[1], "")
			c := len(r[1]) - len(s)
			p = readPopLine(rdr)
			// Ignore the lines without dots, they are a summary we do not need
			if c > 0 {
				if c <= pDots {
					v.Pop()
					if c < pDots {
						v.Pop()
					}
				}
				v.Push(s)

				u := EHAK.GetUnitByArray(*v.Content(), ehak)
				if u != nil {
					p.EhakCode = u.Code
					ps[u.Code] = *p
				}

				pDots = c
			}

		}
	}
	return &ps
}
