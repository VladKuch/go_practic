package model

type Search struct {
	ID         int
	SearchType SearchType
	DateFrom   string
	DateTo     string
}

type SearchType struct {
	ID   int
	Type string
}
