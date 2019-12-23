package day6

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type node struct {
	name     string
	parent   *node
	children []*node
}

func newNode(name string) *node {
	return &node{name, nil, nil}
}

func (n *node) addNewChild(sat string) *node {
	child := newNode(sat)
	n.addChild(child)
	return child
}

func (n *node) addChild(child *node) {
	n.children = append(n.children, child)
}

func (n *node) countParents() int {
	rval := 0
	parent := n.parent
	for parent != nil {
		rval++
		parent = parent.parent
	}
	return rval
}

func run(in io.Reader) (int, error) {
	nodeMap, _, err := createTree(in)
	if err != nil {
		return 0, err
	}
	rval := 0
	for _, n := range nodeMap {
		rval += n.countParents()
	}
	return rval, nil
}

func getNumJumps(in io.Reader, from, to string) (int, error) {
	nodeMap, _, err := createTree(in)
	if err != nil {
		return 0, err
	}
	fromLine, err := getLineage(nodeMap, from)
	if err != nil {
		return 0, err
	}
	toLine, err := getLineage(nodeMap, to)
	if err != nil {
		return 0, err
	}
	commonNode, err := findLastCommonNode(fromLine, toLine)
	if err != nil {
		return 0, err
	}
	dist1, err := directDist(nodeMap, from, commonNode)
	if err != nil {
		return 0, err
	}
	dist2, err := directDist(nodeMap, to, commonNode)
	if err != nil {
		return 0, err
	}
	return dist1 + dist2, nil
}

func directDist(nodeMap map[string]*node, from, to string) (dist int, err error) {
	n := nodeMap[from]
	if n == nil {
		return 0, fmt.Errorf("Unable to find node %v", from)
	}
	parent := n.parent
	for parent != nil {
		if parent.name == to {
			return
		}
		dist++
		parent = parent.parent
	}
	return dist, fmt.Errorf("Could not find %v in line of %v", to, from)
}

func findLastCommonNode(fromLine, toLine []string) (name string, err error) {
	fromIndex := len(fromLine) - 1
	toIndex := len(toLine) - 1
	if fromLine[fromIndex] != toLine[toIndex] {
		return "", fmt.Errorf("Lines do not share a common root: %v vs %v", fromLine[fromIndex], toLine[toIndex])
	}
	for ; fromIndex > -1 && toIndex > -1; fromIndex, toIndex = fromIndex-1, toIndex-1 {
		toName := toLine[toIndex]
		fromName := fromLine[fromIndex]
		if toName != fromName {
			return
		}
		name = toName
	}
	err = fmt.Errorf("Problem walking lines")
	return
}

func getLineage(nodeMap map[string]*node, name string) ([]string, error) {
	n := nodeMap[name]
	if n == nil {
		return nil, fmt.Errorf("Missing node %v", name)
	}
	var rval []string
	parent := n.parent
	for parent != nil {
		rval = append(rval, parent.name)
		parent = parent.parent
	}
	return rval, nil
}

func createTree(in io.Reader) (nodeMap map[string]*node, root *node, err error) {
	nodeMap = make(map[string]*node)
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		body, sat, err := parseLine(scanner.Text())
		if err != nil {
			return nodeMap, nil, err
		}
		parent := nodeMap[body]
		if parent == nil {
			parent = newNode(body)
			nodeMap[body] = parent
		}
		child := nodeMap[sat]
		if child != nil {
			parent.addChild(child)
			if child.parent != nil {
				if child.parent.name != body {
					return nodeMap, nil, fmt.Errorf("Parent conflict %v != %v", child.parent.name, body)
				}
			} else {
				child.parent = parent
			}
		} else {
			child = parent.addNewChild(sat)
			child.parent = parent
			nodeMap[sat] = child
		}
	}
	root = nodeMap["COM"]
	if root == nil {
		err = fmt.Errorf("Missing COM root")
	}
	return
}

func parseLine(line string) (body, sat string, err error) {
	parts := strings.Split(line, ")")
	if len(parts) != 2 {
		err = fmt.Errorf("Invalid line: %v", line)
	} else {
		body, sat = parts[0], parts[1]
	}
	return
}
