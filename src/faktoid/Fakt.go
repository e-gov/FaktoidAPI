package faktoid

import(
	"net/http"
)

// Fakt is the interface that all fact providers must implement
type Fakt interface{
	// GetOne returns one random factoid from the dataset
	GetOne() *Faktoid
	// GetOneFiltered returns a random factoid matching the filter.
	// Semantics of the filter is determined by the implementation
	GetOneFiltered(string) *Faktoid
	// WriteData writes the entire raw dataset to the writer as a json string
	WriteData(http.ResponseWriter)
	GetMeta() *Meta
	// Init initialises the fact implementation
	Init()
}

type Meta struct {
	Updated string `json:"uuendatud"`
	Source 	string `json:"allikas"`
}