package excuses

import (
	"testing"

	"github.com/kyleroush/eaas/models"
	"github.com/stretchr/testify/assert"
)

func Test_getWelp(t *testing.T) {

	expected := models.SimpleExcuse{
		Key:     "welp",
		Message: "Welp I guess I cant make it.",
	}

	actual := getWelp()

	assert.Equal(t, actual.Key, expected.Key)
	assert.Equal(t, actual.Message, expected.Message)
}

func Test_MapExcuses_has_welp(t *testing.T) {

	actual := MapExcuses()[getWelp().Key].(models.SimpleExcuse)

	assert.Equal(t, actual, getWelp())
}
