package day5

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	add         = 1
	mult        = 2
	read        = 3
	write       = 4
	jumpIfTrue  = 5
	jumpIfFalse = 6
	lessThan    = 7
	equals      = 8
	quit        = 99
)

type instruction struct {
	opcode    int
	modes     []int
	arguments []int
}

// Intsrc reads integers
type Intsrc interface {
	Read() (int, error)
}

// Intdest writes integers
type Intdest interface {
	Write(i int) error
}

// Run runs the provided data getting input from in and outputting to out
func Run(data []int, in Intsrc, out Intdest) error {
	ptr := 0
	for ptr < len(data) {
		inst, err := parseInstruction(ptr, data)
		if err != nil {
			return fmt.Errorf("Problem parsing at %v: %v", ptr, err)
		}
		if inst.opcode == quit {
			return nil
		}
		var jumpDest int
		jumpDest, err = execute(data, inst, in, out)
		if err != nil {
			return fmt.Errorf("Problem running at %v: %v", ptr, err)
		}
		if jumpDest > 0 {
			ptr = jumpDest
		} else {
			ptr += len(inst.arguments) + 1
		}
	}
	return fmt.Errorf("Ran out of data before termination")
}

func execute(data []int, inst *instruction, in Intsrc, out Intdest) (int, error) {
	rval := 0
	switch inst.opcode {
	case add:
		arg1 := getVal(0, inst, data)
		arg2 := getVal(1, inst, data)
		arg3 := inst.arguments[2]
		fmt.Printf("%v + %v -> %v\n", arg1, arg2, arg3)
		data[arg3] = arg1 + arg2
	case mult:
		arg1 := getVal(0, inst, data)
		arg2 := getVal(1, inst, data)
		arg3 := inst.arguments[2]
		fmt.Printf("%v * %v -> %v\n", arg1, arg2, arg3)
		data[arg3] = arg1 * arg2
	case read:
		i, err := in.Read()
		if err != nil {
			return 0, err
		}
		index := inst.arguments[0]
		fmt.Printf("read %v -> %v\n", i, index)
		data[index] = i
	case write:
		arg := getVal(0, inst, data)
		fmt.Printf("write %v\n", arg)
		err := out.Write(arg)
		if err != nil {
			return 0, err
		}
	case jumpIfTrue:
		arg1 := getVal(0, inst, data)
		arg2 := getVal(1, inst, data)
		if arg1 != 0 {
			fmt.Printf("%v != 0, jumping to %v\n", arg1, arg2)
			rval = arg2
		} else {
			fmt.Printf("not jumping\n")
		}
	case jumpIfFalse:
		arg1 := getVal(0, inst, data)
		arg2 := getVal(1, inst, data)
		if arg1 == 0 {
			fmt.Printf("%v == 0, jumping to %v\n", arg1, arg2)
			rval = arg2
		} else {
			fmt.Printf("not jumping\n")
		}
	case lessThan:
		arg1 := getVal(0, inst, data)
		arg2 := getVal(1, inst, data)
		arg3 := inst.arguments[2]
		if arg1 < arg2 {
			fmt.Printf("%v < %v, storing 1 in %v\n", arg1, arg2, arg3)
			data[arg3] = 1
		} else {
			fmt.Printf("%v >= %v, storing 0 in %v\n", arg1, arg2, arg3)
			data[arg3] = 0
		}
	case equals:
		arg1 := getVal(0, inst, data)
		arg2 := getVal(1, inst, data)
		arg3 := inst.arguments[2]
		if arg1 == arg2 {
			fmt.Printf("%v == %v, storing 1 in %v", arg1, arg2, arg3)
			data[arg3] = 1
		} else {
			fmt.Printf("%v != %v, storing 0 in %v", arg1, arg2, arg3)
			data[arg3] = 0
		}
	default:
		return 0, fmt.Errorf("Unknown opcode: %v", inst.opcode)
	}
	return rval, nil
}

func getVal(argIndex int, inst *instruction, data []int) int {
	arg := inst.arguments[argIndex]
	mode := inst.modes[argIndex]
	var rval int
	if mode == 0 {
		rval = data[arg]
		fmt.Printf("get %v from %v\n", rval, arg)
	} else {
		rval = arg
	}
	return rval
}

func parseInstruction(ptr int, data []int) (*instruction, error) {
	opStr := strconv.Itoa(data[ptr])
	rval := instruction{0, nil, nil}
	var err error
	opLen := len(opStr)
	var modeRunes []rune
	if opLen > 2 {
		index := opLen - 2
		rval.opcode, err = strconv.Atoi(opStr[index:])
		if err != nil {
			return nil, err
		}
		if index > 0 {
			modeStr := opStr[0:index]
			modeRunes = []rune(modeStr)
			for i, j := 0, len(modeStr)-1; i < j; i, j = i+1, j-1 {
				modeRunes[i], modeRunes[j] = modeRunes[j], modeRunes[i]
			}
		}
	} else {
		rval.opcode = data[ptr]
	}
	var argCount int
	switch rval.opcode {
	case add:
		argCount = 3
	case mult:
		argCount = 3
	case read:
		argCount = 1
	case write:
		argCount = 1
	case quit:
		argCount = 0
	case jumpIfFalse:
		argCount = 2
	case jumpIfTrue:
		argCount = 2
	case lessThan:
		argCount = 3
	case equals:
		argCount = 3
	}
	rval.modes, rval.arguments, err = parseArgs(ptr+1, data, argCount, modeRunes)
	if err != nil {
		return nil, err
	}
	return &rval, nil
}

func parseArgs(ptr int, data []int, argCount int, modeRunes []rune) (modes, args []int, err error) {
	for i := 0; i < argCount; i++ {
		var mode int
		if i < len(modeRunes) {
			mode, err = strconv.Atoi(string(modeRunes[i]))
			if err != nil {
				return
			}
			if mode < 0 || mode > 1 {
				err = fmt.Errorf("Invalid mode %v: ", mode)
				return
			}
		}
		modes = append(modes, mode)
		args = append(args, data[ptr+i])
	}
	return
}

func readData() ([]int, error) {
	return ReadDataFile("../../resources/day5/input.txt")
}

//ReadDataFile reads data from a given file path
func ReadDataFile(fname string) ([]int, error) {
	content, err := ioutil.ReadFile(fname)
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
