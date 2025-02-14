package gamelogic

import (
	"errors"
	"fmt"
	"strconv"
)

func (gs *GameState) CommandSpam(words []string) (int, error) {
	if len(words) < 2 {
		return -1, errors.New("usage: spam <count>")
	}

	count, err := strconv.Atoi(words[1])
	if err != nil {
		return -1, fmt.Errorf("error: %s is not a valid count", words[1])
	}

	return count, nil
}
