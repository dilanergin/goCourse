package variables

import "fmt"

func Demo1() {

	var name string = "dilan"
	fmt.Println(name)

	name2 := "dilanergin"
	fmt.Println(name2)
	fmt.Printf("veri tipi : %T \n", name2)

	var durum bool
	var object1 int = 6
	var object2 int = 8
	durum = object1 != object2
	fmt.Println(durum)
}
