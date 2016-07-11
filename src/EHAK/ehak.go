package EHAK

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

// Unit contains the structure of the EHAK classificator
type Unit struct {
	Code         string
	Name         string
	OtherName    string
	TypeCode     string
	TypeName     string
	BoroughCode  string
	BoroughName  string
	ProvinceCode string
	Province     string
}

// Load loads the EHAK classificator into memory from the file provided.
// Yes, we should be smarter than that
func Load(file string) (ehak *[]string, err error) {
	var content bool
	var e []string
	if _, fErr := os.Stat(file); os.IsNotExist(fErr) {
		err = fErr
		return
	}

	f, err := os.Open(file)
	defer f.Close()

	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(f)

	// Skip the lines we do not care about
	content = false
	for !content {
		s.Scan()
		v := strings.Split(s.Text(), "\t")

		// There will be some header lines, ignore those
		if len(v) == 9 {
			// The first column should be a string of digits
			// We do this to ignore the actual column headers
			if isValidCode(v[0]) {
				e = append(e, s.Text())
				content = true
			}
		}
	}

	for s.Scan() {
		e = append(e, s.Text())
	}

	return &e, nil
}

func isValidCode(s string) bool {
	b, err := regexp.Match(`^[0-9]+$`, []byte(s))
	return (err == nil && b)
}

// GetUnitByName returns a first matching EHAK unit by name
func GetUnitByName(name string, ehak *[]string) *Unit {
	return getUnitByColumn(name, 1, ehak)
}

// GetUnitByCode returns a EHAK unit by code
func GetUnitByCode(code string, ehak *[]string) *Unit {
	return getUnitByColumn(code, 0, ehak)
}

// GetUnitByArray returns a EHAK unit trying to match the hierarchy to the array
func GetUnitByArray(pNames []string, ehak *[]string) *Unit {
	var nList []string
	var bList []string
	var pList []string

	var names []string
	
	if len(pNames) == 0 {
		return nil
	}
	
	// Clean some stupid remarks and things from the input
	vallasisene := regexp.MustCompile(" \\(vallasisene\\)")
	remark := regexp.MustCompile("\\*")
	for _, n := range pNames{
		s := vallasisene.ReplaceAllString(n,"")
		names = append(names, remark.ReplaceAllString(s,""))
	}
	
	// Filter out exact name matches
	for _, line := range *ehak {
		v := strings.Split(line, "\t")
		if v[1] == names[len(names)-1] {
			nList = append(nList, line)
		}
	}
	// Nothing was found
	if len(nList) == 0 {
		return nil
	}

	// This is our current best guess to the match
	bestGuess := makeUnit(strings.Split(nList[0], "\t"))

	// Regardless of the number of matches, there is no point in looking further if
	// only one name was provided as input. Also return the best guess if it's the only one

	if len(names) == 1 || len(nList) == 1 {
		return bestGuess
	}

	// We found more than one match, let's look at the second array item
	// by iterating over the results of the previous filter
	for _, line := range nList {
		v := strings.Split(line, "\t")
		if v[6] == names[len(names)-2] {
			bList = append(bList, line)
		}
	}

	// So filtering by second term did not change our guess
	if len(bList) == 0 {
		return bestGuess
	}

	// The first element of the new list is our current best guess
	bestGuess = makeUnit(strings.Split(bList[0], "\t"))

	// No point in lookin further (no more terms or exactly one match was found)
	if len(names) == 2 || len(bList) == 1 {
		return bestGuess
	}

	// OK Still going. Let's look at the province level
	for _, line := range bList {
		v := strings.Split(line, "\t")
		if v[6] == names[len(names)-3] {
			pList = append(pList, line)
		}
	}

	// Filtering yielded nothing, return the previous best guess
	if len(pList) == 0 {
		return bestGuess
	}

	// There might me more than one but we'll return the first element of the filtered list
	return makeUnit(strings.Split(pList[0], "\t"))
}

func getUnitByColumn(s string, c int, ehak *[]string) *Unit {
	for _, line := range *ehak {
		// Regexp matching would probably be more efficient
		v := strings.Split(line, "\t")
		if v[c] == s {
			return makeUnit(v)
		}
	}
	return nil
}

func makeUnit(v []string) *Unit {
	return &Unit{
		Code:         v[0],
		Name:         v[1],
		OtherName:    v[2],
		TypeCode:     v[3],
		TypeName:     v[4],
		BoroughCode:  v[5],
		BoroughName:  v[6],
		ProvinceCode: v[7],
		Province:     v[8],
	}
}
// Elative returns Estonian elative of a string. Only meant to work in the EHAK context!
func Elative(input string) string{
	vowels := "AEIOUÕÄÖÜ"
	last := strings.ToUpper(string(input[len(input) - 1]))
	if strings.Contains(vowels, last){
		return input + "s"
	}

	if strings.HasSuffix(strings.ToLower(input),"vald"){
		return strings.Replace(input, "vald", "vallas", 1)
	}

	if last == "N"{
		return input + "as"
	}

	if last == "K"{
		return input + "us"
	}

	return input
}