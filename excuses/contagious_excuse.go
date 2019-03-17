package excuses

import "github.com/kyleroush/eaas/models"

func getContagious() models.SimpleExcuse {
	return models.SimpleExcuse{
		Key:     "contagious",
		Message: "I am not going to be able to make it cause I am contagious.",
	}
}
