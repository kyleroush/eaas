package excuses

import "github.com/kyleroush/eaas/models"

func getToilet() models.SimpleExcuse {
	return models.SimpleExcuse{
		Key:     "toilet",
		Message: "I won't be there because I can't get off the toilet.",
	}
}
