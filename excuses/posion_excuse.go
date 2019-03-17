package excuses

import "github.com/kyleroush/eaas/models"

func getPosion() models.SimpleExcuse {
	return models.SimpleExcuse{
		Key:     "posion",
		Message: "I cant make cause the food I eat is trying to kill me.",
	}
}
