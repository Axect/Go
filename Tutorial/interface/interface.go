package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("%s: %s\n", u.name, u.email)
}

type admin user

func (a *admin) notify() {
	fmt.Printf("%s: %s\n", a.name, a.email)
}

func main() {
	Tom := user{"Tom", "abcd@naver.com"}
	sendNotification(&Tom)

	Lisa := admin{"Lisa", "qwer@naver.com"}
	sendNotification(&Lisa)
}

func sendNotification(n notifier) {
	n.notify()
}
