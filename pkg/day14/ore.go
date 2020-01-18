package day14

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var quantRegex = regexp.MustCompile(`\s*([0-9]+)\s+(\w+)\s*`)

type quantity struct {
	id    string
	count int64
}

type reaction struct {
	output quantity
	inputs []quantity
}

type nanofactory struct {
	reactions map[string]reaction
}

func newFactory(fname string) (*nanofactory, error) {
	in, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer in.Close()
	scanner := bufio.NewScanner(in)
	reactions := make(map[string]reaction)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "=>")
		if len(parts) != 2 {
			return nil, fmt.Errorf("Invalid initial split on line: %v", line)
		}
		inputs, err := parseInput(parts[0])
		if err != nil {
			return nil, err
		}
		output, err := parseQuantity(parts[1])
		if err != nil {
			return nil, err
		}
		reactions[output.id] = reaction{output, inputs}
	}
	if err = scanner.Err(); err != nil {
		return nil, err
	}
	return &nanofactory{reactions}, nil
}

func parseInput(input string) (quants []quantity, err error) {
	parts := strings.Split(input, ",")
	quants = make([]quantity, len(parts))
	for i, p := range parts {
		quants[i], err = parseQuantity(p)
		if err != nil {
			return
		}
	}
	return
}

func parseQuantity(qstr string) (q quantity, err error) {
	parts := quantRegex.FindStringSubmatch(qstr)
	if len(parts) != 3 {
		err = fmt.Errorf("invalid quantity: %v, got %v", qstr, parts)
		return
	}
	q.id = parts[2]
	var i int
	i, err = strconv.Atoi(parts[1])
	q.count = int64(i)
	return
}

func (nf *nanofactory) getMaxFuel(oreAvailable int64) (int64, error) {
	singleFuel, err := nf.getOreRequirement(1)
	if err != nil {
		return -1, err
	}
	lower := oreAvailable / singleFuel
	if err != nil {
		return -1, err
	}
	i, j := lower, lower*2
	for i < j {
		m := i + ((j - i) / 2)
		fmt.Printf("i: %v, j: %v, m: %v\n", i, j, m)
		ore, err := nf.getOreRequirement(m)
		if err != nil {
			return -1, err
		}
		if ore == oreAvailable {
			return m, nil
		}
		if ore > oreAvailable {
			fmt.Printf("too high\n")
			j = m - 1
		} else {
			if j == i+1 {
				break
			}
			fmt.Printf("too low\n")
			i = m
		}
	}
	return i, nil
}

func (nf *nanofactory) getOreRequirement(fuelTarget int64) (int64, error) {
	bank := make(map[string]int64)
	return nf.getOreForReaction("FUEL", fuelTarget, bank)
}

func (nf *nanofactory) getOreForReaction(element string, amount int64, bank map[string]int64) (int64, error) {
	if element == "ORE" {
		return amount, nil
	}
	diff := amount - bank[element]
	rval := int64(0)
	if diff > 0 {
		react, ok := nf.reactions[element]
		if !ok {
			return -1, fmt.Errorf("Unable to find reaction for %v", element)
		}
		reactionsNeeded := int64(math.Ceil(float64(diff) / float64(react.output.count)))
		gain := react.output.count * int64(reactionsNeeded)
		excess := gain - diff
		for _, in := range react.inputs {
			ore, err := nf.getOreForReaction(in.id, in.count*reactionsNeeded, bank)
			if err != nil {
				return -1, err
			}
			rval += ore
		}
		bank[element] = excess
	} else {
		bank[element] -= amount
	}
	return rval, nil
}
