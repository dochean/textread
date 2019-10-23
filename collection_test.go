package textread

import (
	"fmt"
	"testing"
)

func TestScanCollectionFile(t *testing.T) {
	// file, err := os.Open("novel.tr")
	// if err != nil {
	// 	t.Error(err)
	// }
	// reader := bufio.NewReader(file)
	// nodes, err := ScanReader(reader)
	// if err != nil {
	// 	t.Log(err)
	// }
	// c, err := ScanCollectionFile("novel.trg", nodes)
	// if err != nil {
	// 	t.Log(err)
	// }
	// t.Logf("%#v", c)

	cc, cl := C("nodes")
	defer cl()
	// var m *manager
	// cc.Insert(&Novel{c, "test"})
	// cc.Find(nil).One(&m)
	var n Novel
	cc.Find(nil).One(&n)
	fmt.Printf("%#v\n", n.C)

	input := []string{"3"}

	var slt string
	fmt.Println(n.C.StartFrom(""))
	for node, _ := n.C.StartFrom(""); node != nil; node = n.C.Select(node.Name, slt) {
		for p := range node.Para {
			fmt.Println(node.Para[p])
		}
		if len(node.Select) == 0 {
			// next = node.Name
			slt = ""
			continue
		}
		for i := range node.Select {
			fmt.Printf("[%s]: %s\n", i, node.Select[i])
		}
		fmt.Printf("input: %s\n", input[0])
		// next = n.C.M[node.Name][input]
		slt = input[0]
		// next = node.Name

	}
}
