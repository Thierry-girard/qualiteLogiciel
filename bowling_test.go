package bowling

import (
	"fmt"
	"testing"
)

func scoreChecker(input []Frame, bonus Frame, expectedScore int, expectedError error) error {
	score, err := GetScore(input, bonus)

	if err != expectedError && !(err != nil && expectedError != nil && err.Error() == expectedError.Error()) {
		return fmt.Errorf("Score error : %+v, expected %+v", err, expectedError)
	}
	if score != expectedScore {
		return fmt.Errorf("Score : %d, expected %d", score, expectedScore)
	}
	return nil
}

func TestScore(t *testing.T) {
	input := []Frame{{0, 2}, {5, 1}, {7, 2}, {8, 1}, {7, 2}, {8, 1}, {7, 2}, {8, 1}, {7, 2}, {8, 1}}
    var bonus Frame;
	expected := 80
	if err := scoreChecker(input, bonus, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}  
}

func TestSizeInf(t *testing.T) {
	input := []Frame{{7, 5}, {13, 2}}
    var bonus Frame;
	expected := 0
	expectedError := fmt.Errorf("Game Size < 10")
	if err := scoreChecker(input, bonus, expected, expectedError); err != nil {
		t.Fatalf("%+v\n", err)
	}
	input = []Frame{{0, 2}, {5, 1}, {7, 2}, {8, 1}, {7, 2}, {8, 1}, {7, 2}, {8, 1}, {7, 2}, {8, 1}}
	expected = 80
	expectedError = nil
	if err := scoreChecker(input, bonus, expected, expectedError); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

func TestSizeSup(t *testing.T) {
    input := []Frame{{0, 2}, {5, 1}, {7, 5}, {13, 2}, {7, 5}, {13, 2}, {7, 5}, {13, 2}, {7, 5}, {7, 2}, {2, 4}}
	var bonus Frame;
    expected := 0
	expectedError := fmt.Errorf("Game Size > 10")
	if err := scoreChecker(input, bonus, expected, expectedError); err != nil {
		t.Fatalf("%+v\n", err)
	}
    input = []Frame{{0, 2}, {5, 1}, {7, 2}, {8, 1}, {7, 2}, {8, 1}, {7, 2}, {8, 1}, {7, 2}, {8, 1}}
	expected = 80
	expectedError = nil
	if err := scoreChecker(input, bonus, expected, expectedError); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

func TestTuplePositifOuNul(t *testing.T) {
	input := []Frame{{0, -2}, {5, 1}, {7, 2}, {8, 2}, {7, 2}, {8, 2}, {7, 2}, {8, 2}, {7, 2}, {8, 1}}
    var bonus Frame;
	expected := 0
	expectedError := fmt.Errorf("Negative value")
	if err := scoreChecker(input, bonus, expected, expectedError); err != nil {
		t.Fatalf("%+v\n", err)
	}
	input = []Frame{{0, 2}, {5, 1}, {7, 2}, {8, 1}, {7, 2}, {8, 1}, {7, 2}, {8, 1}, {7, 2}, {8, 1}}
	expected = 80
	expectedError = nil
	if err := scoreChecker(input, bonus, expected, expectedError); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

func TestScoreFrame(t *testing.T) {
	input := []Frame{{9, 2}, {5, 1}, {7, 5}, {13, 2}, {7, 5}, {13, 2}, {7, 5}, {13, 2}, {7, 5}, {7, 2}}
    var bonus Frame;
    expected := 0
	expectedError := fmt.Errorf("Value Error : Frame > 10")
	if err := scoreChecker(input, bonus, expected, expectedError); err != nil {
		t.Fatalf("%+v\n", err)
	}
	input = []Frame{{0, 2}, {5, 1}, {7, 3}, {3, 2}, {7, 2}, {3, 2}, {7, 1}, {3, 2}, {4, 5}, {3, 2}}
	expected = 67
	expectedError = nil
	if err := scoreChecker(input, bonus, expected, expectedError); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

func TestScoreSpare(t *testing.T) {
	input := []Frame{{0, 2}, {5, 1}, {7, 1}, {3, 2}, {7, 1}, {3, 2}, {7, 1}, {3, 2}, {4, 5}, {3, 7}}
    bonus := Frame{7,0}
	expected := 66 + bonus.firstThrow
	var expectedError error;
	expectedError = nil
	if err := scoreChecker(input, bonus, expected, expectedError); err != nil {
		t.Fatalf("%+v\n", err)
	}
	input = []Frame{{0, 2}, {5, 1}, {7, 1}, {3, 2}, {7, 1}, {3, 2}, {7, 1}, {3, 2}, {4, 5}, {3, 6}}
	expected = 65
	expectedError = nil
	if err := scoreChecker(input, bonus, expected, expectedError); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

func TestScoreStrike(t *testing.T) {
	input := []Frame{{0, 2}, {5, 1}, {7, 1}, {3, 2}, {7, 1}, {3, 2}, {7, 1}, {3, 2}, {4, 5}, {10, 0}}
    bonus := Frame{7,3}
	expected := 66 + bonus.firstThrow + bonus.secondThrow
	var expectedError error;
	expectedError = nil
	if err := scoreChecker(input, bonus, expected, expectedError); err != nil {
		t.Fatalf("%+v\n", err)
	}
	input = []Frame{{0, 2}, {5, 1}, {7, 1}, {3, 2}, {7, 1}, {3, 2}, {7, 1}, {3, 2}, {4, 5}, {8, 1}}
	expected = 65
	expectedError = nil
	if err := scoreChecker(input, bonus, expected, expectedError); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

func TestBonusEmpty(t *testing.T) {
	input := []Frame{{0, 2}, {5, 1}, {7, 1}, {3, 2}, {7, 1}, {3, 2}, {7, 1}, {3, 2}, {4, 5}, {10, 0}}
    var bonus Frame;
	expected := 0
	var expectedError error;
	expectedError = fmt.Errorf("Bonus is empty")
	if err := scoreChecker(input, bonus, expected, expectedError); err != nil {
		t.Fatalf("%+v\n", err)
	}
	input = []Frame{{0, 2}, {5, 1}, {7, 1}, {3, 2}, {7, 1}, {3, 2}, {7, 1}, {3, 2}, {4, 5}, {10, 0}}
    bonus = Frame{1,2}
	expected = 66 + bonus.firstThrow + bonus.secondThrow
	expectedError = nil
	if err := scoreChecker(input, bonus, expected, expectedError); err != nil {
		t.Fatalf("%+v\n", err)
	}
}