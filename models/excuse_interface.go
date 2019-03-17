package models

// Excuse the interface for excuses of the excuse
type Excuse interface {
	GetKey() string
	// GetParams() map[string]string
	GetDoc() string
	BuildText(request Input) string
}

// type Param struct {
// }
