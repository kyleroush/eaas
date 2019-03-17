package excuses

import (
	"testing"

	"github.com/kyleroush/eaas/models"
	"github.com/stretchr/testify/assert"
)

func Test_getEvicted(t *testing.T) {

	expected := models.SimpleExcuse{
		Key:     "evicted",
		Message: "We cant use my house I am getting evicted.",
	}

	actual := getEvicted()

	assert.Equal(t, actual.Key, expected.Key)
	assert.Equal(t, actual.Message, expected.Message)
}

func Test_MapExcuses_has_evicted(t *testing.T) {

	actual := MapExcuses()[getEvicted().Key].(models.SimpleExcuse)

	assert.Equal(t, actual, getEvicted())
}
