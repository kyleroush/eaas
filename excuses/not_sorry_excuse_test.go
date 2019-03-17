package excuses

import (
	"testing"

	"github.com/kyleroush/eaas/models"
	"github.com/stretchr/testify/assert"
)

func Test_getNotSorry(t *testing.T) {

	expected := models.SimpleExcuse{
		Key:     "notSorry",
		Message: "not sorry I just don't wanna.",
	}

	actual := getNotSorry()

	assert.Equal(t, actual.Key, expected.Key)
	assert.Equal(t, actual.Message, expected.Message)
}

func Test_MapExcuses_has_notSorry(t *testing.T) {

	actual := MapExcuses()[getNotSorry().Key].(models.SimpleExcuse)

	assert.Equal(t, actual, getNotSorry())
}
