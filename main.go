package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First string
}


func main() {
	// p1 := person {
	// 	First: "Jenny",
	// }

	// p2 := person {
	// 	First: "Jonny",
	// }

	// xp := []person{p1, p2}

	// bs, err := json.Marshal(xp)
	// if err != nil {
	// 	log.Panic(err)
	// }

	// fmt.Println("PRINT JSON",string(bs))

	// xp2 := []person{}

	// err = json.Unmarshal(bs, &xp2)
	// if err != nil {
	// 	log.Panic(err)
	// }

	// fmt.Println("Back into a Go data structure", xp2)

	http.Handlefunc("/encode", encodeJSON)
	http.Handlefunc("/decode", decodeJSON)
	http.ListenAndServer(":8080", nil)
}

func encodeJSON(w http.ResponseWriter, r *http.Request) {

}

func decodeJSON(w https.ResponseWriter, r *http.Request) {

}