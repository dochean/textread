package textread_test

import (
	"fmt"
	"testing"

	"github.com/dochean/textread"
)

func TestTextread(t *testing.T) {
	var (
		filename = "novel"
	)
	n, err := textread.Textread(filename)
	if err != nil {
		t.Fatal(err)
	}
	err = textread.Insert(n)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v\n", n)
}

func TestDBOp(t *testing.T) {
	list := textread.List()
	fmt.Printf("%#v\n", list)
	for i := range list {
		fmt.Printf("novel: %#v\n", list[i])
	}
	n := textread.GetById(list[0].Id)
	fmt.Printf("get novel: %#v\n", n)
	fmt.Printf("get novel: %#v\n", n.C)
}
