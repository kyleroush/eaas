package excuses

import (
	"testing"

	"github.com/kyleroush/eaas/models"
	"github.com/stretchr/testify/assert"
)

func Test_getYoung(t *testing.T) {

	expected := models.SimpleExcuse{
		Key:     "young",
		Message: "But I'm too young!!!",
	}

	actual := getWelp()

	assert.Equal(t, actual.Key, expected.Key)
	assert.Equal(t, actual.Message, expected.Message)
}

func Test_MapExcuses_has_welp(t *testing.T) {

	actual := MapExcuses()[getYoung().Key].(models.SimpleExcuse)

	assert.Equal(t, actual, getYoung())
}
