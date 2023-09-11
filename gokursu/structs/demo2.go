package structs

import "fmt"

type customer struct {
	firstName string
	lastName  string
	age       int
}

func (c customer) update() {
	fmt.Println("güncellendi : ", c.firstName)
}
func Demo2() {
	c := customer{firstName: "dilan", lastName: "ergin", age: 24}
	c.update()
}
