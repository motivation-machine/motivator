package app

// Post is a normalized motivation object, that can be from any source
type Post struct {
	Text         string
	Picture      string
	Sayer        string
	Photographer string
	PhotoSource  string
	SourceLink   string
}
