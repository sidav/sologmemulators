package tiny_solitary_soldiers

import "math/rand"

func GetOracleDiceString(rnd *rand.Rand) string {
	ansString := QuestionDice[rnd.Int()%6]
	twistString := ""
	if rnd.Int()%6 == 0 {
		twistString = "\nWENDUNG: " + TwistDice1[rnd.Int()%6] + " " + TwistDice2[rnd.Int()%6]
	}
	return ansString + twistString
}

func GetTwistDiceString(rnd *rand.Rand) string {
	return "\nDie Wendung ist: " + TwistDice1[rnd.Int()%6] + " " + TwistDice2[rnd.Int()%6]
}

func GetSceneDiceString(rnd *rand.Rand) string {
	rndVal := rnd.Int() % 6
	str := SceneDice[rndVal]
	if rndVal == 5 {
		str += TwistDice1[rnd.Int()%6] + " " + TwistDice2[rnd.Int()%6]
	}
	return str
}
