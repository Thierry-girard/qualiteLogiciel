package bowling

import "fmt"

type Frame struct {
	firstThrow  int
	secondThrow int
}

func GetScore(game []Frame) (int, error) {
	score := 0
	if len(game) < 10{
		return 0, fmt.Errorf("Game Size < 10")
	}
	if len(game) > 10{
		return 0, fmt.Errorf("Game Size > 10")
	}

	for i := 0; i < len(game); i++ {
		if game[i].firstThrow < 0 || game[i].secondThrow < 0{
			return 0, fmt.Errorf("Negative value")
		}
		if game[i].firstThrow + game[i].secondThrow > 10{
			return 0, fmt.Errorf("Value Error : Frame > 10")
		}
		score = score + game[i].firstThrow + game[i].secondThrow
	}
	return score, nil
}