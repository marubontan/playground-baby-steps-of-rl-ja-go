package main

import (
	"fmt"
	"strings"
)

func V(s string, gamma float64) float64 {
	value := R(s) + gamma*maxVOnNextState(s)
	return value
}

func R(s string) float64 {
	switch s {
	case "happy_end":
		return 1.0
	case "bad_end":
		return -1.0
	default:
		return 0.0
	}
}

func getMax(a *[]float64) float64 {
	slice := *a
	max := slice[0]
	for _, element := range slice {
		if element > max {
			max = element
		}
	}
	return max
}

func maxVOnNextState(s string) float64 {
	contained := false
	for _, element := range []string{"happy_end", "bad_end"} {
		if s == element {
			contained = true
			break
		}
	}
	if contained {
		return 0.0
	}
	actions := []string{"up", "down"}
	values := []float64{}
	for _, action := range actions {
		transactionProb := transactionFunc(s, action)
		v := 0.0
		for nextStep, prob := range transactionProb {
			v += prob * V(nextStep, 0.99)
		}
		values = append(values, v)
	}
	return getMax(&values)
}

func nextState(state string, action string) string {
	return strings.Join([]string{state, action}, "_")
}

func transactionFunc(s string, a string) map[string]float64 {
	actions := strings.Split(s, "_")[1:]
	const LimitGameCount = 5
	const HappyEndBorder = 4
	const MoveProb = 0.9
	if len(actions) == LimitGameCount {
		upCount := 0
		for _, action := range actions {
			if action == "up" {
				upCount += 1
			}
		}
		var state string
		if upCount >= HappyEndBorder {
			state = "happy_end"
		} else {
			state = "bad_end"
		}
		return map[string]float64{
			state: 1.0,
		}
	} else {
		var opposite string
		if a == "down" {
			opposite = "up"
		} else {
			opposite = "down"
		}
		return map[string]float64{
			nextState(s, a):        MoveProb,
			nextState(s, opposite): 1.0 - MoveProb,
		}
	}
}

func main() {
	fmt.Println(V("state", 0.99))
	fmt.Println(V("state_up_up_up", 0.99))
	fmt.Println(V("state_down_down", 0.99))

}
