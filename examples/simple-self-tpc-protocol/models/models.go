package models

// FixSizeModel - for translate to bytes
type FixSizeModel struct {
	B     byte
	Int64 int64
}

// JSONModel - for serialize to JSON
type JSONModel struct {
	Name string
}
