package excuses

import (
	"testing"

	"github.com/kyleroush/eaas/models"
	"github.com/stretchr/testify/assert"
)

func Test_getPosion(t *testing.T) {

	expected := models.SimpleExcuse{
		Key:     "posion",
		Message: "I cant make cause the food I eat is trying to kill me.",
	}

	actual := getPosion()

	assert.Equal(t, actual.Key, expected.Key)
	assert.Equal(t, actual.Message, expected.Message)
}

func Test_MapExcuses_has_posion(t *testing.T) {

	actual := MapExcuses()[getPosion().Key].(models.SimpleExcuse)

	assert.Equal(t, actual, getPosion())
}
