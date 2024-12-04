package pirates

import (
	"errors"
	"strings"
)

type Pirate struct {
	Name  string
	Prime string
	Img   string
}

func New(name, prime, img string) (*Pirate, error) {

	if name != strings.ToUpper(name) { // Vérification que le nom soit en majuscules
		return nil, errors.New("le nom du pirate doit être en majuscules")
	}
	if prime == "" { // Vérification que la prime n'est pas vide
		return nil, errors.New("la prime ne peut pas être vide")
	}
	if img == "" { // Vérification que l'image n'est pas vide
		return nil, errors.New("l'image ne peut pas être vide")
	}

	pirate := &Pirate{name, prime, img}

	return pirate, nil
}
