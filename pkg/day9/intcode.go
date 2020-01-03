package day9

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
	offsetBase  = 9
	quit        = 99
)

// Computer contains the memory and IO for an intcode computer
type Computer struct {
	data  []int64
	base  int
	in    Intsrc
	out   Intdest
	debug bool
}

// NewComputer creates a new computer and returns a pointer to it
func NewComputer(data []int64, in Intsrc, out Intdest, debug bool) *Computer {
	return &Computer{data, 0, in, out, debug}
}

type instruction struct {
	opcode    int
	modes     []int
	arguments []int64
}

// Intsrc reads integers
type Intsrc interface {
	Read() (int64, error)
}

// Intdest writes integers
type Intdest interface {
	Write(i int64) error
}

// SliceIO uses slices as input and out for a program
type SliceIO struct {
	input  []int64
	pos    int
	output []int64
}

// Read gets the next input or an error
func (io *SliceIO) Read() (int64, error) {
	if io.pos > len(io.input) {
		return 0, fmt.Errorf("Ran out of input at %v", io.pos)
	}
	rval := io.input[io.pos]
	io.pos++
	return rval, nil
}

// Write writes the output to a slice
func (io *SliceIO) Write(i int64) error {
	io.output = append(io.output, i)
	return nil
}

// NewSliceIO creates input/output storage backed by slices
func NewSliceIO(input []int64) *SliceIO {
	return &SliceIO{input, 0, nil}
}

// RunAsync should be ran as a gorutine and passes an error to the
// channel on error or nil otherwise
func (comp *Computer) RunAsync(errc chan error) {
	rval := comp.Run()
	errc <- rval
}

// Run runs the provided data getting input from in and outputting to out
func (comp *Computer) Run() error {
	ptr := 0
	for ptr < len(comp.data) {
		inst, err := parseInstruction(ptr, comp.data)
		if err != nil {
			return fmt.Errorf("Problem parsing at %v: %v", ptr, err)
		}
		if inst.opcode == quit {
			return nil
		}
		var jumpDest int
		jumpDest, err = comp.execute(inst)
		if err != nil {
			return fmt.Errorf("Problem running at %v: %v", ptr, err)
		}
		if jumpDest > -1 {
			ptr = jumpDest
		} else {
			ptr += len(inst.arguments) + 1
		}
	}
	return fmt.Errorf("Ran out of data before termination")
}

func (comp *Computer) execute(inst *instruction) (int, error) {
	rval := -1
	switch inst.opcode {
	case add:
		arg1 := comp.getVal(0, inst)
		arg2 := comp.getVal(1, inst)
		arg3 := comp.getIndex(2, inst)
		if comp.debug {
			fmt.Printf("%v + %v -> %v\n", arg1, arg2, arg3)
		}
		comp.setValue(arg3, arg1+arg2)
	case mult:
		arg1 := comp.getVal(0, inst)
		arg2 := comp.getVal(1, inst)
		arg3 := comp.getIndex(2, inst)
		if comp.debug {
			fmt.Printf("%v * %v -> %v\n", arg1, arg2, arg3)
		}
		comp.setValue(arg3, arg1*arg2)
	case read:
		if comp.debug {
			fmt.Printf("waiting to read\n")
		}
		i, err := comp.in.Read()
		if err != nil {
			return 0, err
		}
		index := comp.getIndex(0, inst)
		if comp.debug {
			fmt.Printf("read %v -> %v\n", i, index)
		}
		comp.setValue(index, i)
	case write:
		arg := comp.getVal(0, inst)
		if comp.debug {
			fmt.Printf("write %v\n", arg)
		}
		err := comp.out.Write(arg)
		if err != nil {
			return 0, err
		}
	case jumpIfTrue:
		arg1 := comp.getVal(0, inst)
		arg2 := comp.getVal(1, inst)
		if arg1 != 0 {
			if comp.debug {
				fmt.Printf("%v != 0, jumping to %v\n", arg1, arg2)
			}
			rval = int(arg2)
		} else {
			if comp.debug {
				fmt.Printf("not jumping\n")
			}
		}
	case jumpIfFalse:
		arg1 := comp.getVal(0, inst)
		arg2 := comp.getVal(1, inst)
		if arg1 == 0 {
			if comp.debug {
				fmt.Printf("%v == 0, jumping to %v\n", arg1, arg2)
			}
			rval = int(arg2)
		} else {
			if comp.debug {
				fmt.Printf("not jumping\n")
			}
		}
	case lessThan:
		arg1 := comp.getVal(0, inst)
		arg2 := comp.getVal(1, inst)
		arg3 := comp.getIndex(2, inst)
		if arg1 < arg2 {
			if comp.debug {
				fmt.Printf("%v < %v, storing 1 in %v\n", arg1, arg2, arg3)
			}
			comp.setValue(arg3, 1)
		} else {
			if comp.debug {
				fmt.Printf("%v >= %v, storing 0 in %v\n", arg1, arg2, arg3)
			}
			comp.setValue(arg3, 0)
		}
	case equals:
		arg1 := comp.getVal(0, inst)
		arg2 := comp.getVal(1, inst)
		arg3 := comp.getIndex(2, inst)
		if arg1 == arg2 {
			if comp.debug {
				fmt.Printf("%v == %v, storing 1 in %v\n", arg1, arg2, arg3)
			}
			comp.setValue(arg3, 1)
		} else {
			if comp.debug {
				fmt.Printf("%v != %v, storing 0 in %v\n", arg1, arg2, arg3)
			}
			comp.setValue(arg3, 0)
		}
	case offsetBase:
		arg1 := comp.getVal(0, inst)
		if comp.debug {
			fmt.Printf("adding %v to base %v\n", arg1, comp.base)
		}
		comp.base += int(arg1)
		if comp.debug {
			fmt.Printf("base is now %v\n", comp.base)
		}
	default:
		return 0, fmt.Errorf("Unknown opcode: %v", inst.opcode)
	}
	return rval, nil
}

