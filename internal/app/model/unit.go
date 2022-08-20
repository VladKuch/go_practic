package model

type Unit struct {
	ID        int
	Type      UnitType
	UnitName  string
	CreatedAt string
}

type UnitType struct {
	ID   int
	Type string
}
