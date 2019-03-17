package models

// Excuse the interface for excuses of the excuse
type Excuse interface {
	GetKey() string
	GetParams() []Param
	GetDoc() string
	BuildText(request Input) string
}

// Param the model to describe the params of an excuse
type Param struct {
	Required bool   `json:"required"`
	Name     string `json:"name"`
}
