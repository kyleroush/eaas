package excuses

import "github.com/kyleroush/eaas/models"

func getNotSorry() models.SimpleExcuse {
	return models.SimpleExcuse{
		Key:     "notSorry",
		Message: "not sorry I just don't wanna.",
	}
}
