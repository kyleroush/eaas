package main

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
