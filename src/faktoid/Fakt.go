package faktoid

type Fakt interface{
	GetOne() *Faktoid
	GetOneFiltered(string) *Faktoid
}