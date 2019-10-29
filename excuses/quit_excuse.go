package excuses

import "github.com/kyleroush/eaas/models"

func getQuit() models.SimpleExcuse {
	return models.SimpleExcuse{
		Key:     "quit",
		Message: "I Quit.",
	}
}
