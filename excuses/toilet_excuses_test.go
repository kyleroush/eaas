package excuses

import (
	"testing"

	"github.com/kyleroush/eaas/models"
	"github.com/stretchr/testify/assert"
)

func Test_getToilet(t *testing.T) {

	expected := models.SimpleExcuse{
		Key:     "toilet",
		Message: "I won't be there because I can't get off the toilet.",
	}

	actual := getToilet()

	assert.Equal(t, actual.Key, expected.Key)
	assert.Equal(t, actual.Message, expected.Message)
}

func Test_MapExcuses_has_toilet(t *testing.T) {

	actual := MapExcuses()[getToilet().Key].(models.SimpleExcuse)

	assert.Equal(t, actual, getToilet())
}
