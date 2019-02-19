package main

func listExcuses() map[string]Excuse {
	excuses := []Excuse{
		getDog(),
		getNotSorry(),
		getPosion(),
		getToliet(),
		getEvicted(),
		getContagious(),
	}

	mapExcuses := map[string]Excuse {}
	for _, e := range excuses {
		mapExcuses[e.getKey()] = e
	}
	return mapExcuses
}

// Message will not be exported but is used several places
type Message struct {
	Memo string `json:"memo"`
	From string `json:"from"`
	To   string `json:"to"`
	Text string `json:"text"`
}

// Excuse the model of the excuse 
type Excuse interface {
    getKey() string
    buildMessage() string
}

func getDog() Excuse {
	return SimpleExcuse {
		key: "dog",
		message: "My dog ate my homework.",
	}
}

func getNotSorry() Excuse {
	return SimpleExcuse {
		key: "notSorry",
		message: "not sorry I just don't wanna.",
	}
}

func getPosion() Excuse {
	return SimpleExcuse {
		key: "posion",
		message: "I cant make cause the food I eat is trying to kill me.",
	}
}

func getToliet() Excuse {
	return SimpleExcuse {
		key: "toliet",
		message: "I won't be there because I can't get off the toliet.",
	}
}

func getContagious() Excuse {
	return SimpleExcuse {
		key: "contagious",
		message: "I am not going to be able to make it cause I am contagious.",
	}
}

func getEvicted() Excuse {
	return SimpleExcuse {
		key: "evicted",
		message: "We cant use my house I am getting evicted.",
	}
}