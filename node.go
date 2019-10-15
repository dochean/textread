package textread

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

type Node struct {
	Name   string            `bson: "name"`
	Para   []string          `bson: "para"`
	Select map[string]string `bson: "select"` //name: used to choose next node value: select placeholder
}

var (
	DefaultParaLength = 20
	nameReg           = regexp.MustCompile(`\[\s*(?P<1>\S+)\s*\]`)
	selectReg         = regexp.MustCompile(`>>\s*([\S]+)\s*:\s*([\S\s*\S]+)\s*`)
)

func ScanFile(filename string) (map[string]*Node, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(file)
	return ScanReader(reader)
}

func ScanReader(reader *bufio.Reader) (map[string]*Node, error) {
	scanner := bufio.NewScanner(reader)
	// var i int
	var para []string
	var name string
	var selects map[string]string
	// var nodes []*Node
	var nodes = make(map[string]*Node)
	for scanner.Scan() {
		token := strings.TrimSpace(scanner.Text())
		switch {
		case nameReg.MatchString(token):
			res := nameReg.FindStringSubmatch(token)
			name = res[1]
			if _, ok := nodes[name]; ok {
				return nil, fmt.Errorf("Err node name %s existed already\n", name)
			}
			para = make([]string, 0, DefaultParaLength)
			selects = make(map[string]string)
		case selectReg.MatchString(token):
			if name == "" {
				continue
			}
			res := selectReg.FindStringSubmatch(token)
			// fmt.Println(res)
			selects[res[1]] = res[2]
		case token == "":
			if name == "" {
				continue
			}
			// nodes = append(nodes, &Node{name, para, selects})
			nodes[name] = &Node{name, para, selects}
			para = nil
			name = ""
			selects = nil
		default:
			para = append(para, token)
		}
	}
	if name != "" {
		// nodes = append(nodes, &Node{name, para, selects})
		nodes[name] = &Node{name, para, selects}
	}
	// NodesPrint(nodes)
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return nodes, nil
}

func NodesPrint(nodes map[string]*Node) {
	out := &bytes.Buffer{}
	for i := range nodes {
		node := nodes[i]
		fmt.Fprintf(out, "Name: %s\n", node.Name)
		for j := range node.Para {
			fmt.Fprintf(out, "[%d] %s %d\n", j, node.Para[j], len(node.Para[j]))
		}
		for j := range node.Select {
			fmt.Fprintf(out, ">>%s: %s %d\n", j, node.Select[j], len(node.Select[j]))
		}
	}
	io.Copy(os.Stdout, out)
}
