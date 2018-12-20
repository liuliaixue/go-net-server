package main

import (
	// "encoding/binary"
	"encoding/json"
	"fmt"
	// "time"
)

func main() {
	fmt.Println("8080")

	message := Message{Cmd: 123}
	fmt.Printf("%+v\n", message)
	var d,_ = json.Marshal(message)
	fmt.Println(d )

	s := []byte(`{
		"cmd": 121,
		"modules": [{"md5":"md5"}]
	}`)
	var obj Message
	err := json.Unmarshal(s, &obj)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", obj)

}

type Module struct {
	Md5  string
	Path string
}

type Message struct {
	Cmd  int  `json:"cmd"`
	Uuid string 
	game_id int
	sver  int
	Modules []Module 
}
