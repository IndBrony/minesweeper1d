package minesweeper1d

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

//CountMinesString is a string wrapper for CountMines function
//it takes a string of number separated with a comma followed by a space ", "
//it also give additional error object if CountMines returned -1
func CountMinesString(s string) (int, error) {
	//split the string into an array of numbers in string format
	stringArr := strings.Split(s, ", ")
	arr := make([]int, 0)
	//for each number, convert it into an int (return err if cant be converted)
	for _, clueString := range stringArr {
		clue, err := strconv.Atoi(clueString)
		if err != nil {
			return -1, err
		}
		arr = append(arr, clue)
	}
	//count the mines in the given clues and add error message if returned -1
	if output := CountMines(arr); output == -1 {
		return output, errors.New("clues-not-valid")
	} else {
		return output, nil
	}
}

func CountMines(arr []int) int {
	cluesSum := 0
	// hasClueBefore := false
	//minesBlock is a block of mines in the array. Opened and closed by "1" clue. ex =
	//[1][R][1]
	//[1][R][2][R][1]
	inAMinesBlock := false
	//for each clue
	for _, clue := range arr {
		if inAMinesBlock && clue == 1 {
			//if the inside a minesBlock and the clue is "1" then close it
			inAMinesBlock = false
		} else if (inAMinesBlock && clue == 0) || (!inAMinesBlock && clue == 2) {
			//if the minesBlock is not yet closed but the clue is "0" then it isn't valid
			return -1
		} else if !inAMinesBlock && clue == 1 {
			//if the not inside a minesBlock and the clue is "1" then open it
			inAMinesBlock = true
		}
		cluesSum += clue
	}
	//the number of mines is always the sum of all the clues divided by 2
	//and then round it upwards
	return int(math.Ceil(float64(cluesSum) / 2))
}
