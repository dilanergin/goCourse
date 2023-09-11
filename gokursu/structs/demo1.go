package structs

import "fmt"

func Demo1() {
	fmt.Println(product{"bilgisayar", 500, "xyz"})
}

type product struct {
	name      string
	unitPrice float64
	brand     string
}
