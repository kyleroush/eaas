package excuses

import "github.com/kyleroush/eaas/models"

// ListExcuses list all of the excuses
func ListExcuses() []models.Excuse {
	return []models.Excuse{
		getDog(),
		getNotSorry(),
		getPosion(),
		getToilet(),
		getEvicted(),
		getContagious(),
		getWelp(),
		getYoung(),
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
