package models

const toName = "to"
const fromName = "from"

// SimpleExcuse This is a simple excuse where the message is simple string
type SimpleExcuse struct {
	Key     string
	Message string
}

// BuildText builds the message of the excuse
func (excuse SimpleExcuse) BuildText(request Input) string {
	to := request.Inputs[toName]
	from := request.Inputs[fromName]

	message := ""
	if to != "" {
		message += "Dear " + to + ", "
	}
	message += excuse.Message

	if from != "" {
		message += " Sincerly " + from
	}

	return message
}

// GetKey returns the key of the simple excuse
func (excuse SimpleExcuse) GetKey() string {
	return excuse.Key
}

// GetParams returns the map of the params
func (excuse SimpleExcuse) GetParams() []Param {

	return []Param{
		Param{
			Required: false,
			Name:     toName,
		},
		Param{
			Required: false,
			Name:     fromName,
		},
	}
}

// GetDoc returns the docs of the simple excuse
func (excuse SimpleExcuse) GetDoc() string {
	input := Input{
		Inputs: map[string]string{
			toName:   ":" + toName,
			fromName: ":" + fromName,
		},
	}

	return "Creates a message in the form of " +
		excuse.BuildText(input) +
		" with the 'Dear :to,' and 'Sincerly :from' being optional only added when to or from are passed in."
}
