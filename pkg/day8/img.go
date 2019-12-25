package day8

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
)

type image struct {
	data       []int
	width      int
	height     int
	layerLen   int
	layerCount int
}

func readImage(in io.Reader, width, height int) (*image, error) {
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanRunes)
	var data []int
	layerLen := width * height
	for scanner.Scan() {
		ch := scanner.Text()
		ch = strings.TrimSpace(ch)
		if ch == "" {
			continue
		}
		i, err := strconv.Atoi(ch)
		if err != nil {
			return nil, err
		}
		data = append(data, i)
	}
	dataLen := len(data)
	if dataLen%layerLen != 0 {
		return nil, fmt.Errorf("read %v values but each layer needs %v", dataLen, layerLen)
	}
	layerCount := dataLen / layerLen
	return &image{data, width, height, layerLen, layerCount}, nil
}

func (img *image) getLayer(index int) []int {
	start := index * img.layerLen
	end := start + img.layerLen
	return img.data[start:end]
}

func (img *image) countLayerDigits(index, digit int) int {
	rval := 0
	layer := img.getLayer(index)
	for i := 0; i < img.layerLen; i++ {
		if layer[i] == digit {
			rval++
		}
	}
	return rval
}

func (img *image) findLayerWithFewest(digit int) (index int) {
	least := math.MaxInt32
	for i := 0; i < img.layerCount; i++ {
		count := img.countLayerDigits(i, digit)
		if count < least {
			least = count
			index = i
		}
	}
	return
}

func (img *image) getChecksum() int {
	index := img.findLayerWithFewest(0)
	ones := img.countLayerDigits(index, 1)
	twos := img.countLayerDigits(index, 2)
	return ones * twos
}

func (img *image) render() string {
	layer := img.flatten()
	var sb strings.Builder
	for j := 0; j < img.height; j++ {
		for i := 0; i < img.width; i++ {
			num := layer[(j*img.width)+i]
			var ch rune
			if num == 0 {
				ch = ' '
			} else {
				ch = '█'
			}
			sb.WriteRune(ch)
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

func (img *image) flatten() []int {
	rval := createTransparentLayer(img.layerLen)
	for i := 0; i < img.layerCount; i++ {
		layer := img.getLayer(i)
		applyLayer(layer, rval)
	}
	return rval
}

func applyLayer(src, dest []int) {
	for i, v := range src {
		if dest[i] == 2 {
			dest[i] = v
		}
	}
}

func createTransparentLayer(size int) []int {
	rval := make([]int, size)
	for i := range rval {
		rval[i] = 2
	}
	return rval
}
