package textread

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	DefaultKey      = "NEXT"
	DefaultStartKey = "START"
	nmlReg          = regexp.MustCompile(`\s*(\S+)\s*====>\s*(\S+)\s*`)
	sltReg          = regexp.MustCompile(`\s*(\S+)\s*==(\S+)==>\s*(\S+)\s*`)
	EOF             = errors.New("The end.")
)

type Collection struct {
	Nodes map[string]*Node             `bson: "nodes"`
	M     map[string]map[string]string `bson: "m"` //M[nodeA][slt] -> nodeB
	Used  map[string]bool
}

func (c *Collection) StartFrom(node string) (*Node, error) {
	if node == "" {
		return c.Nodes[DefaultStartKey], nil
	}
	if c.Used[node] {
		return c.Nodes[node], nil
	}
	return nil, fmt.Errorf("Err no this node.")
}

func (c *Collection) Select(node, slt string) *Node {
	if node == "" {
		node = DefaultStartKey
	}
	if slt == "" {
		slt = DefaultKey
	}
	if c.Used[node] && c.M[node][slt] != "" {
		return c.Nodes[c.M[node][slt]] //maybe nil, meaning no subsequent end.
	}
	return nil
}

func newCollection(nodes map[string]*Node) *Collection {
	m := make(map[string]map[string]string)
	used := make(map[string]bool)
	for i := range nodes {
		used[i] = true
	}
	// fmt.Printf("%T, %[1]v", m)
	return &Collection{Nodes: nodes, M: m, Used: used}
}

func ScanCollection(reader *bufio.Reader, nodes map[string]*Node) (*Collection, error) {
	scanner := bufio.NewScanner(reader)
	c := newCollection(nodes)
	used := make(map[string]bool)
	used[DefaultStartKey] = true
	c.Used[DefaultStartKey] = true
	c.Nodes[DefaultStartKey] = &Node{DefaultStartKey, []string{"Text Start."}, nil}

	for scanner.Scan() {
		token := strings.TrimSpace(scanner.Text())
		switch {
		case nmlReg.MatchString(token):
			res := nmlReg.FindStringSubmatch(token)
			if !c.Used[res[1]] || !c.Used[res[2]] {
				return nil, fmt.Errorf("Err this node [%s %s] has not been decleared.", res[0], res[1])
			}
			// if !used[res[0]] || used[res[1]]
			// how to keep away from circle
			if !used[res[1]] {
				return nil, fmt.Errorf("Err this node [%s %s] in wrong logic", res[0], res[1])
			}
			used[res[2]] = true
			// fmt.Printf("%T, %[1]v\n", c.M)
			if c.M[res[1]] == nil {
				c.M[res[1]] = make(map[string]string)
			}
			c.M[res[1]][DefaultKey] = res[2]
		case sltReg.MatchString(token):
			res := sltReg.FindStringSubmatch(token)
			if !c.Used[res[1]] || !c.Used[res[3]] {
				return nil, fmt.Errorf("Err this node [%s %s] has not been decleared.", res[0], res[1])
			}
			// if !used[res[0]] || used[res[1]]
			// how to keep away from circle
			if !used[res[1]] {
				return nil, fmt.Errorf("Err this node [%s %s] in wrong logic", res[0], res[1])
			}
			used[res[3]] = true
			if c.M[res[1]] == nil {
				c.M[res[1]] = make(map[string]string)
			}
			c.M[res[1]][res[2]] = res[3]
		default:
			return nil, fmt.Errorf("Err this token [%s]", token)
		}
	}
	return c, nil
}

func ScanCollectionFile(filename string, nodes map[string]*Node) (*Collection, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(file)
	return ScanCollection(reader, nodes)
}
