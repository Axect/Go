package prob2

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/Axect/Go/Package/check"
)

// =============================================================================
// Type & Method for RSP
// =============================================================================

// Score is type for scoring
type Score struct {
	player int
	enemy  int
	draw   int
}

// =============================================================================
// Map for RSP
// =============================================================================

// RSP for rsp
var RSP = map[string]int{
	"가위": -1,
	"바위": 0,
	"보":  1,
}

// Judge for Score
func (s *Score) Judge(x, y int) {
	// At least one of them put rock
	if x*y == 0 {
		if x > y {
			s.player++
		} else if x < y {
			s.enemy++
		} else {
			s.draw++
		}
	} else if x*y < 0 { // At least one of them put Scissor
		if x < 0 {
			s.player++
		} else {
			s.enemy++
		}
	} else {
		s.draw++
	}
}

// Play is method for Score
func (s *Score) Play(n int) {
	for i := 0; i < n; i++ {
		fmt.Println("가위, 바위, 보를 시작합니다. 가위, 바위, 보!")
		var player string
		_, err := fmt.Scan(&player)
		temp := []string{"가위", "바위", "보"}
		if err != nil || !check.Contains(player, temp) {
			fmt.Println("똑바로 입력하세요.")
			os.Exit(1)
		}
		enemy := AI()
		s.Judge(RSP[player], RSP[enemy])
	}
}

// =============================================================================
// Make Random Engine for RSP
// =============================================================================

//AI for Rsp
func AI() string {
	rand.Seed(int64(time.Now().Nanosecond()))
	number := rand.Intn(3)
	List := [3]string{"가위", "바위", "보"}
	return List[number]
}

//DoRSP !
func DoRSP() {
	fmt.Println("---------------------------------------")
	fmt.Println("Rock Scissors Paper!")
	fmt.Println("---------------------------------------")
	var n int
	fmt.Println("How many times you want to play RSP?")
	fmt.Scanln(&n)
	var Count Score
	Count.Play(n)
	fmt.Printf("Player: %v, Computer: %v, Draw:%v\n", Count.player, Count.enemy, Count.draw)
}
