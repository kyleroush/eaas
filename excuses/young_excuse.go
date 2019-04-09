package excuses

import "github.com/kyleroush/eaas/models"

func getYoung() models.SimpleExcuse {
	return models.SimpleExcuse{
		Key:     "young",
		Message: "But I'm too young!!!",
	}
}
