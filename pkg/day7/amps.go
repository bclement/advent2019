package day7

import (
	"math"
	"sync"

	"github.com/bclement/advent2019/pkg/day5"
)

type amp struct {
	data       []int
	input      chan int
	output     chan int
	lastOutput int
	err        error
}

func newAmp(data []int, input, output chan int, phase int) *amp {
	clone := make([]int, len(data))
	copy(clone, data)
	input <- phase
	return &amp{clone, input, output, 0, nil}
}

func (a *amp) Read() (i int, err error) {
	return <-a.input, nil
}

func (a *amp) Write(i int) error {
	a.output <- i
	a.lastOutput = i
	return nil
}

func (a *amp) run(wg *sync.WaitGroup) {
	defer wg.Done()
	a.err = day5.Run(a.data, a, a)
}

func findMax(data []int, minAmp, maxAmp int) (settings []int, signal int, err error) {
	signal = math.MinInt32
	phases := make([]int, maxAmp-minAmp+1)
	for i := range phases {
		phases[i] = minAmp + i
	}
	perms := getPerms(phases, len(phases))
	for _, testSettings := range perms {
		output, err := run(data, testSettings)
		if err != nil {
			return nil, 0, err
		}
		if output > signal {
			signal = output
			settings = testSettings
		}
	}
	return
}

func getPerms(choices []int, n int) [][]int {
	if n == 1 {
		clone := make([]int, len(choices))
		copy(clone, choices)
		return [][]int{clone}
	}
	var rval [][]int
	for i := 0; i < n; i++ {
		choices[i], choices[n-1] = choices[n-1], choices[i]
		perms := getPerms(choices, n-1)
		rval = append(rval, perms...)
		choices[i], choices[n-1] = choices[n-1], choices[i]
	}
	return rval
}

func run(data, settings []int) (result int, err error) {
	amps := makeAmps(data, settings)
	var wg sync.WaitGroup
	for _, a := range amps {
		wg.Add(1)
		go a.run(&wg)
	}
	amps[0].input <- 0
	wg.Wait()

	for _, a := range amps {
		close(a.input)
		if a.err != nil {
			err = a.err
		}
	}
	result = amps[len(amps)-1].lastOutput
	return
}

func makeAmps(data, settings []int) (amps []*amp) {
	var chans []chan int
	for i := 0; i < len(settings); i++ {
		chans = append(chans, make(chan int, 10))
	}
	for i, phase := range settings {
		next := i + 1
		if next >= len(settings) {
			next = 0
		}
		amps = append(amps, newAmp(data, chans[i], chans[next], phase))
	}
	return
}
