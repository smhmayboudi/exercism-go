package wordy

import (
	"math"
	"strconv"
	"strings"
)

func delimiters(ch rune) bool {
	return ch == ' ' || ch == '?'
}

type info struct {
	num   int
	op    string
	state string
}
type handler func(*info, string) bool

var stateMap = make(map[string]handler)

func init() {
	// generate the map of states to functions */
	stateMap["initial"] = wordState([]string{"What"})
	stateMap["What"] = wordState([]string{"is"})
	stateMap["is"] = numState()
	stateMap["number"] = wordState([]string{"plus", "minus", "multiplied", "divided", "to"})
	stateMap["plus"] = numState()
	stateMap["minus"] = numState()
	stateMap["multiplied"] = wordState([]string{"by"})
	stateMap["divided"] = wordState([]string{"by"})
	stateMap["by"] = numState()
	stateMap["to"] = wordState([]string{"the"})
	stateMap["the"] = wordState([]string{"power"})
	stateMap["power"] = wordState([]string{"of"})
	stateMap["of"] = numState()
}

// wordState returns a function that handles expected word inputs
func wordState(nextState []string) handler {
	return func(stateInfo *info, word string) bool {
		for _, v := range nextState {
			if word == v {
				stateInfo.state = v
				if stateInfo.op == "" {
					stateInfo.op = v
				}
				return true
			}
		}
		return false
	}
}

// numState returns a function that handles expected numeric inputs
func numState() handler {
	return func(stateInfo *info, word string) bool {
		n, err := strconv.Atoi(word)
		if err != nil {
			return false
		}
		if stateInfo.op == "What" {
			stateInfo.num = n
		} else if stateInfo.op == "plus" {
			stateInfo.num += n
		} else if stateInfo.op == "minus" {
			stateInfo.num -= n
		} else if stateInfo.op == "multiplied" {
			stateInfo.num *= n
		} else if stateInfo.op == "divided" && n != 0 {
			stateInfo.num /= n
		} else if stateInfo.op == "to" {
			stateInfo.num = int(math.Pow(float64(stateInfo.num), float64(n)))
		} else {
			return false
		}
		stateInfo.op = ""
		stateInfo.state = "number"
		return true
	}
}

// Answer answers simple integer arithmetic questions
func Answer(input string) (int, bool) {
	words := strings.FieldsFunc(input, delimiters)
	stateInfo := info{0, "", "initial"}
	for _, v := range words {
		if v == "" {
			continue
		}
		if !stateMap[stateInfo.state](&stateInfo, v) {
			return 0, false
		}
	}
	if stateInfo.state != "number" {
		return 0, false
	}
	return stateInfo.num, true
}
