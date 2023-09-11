package main

import (
	"golesson/restful"
	"net/http"
)

func main() {

	// sayi := 20
	// pointers.Demo1(&sayi)
	// fmt.Println("maindeki sayı: ", sayi)

	// sayilar := []int{1, 2, 3}
	// pointers.Demo2(sayilar)
	// fmt.Println("maindeki sayı : ", sayilar[0])

	http.HandleFunc("/get", restful.HandleGet)
	http.HandleFunc("/post", restful.HandlePost)
	http.ListenAndServe(":8080", nil)

}
