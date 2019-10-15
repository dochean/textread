package textread

import (
	"bufio"
	"errors"

	// "fmt"
	"os"
	"testing"
	// "gopkg.in/mgo.v2/bson"
)

func TestScanReader(t *testing.T) {
	file, err := os.Open("preface.tr")
	if err != nil {
		t.Error(err)
	}

	reader := bufio.NewReader(file)
	_, err = ScanReader(reader)
	if err != nil {
		t.Log(err)
	}
}

func TestScanFile(t *testing.T) {
	_, err := ScanFile("preface1.tr")
	if err != nil {
		t.Error(err)
		t.Log(os.IsNotExist(err))
		t.Log(errors.Is(err, os.ErrNotExist))
	}

}

type nodetext struct {
	novel map[string]*Node `bson: "novel"`
}

func TestScanToMongo(t *testing.T) {
	nodes, err := ScanFile("novel.tr")
	if err != nil {
		t.Error(err)
		t.Log(os.IsNotExist(err))
		t.Log(errors.Is(err, os.ErrNotExist))
	}
	c, cl := C("nodes")
	defer cl()
	c.Insert(nodes)
	// var nodes = make(map[string]*Node)
	// var n = nodetext{make(map[string]*Node)}

	// var n map[string]*Node
	// c.Find(nil).One(&n)
	// t.Logf("%v", n)
	// NodesPrint(n)

	// m, err := ScanManagerFile("novel.trg", n)
	// if err != nil {
	// 	t.Log(err)
	// }
	// s := m.Start()
	// fmt.Println(s.Show())
	// input := ""
	// for {

	// 	sn, err := s.Next(input)
	// 	if err != nil || sn == nil {
	// 		t.Log(err, sn)
	// 		break
	// 	}
	// 	fmt.Println(sn.Show())
	// 	if sn.Select() != nil {
	// 		input = "2"
	// 	}
	// 	s = sn
	// }
}
