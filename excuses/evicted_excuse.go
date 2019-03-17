package excuses

import "github.com/kyleroush/eaas/models"

func getEvicted() models.SimpleExcuse {
	return models.SimpleExcuse{
		Key:     "evicted",
		Message: "We cant use my house I am getting evicted.",
	}
}
