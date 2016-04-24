package rahvafakt

import(
	"faktoid"
)
type PopulationFakt struct{
	
}

func (fakt *PopulationFakt)GetOne() *faktoid.Faktoid{
	f := faktoid.Faktoid{
		Language: "EST",
		Content: "Juhuslik fakt",
	}
	return &f
}


func (fakt *PopulationFakt)GetOneFiltered(filter string) *faktoid.Faktoid{
	f := faktoid.Faktoid{
		Language: "EST",
		Content: "Juhuslik fakt",
	}
	return &f
}