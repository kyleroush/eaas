package excuses

import (
	"testing"

	"github.com/kyleroush/eaas/models"
	"github.com/stretchr/testify/assert"
)

func Test_getContagious(t *testing.T) {

	expected := models.SimpleExcuse{
		Key:     "contagious",
		Message: "I am not going to be able to make it cause I am contagious.",
	}

	actual := getContagious()

	assert.Equal(t, actual.Key, expected.Key)
	assert.Equal(t, actual.Message, expected.Message)
}

func Test_MapExcuses_has_contagious(t *testing.T) {

	actual := MapExcuses()[getContagious().Key].(models.SimpleExcuse)

	assert.Equal(t, actual, getContagious())
}
