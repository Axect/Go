package main

import (
	"fmt"
	"log"
	"time"
)

// MAP is a map
var MAP = map[string]int{
	"왼쪽":  1,
	"오른쪽": 2,
}

// Choose for choose
var Choose int

func main() {
	PRINT("예디게임을 시작합니다!")
	fmt.Println()
	PRINT("- 1024호")
	fmt.Println()
	PRINT("문득 잠에서 깬 예디는 태고미한테 연락을 합니다.")
	fmt.Println()
	for i := 0; i < 3; i++ {
		SLEEP(1)
		PRINT("뚜루루루..")
	}
	fmt.Println()
	PRINT("전화를 받지 않는 태고미! 자는 것이 분명합니다!")
	PRINT("열 받은 예디는 계속 전화를 겁니다. 3번의 전화 끝에 태고미가 받았습니다.")
	SLEEP(1)
	fmt.Println()
	PRINT("예디: 내가 일찍오라고 했쥐!")
	PRINT("태곰: 어제 늦게 자서..")
	PRINT("예디: 30분안에 나와라.")
	PRINT("태곰: 응!")
	fmt.Println()
	PRINT("첫번째 문제: 현재시간은 10시! 택미는 몇분후에 나올 확률이 가장 높을까요?")
	fmt.Println("1. 30분 후")
	fmt.Println("2. 1시간 후")
	fmt.Println("3. 1시간 30분 후")

	fmt.Scanln(&Choose)

	First(Choose)
}

func SLEEP(i int) {
	time.Sleep(time.Duration(i) * time.Second)
}

func PRINT(str string) {
	fmt.Println(str)
	SLEEP(1)
}

func YES() {
	PRINT("정답입니다!\n")
}

func First(i int) {
	switch i {
	case 1:
		fmt.Println("아직도 태고미를 모르는군요!")
		log.Fatal("흥!")
	case 2:
		fmt.Println("가끔은 그러지만 평소는 아닙니다! 지금 시각은 10시죠!")
		log.Fatal("시간을 제대로 보세요.")
	case 3:
		YES()
	}
}
