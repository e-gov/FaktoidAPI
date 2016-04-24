package faktoid

import(
	"net/http"
)

type Fakt interface{
	GetOne() *Faktoid
	GetOneFiltered(string) *Faktoid
	WriteData(http.ResponseWriter)
}