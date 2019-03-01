package excuses

import "github.com/kyleroush/eaas/models"

func getWelp() models.Excuse {
	return models.SimpleExcuse{
		Key:     "welp",
		Message: "Welp I guess I cant make it.",
	}
}
