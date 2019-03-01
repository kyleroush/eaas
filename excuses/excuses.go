package excuses

import "github.com/kyleroush/eaas/models"

// ListExcuses list all of the excuses
func ListExcuses() []models.Excuse {
	return []models.Excuse{
		getDog(),
		getNotSorry(),
		getPosion(),
		getToliet(),
		getEvicted(),
		getContagious(),
		getWelp(),
	}
}

// MapExcuses returns a map of the excuses
func MapExcuses() map[string]models.Excuse {
	excuses := ListExcuses()

	mapExcuses := map[string]models.Excuse{}
	for _, e := range excuses {
		mapExcuses[e.GetKey()] = e
	}
	return mapExcuses
}

// move to there own files

func getDog() models.Excuse {
	return models.SimpleExcuse{
		Key:     "dog",
		Message: "My dog ate my homework.",
	}
}

func getNotSorry() models.Excuse {
	return models.SimpleExcuse{
		Key:     "notSorry",
		Message: "not sorry I just don't wanna.",
	}
}

func getPosion() models.Excuse {
	return models.SimpleExcuse{
		Key:     "posion",
		Message: "I cant make cause the food I eat is trying to kill me.",
	}
}

func getToliet() models.Excuse {
	return models.SimpleExcuse{
		Key:     "toliet",
		Message: "I won't be there because I can't get off the toliet.",
	}
}

func getContagious() models.Excuse {
	return models.SimpleExcuse{
		Key:     "contagious",
		Message: "I am not going to be able to make it cause I am contagious.",
	}
}

func getEvicted() models.Excuse {
	return models.SimpleExcuse{
		Key:     "evicted",
		Message: "We cant use my house I am getting evicted.",
	}
}
