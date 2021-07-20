package main

import (
	"fmt"
	"math/rand"
	"sologmemulators/tiny_solitary_soldiers"
	"strconv"
	"strings"
	"time"
)

var (
	rnd     *rand.Rand
	console *consoleWrapperStdout
)

func throwDice(num, sides, mod int) string {
	str := fmt.Sprintf("%dd%d+%d: ", num, sides, mod)
	resultsStr := ""
	res := 0
	for i := 0; i < num; i++ {
		if len(resultsStr) > 0 {
			resultsStr = "," + resultsStr
		}
		currThrow := rnd.Int() % sides + 1
		resultsStr += fmt.Sprintf("%d", currThrow)
		res += currThrow
	}
	str += fmt.Sprintf("%d (%s)", res, resultsStr)
	return str
}

func main() {
	console = &consoleWrapperStdout{}
	console.init()
	defer console.closeConsole()
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		console.print("Wähle einen Würfel zum Würfeln:")
		console.print("\n(F)ragewürfel, (W)endungwürfel, (S)zenewürfel, order XdY: ")
		console.flush()
		typeDice := console.read()
		if typeDice == "" {
			return
		}
		console.print("\n")
		if strings.Contains(typeDice, "d") && len(typeDice) >= 3 {
			num, err := strconv.Atoi(strings.Split(typeDice, "d")[0])
			if err != nil {
				continue
			}
			sides, err := strconv.Atoi(strings.Split(typeDice, "d")[1])
			if err != nil {
				continue
			}
			console.println(throwDice(num, sides, 0))
		}
		switch rune(typeDice[0]) {
		case 'f':
			console.println(tiny_solitary_soldiers.GetOracleDiceString(rnd))
		case 'w':
			console.println(tiny_solitary_soldiers.GetTwistDiceString(rnd))
		case 's':
			console.println(tiny_solitary_soldiers.GetSceneDiceString(rnd))
		}
	}
}
