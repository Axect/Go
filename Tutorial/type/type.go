package main

import "fmt"

type user struct {
	name       string
	email      string
	ext        int
	privileged bool
}

type admin struct {
	person user
	level  string
}

var tom = admin{
	person: user{
		name:       "tom",
		email:      "abcd@naver.com",
		ext:        123,
		privileged: true,
	},
	level: "super",
}

func main() {
	fmt.Println(tom.person.name)
}
