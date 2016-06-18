package faktoid

// Faktoid is the structure containing the abstract factoid.
// Language is the language used, content is the fact content
// See the API spec for details
type Faktoid struct{
	Language 	string	`json:"keel,omitempty"`
	Content		string 	`json:"tekst"`
}