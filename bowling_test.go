package bowling

import (
	"fmt"
	"testing"
)

func scoreChecker(input []Frame, expectedScore int, expectedError error) error {
	score, err := GetScore(input)

	if err != expectedError && !(err != nil && expectedError != nil && err.Error() == expectedError.Error()) {
		return fmt.Errorf("Score error : %+v, expected %+v", err, expectedError)
	}
	if score != expectedScore {
		return fmt.Errorf("Score : %d, expected %d", score, expectedScore)
	}
	return nil
}

/*func TestNullScore(t *testing.T) {
	input := []Frame{}
	expected := 0
	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}
func TestScore(t *testing.T) {
	input := []Frame{{7, 5}, {13, -2}}
	expected := 23
	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
	input = []Frame{{0, 2}, {5, 1}}
	expected = 8
	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}*/


func TestSizeScore(t *testing.T) {
	input := []Frame{{7, 5}, {13, 2}}
	expected := 0
	expectedError := fmt.Errorf("Game Size < 10")
	if err := scoreChecker(input, expected, expectedError); err != nil {
		t.Fatalf("%+v\n", err)
	}
	input = []Frame{{0, 2}, {5, 1}, {7, 5}, {13, 2}, {7, 5}, {13, 2}, {7, 5}, {13, 2}, {7, 5}, {13, 2}}
	expected = 116
	expectedError = nil
	if err := scoreChecker(input, expected, expectedError); err != nil {
		t.Fatalf("%+v\n", err)
	}
	input = []Frame{{0, 2}, {5, 1}, {7, 5}, {13, 2}, {7, 5}, {13, 2}, {7, 5}, {13, 2}, {7, 5}, {13, 2}, {13, 2}}
	expected = 0
	expectedError = fmt.Errorf("Game Size > 10")
	if err := scoreChecker(input, expected, expectedError); err != nil {
		t.Fatalf("%+v\n", err)
	}
}


func TestScorePositif(t *testing.T) {
	input := []Frame{{0, 2}, {5, 1}, {7, 5}, {13, -2}, {7, 5}, {13, -2}, {7, 5}, {13, -2}, {7, 5}, {13, -2}}
	expected := 0
	expectedError := fmt.Errorf("Negative value")
	if err := scoreChecker(input, expected, expectedError); err != nil {
		t.Fatalf("%+v\n", err)
	}
	input = []Frame{{0, 2}, {5, 1}, {7, 5}, {13, 2}, {7, 5}, {13, 2}, {7, 5}, {13, 2}, {7, 5}, {13, 2}}
	expected = 116
	expectedError = nil
	if err := scoreChecker(input, expected, expectedError); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

func TestScoreFrame(t *testing.T) {
	input := []Frame{{0, 2}, {5, 1}, {7, 5}, {13, 2}, {7, 5}, {13, 2}, {7, 5}, {13, 2}, {7, 5}, {13, 2}}
	expected := 0
	expectedError := fmt.Errorf("Negative value")
	if err := scoreChecker(input, expected, expectedError); err != nil {
		t.Fatalf("%+v\n", err)
	}
	input = []Frame{{0, 2}, {5, 1}, {7, 3}, {3, 2}, {7, 2}, {3, 2}, {7, 1}, {3, 2}, {4, 5}, {3, 2}}
	expected = 64
	expectedError = nil
	if err := scoreChecker(input, expected, expectedError); err != nil {
		t.Fatalf("%+v\n", err)
	}
}