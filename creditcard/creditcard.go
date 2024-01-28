package creditcard

import (
	"errors"
)

type card struct {
	number string
}

func (cc card) Number() string {
	return cc.number
}

func New(number string) (card, error) {
	if number == "" {
		return card{}, errors.New("number field cannot be emtpy")
	}
	cc := card{number: number}

	return cc, nil
}
