package excuses

import "github.com/kyleroush/eaas/models"

func getWelp() models.SimpleExcuse {
	return models.SimpleExcuse{
		Key:     "welp",
		Message: "Welp I guess I cant make it.",
	}
}
