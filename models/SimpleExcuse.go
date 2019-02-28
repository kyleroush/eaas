package models

// SimpleExcuse This is a simple excuse where the message is simple string
type SimpleExcuse struct {
	key string
	message string
}

func (excuse SimpleExcuse) buildMessage() string {
	return excuse.message
}

func (excuse SimpleExcuse) getKey() string {
	return excuse.key
}
