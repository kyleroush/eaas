package models

// SimpleExcuse This is a simple excuse where the message is simple string
type SimpleExcuse struct {
	Key     string
	Message string
}

// BuildText builds the message of the excuse
func (excuse SimpleExcuse) BuildText(request Input) string {
	to := request.Inputs["to"]
	from := request.Inputs["from"]

	message := ""
	if to != "" {
		message += "Dear " + to + ", "
	}
	message += excuse.Message

	if from != "" {
		message += " Sincerly " + from
	}

	return excuse.Message
}

// GetKey returns the key of the simple excuse
func (excuse SimpleExcuse) GetKey() string {
	return excuse.Key
}
