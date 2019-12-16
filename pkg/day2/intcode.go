package day2

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func findNounVerb(expected int) (noun, verb int, err error) {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			result, err := runWithNounVerb(i, j)
			if err != nil {
				return i, j, err
			}
			if result == expected {
				return i, j, nil
			}
		}
	}
	return -1, -1, fmt.Errorf("No solution found")
}

func runWithNounVerb(noun, verb int) (int, error) {
	data, err := readInput("../../resources/day2/input.txt")
	if err != nil {
		return -1, err
	}
	data[1] = noun
	data[2] = verb
	return run(data)
}

func run(data []int) (int, error) {
	for i := 0; i < len(data); i += 4 {
		opcode := data[i]
		if opcode == 99 {
			return data[0], nil
		}
		arg1 := data[data[i+1]]
		arg2 := data[data[i+2]]
		var err error
		data[data[i+3]], err = compute(opcode, arg1, arg2)
		if err != nil {
			return -1, err
		}
	}
	return data[0], nil
}

func compute(opcode, arg1, arg2 int) (int, error) {
	if opcode == 1 {
		return arg1 + arg2, nil
	} else if opcode == 2 {
		return arg1 * arg2, nil
	} else if opcode == 99 {
		return -1, fmt.Errorf("Called compute on terminate opcode: %v", opcode)
	} else {
		return -1, fmt.Errorf("Unknown opcode: %v", opcode)
	}
}

func readInput(path string) ([]int, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	texts := strings.Split(string(content), ",")
	rval := make([]int, len(texts))
	for i := 0; i < len(texts); i++ {
		num, err := strconv.Atoi(strings.TrimSpace(texts[i]))
		if err != nil {
			return nil, err
		}
		rval[i] = num
	}
	return rval, nil
}