func (comp *Computer) getIndex(argIndex int, inst *instruction) int {
	rval := int(inst.arguments[argIndex])
	if inst.modes[argIndex] == 2 {
		rval += comp.base
	}
	return rval
}

func (comp *Computer) setValue(index int, val int64) {
	comp.extendMemory(index + 1)
	comp.data[index] = val
}

func (comp *Computer) getValue(index int) int64 {
	comp.extendMemory(index + 1)
	return comp.data[index]
}

func (comp *Computer) extendMemory(newLen int) {
	oldLen := len(comp.data)
	for i := oldLen; i < newLen; i++ {
		comp.data = append(comp.data, 0)
	}
}

func (comp *Computer) getVal(argIndex int, inst *instruction) int64 {
	arg := inst.arguments[argIndex]
	mode := inst.modes[argIndex]
	var rval int64
	if mode == 0 {
		rval = comp.getValue(int(arg))
		if comp.debug {
			fmt.Printf("reference: get %v from %v\n", rval, arg)
		}
	} else if mode == 1 {
		rval = arg
		if comp.debug {
			fmt.Printf("get %v directly\n", rval)
		}
	} else if mode == 2 {
		offset := int(arg)
		index := offset + comp.base
		rval = comp.getValue(index)
		if comp.debug {
			fmt.Printf("offset: get %v from %v\n", rval, index)
		}
	}
	return rval
}

func parseInstruction(ptr int, data []int64) (*instruction, error) {
	opStr := strconv.FormatInt(data[ptr], 10)
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
		rval.opcode = int(data[ptr])
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
	case offsetBase:
		argCount = 1
	}
	rval.modes, rval.arguments, err = parseArgs(ptr+1, data, argCount, modeRunes)
	if err != nil {
		return nil, err
	}
	return &rval, nil
}

func parseArgs(ptr int, data []int64, argCount int, modeRunes []rune) (modes []int, args []int64, err error) {
	for i := 0; i < argCount; i++ {
		var mode int
		if i < len(modeRunes) {
			mode, err = strconv.Atoi(string(modeRunes[i]))
			if err != nil {
				return
			}
			if mode < 0 || mode > 2 {
				err = fmt.Errorf("Invalid mode %v: ", mode)
				return
			}
		}
		modes = append(modes, mode)
		args = append(args, data[ptr+i])
	}
	return
}

//ReadDataFile reads data from a given file path
func ReadDataFile(fname string) ([]int64, error) {
	content, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}
	texts := strings.Split(string(content), ",")
	rval := make([]int64, len(texts))
	for i := 0; i < len(texts); i++ {
		num, err := strconv.ParseInt(strings.TrimSpace(texts[i]), 10, 64)
		if err != nil {
			return nil, err
		}
		rval[i] = num
	}
	return rval, nil
}
