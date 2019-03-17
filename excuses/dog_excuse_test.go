package excuses

import (
	"testing"

	"github.com/kyleroush/eaas/models"
	"github.com/stretchr/testify/assert"
)

func Test_getDog(t *testing.T) {
	expected := models.SimpleExcuse{
		Key:     "dog",
		Message: "My dog ate my homework.",
	}

	actual := getDog()

	assert.Equal(t, actual.Key, expected.Key)
	assert.Equal(t, actual.Message, expected.Message)
}

func Test_MapExcuses_has_dog(t *testing.T) {

	actual := MapExcuses()[getDog().Key].(models.SimpleExcuse)

	assert.Equal(t, actual, getDog())
}
