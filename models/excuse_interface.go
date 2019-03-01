package models

// Excuse the interface for excuses of the excuse
type Excuse interface {
	GetKey() string
	BuildText(request Input) string
}
