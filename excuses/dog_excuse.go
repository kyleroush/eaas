package excuses

import "github.com/kyleroush/eaas/models"

func getDog() models.SimpleExcuse {
	return models.SimpleExcuse{
		Key:     "dog",
		Message: "My dog ate my homework.",
	}
}
