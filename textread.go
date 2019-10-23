package textread

import (
	"fmt"
	"strings"
)

var (
	SUF_TXT = ".tr"
	SUF_GPH = ".trg"
)

type Novel struct {
	C     *Collection `bson:"collection"`
	Title string      `bson:"title"`
}

func Textread(fileprefix string) (*Novel, error) {
	if fileprefix == "" {
		return nil, fmt.Errorf("Err filename is null.")
	}
	//need to check file
	fileprefix = strings.TrimSpace(fileprefix)
	nodes, err := ScanFile(fileprefix + SUF_TXT)
	if err != nil {
		return nil, fmt.Errorf("Err when scan txt file.")
	}
	co, err := ScanCollectionFile(fileprefix+SUF_GPH, nodes)
	if err != nil {
		return nil, fmt.Errorf("Err when scan graph file.")
	}

	return &Novel{co, fileprefix}, nil
}
