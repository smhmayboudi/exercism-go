// Package poker solves the exercism poker problem
package poker

import (
	"errors"
	"sort"
	"strings"
	"unicode/utf8"
)

// A Card is a byte (to represent pip values) and a rune (to represent the suit)
type Card struct {
	pip  byte
	suit rune
}

// A Hand is a list of Cards
type Hand []Card

// parseCard takes string representation of a single card, such as 10♡, and returns a Card
func parseCard(c string) (Card, error) {
	suit, size := utf8.DecodeLastRuneInString(c)
	if suit != '♤' && suit != '♢' && suit != '♡' && suit != '♧' {
		return Card{}, errors.New("invalid suit in " + c)
	}
	c = c[:len(c)-size]
	pip := map[string]byte{"2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "10": 10, "J": 11, "Q": 12, "K": 13, "A": 14}[c]
	if pip == 0 {
		return Card{}, errors.New("failed to find pipvalue in " + c)
	}
	return Card{pip: pip, suit: suit}, nil
}

// Parse takes a space separated representation of a hand and returns an array of Card
// sorted by pip value
func Parse(h string) (Hand, error) {
	cardStrings := strings.Split(h, " ")
	if len(cardStrings) != 5 {
		return Hand{}, errors.New("Wrong number of cards")
	}
	hand := make(Hand, 5)
	var err error
	for i, c := range cardStrings {
		hand[i], err = parseCard(c)
		if err != nil {
			return hand, err
		}
	}
	sort.Slice(hand, func(i, j int) bool { return hand[i].pip > hand[j].pip })
	return hand, nil
}

// High etc are constant values to name types of hands
const (
	High = byte(iota)
	Pair
	TwoPair
	Three
	Straight
	Flush
	FullHouse
	Four
	StraightFlush
)

func valueHigh(hand Hand) []byte {
	v := make([]byte, 6)
	v[0] = High
	for i := range hand {
		v[i+1] = byte(hand[i].pip)
	}
	return v
}

func valuePair(hand Hand) []byte {
	v := make([]byte, 6)
	if hand[0].pip == hand[1].pip {
		v[0] = Pair
		v[1] = hand[0].pip //the pip value of the pair
		v[2] = hand[2].pip //the pip value of the remaining cards
		v[3] = hand[3].pip //descending order
		v[4] = hand[4].pip
		return v
	}
	if hand[1].pip == hand[2].pip {
		v[0] = Pair
		v[1] = hand[1].pip
		v[2] = hand[0].pip
		v[3] = hand[3].pip
		v[4] = hand[4].pip
		return v
	}
	if hand[2].pip == hand[3].pip {
		v[0] = Pair
		v[1] = hand[2].pip
		v[2] = hand[0].pip
		v[3] = hand[1].pip
		v[4] = hand[4].pip
		return v
	}
	if hand[3].pip == hand[4].pip {
		v[0] = Pair
		v[1] = hand[3].pip
		v[2] = hand[0].pip
		v[3] = hand[1].pip
		v[4] = hand[2].pip
		return v
	}
	return v // array of zeros
}

func valueTwoPair(hand Hand) []byte {
	v := make([]byte, 6)
	if hand[0].pip == hand[1].pip && hand[2].pip == hand[3].pip {
		v[0] = TwoPair
		v[1] = hand[0].pip //pip of top pair
		v[2] = hand[2].pip //pip of second pair
		v[3] = hand[4].pip //pip of kicker
		return v
	}
	if hand[0].pip == hand[1].pip && hand[3].pip == hand[4].pip {
		v[0] = TwoPair
		v[1] = hand[0].pip //pip of top pair
		v[2] = hand[3].pip //pip of second pair
		v[3] = hand[2].pip //pip of kicker
		return v
	}
	if hand[1].pip == hand[2].pip && hand[3].pip == hand[4].pip {
		v[0] = TwoPair
		v[1] = hand[1].pip //pip of top pair
		v[2] = hand[3].pip //pip of second pair
		v[3] = hand[0].pip //pip of kicker
		return v
	}
	return v
}

func valueThree(hand Hand) []byte {
	v := make([]byte, 6)
	if hand[0].pip == hand[2].pip {
		v[0] = Three
		v[1] = hand[0].pip //the pip value of the three
		v[2] = hand[3].pip //the pip value of the remaining cards
		v[3] = hand[4].pip //descending order
		return v
	}
	if hand[1].pip == hand[3].pip {
		v[0] = Three
		v[1] = hand[1].pip //the pip value of the three
		v[2] = hand[0].pip //the pip value of the remaining cards
		v[3] = hand[4].pip //descending order
		return v
	}
	if hand[2].pip == hand[4].pip {
		v[0] = Three
		v[1] = hand[2].pip //the pip value of the three
		v[2] = hand[0].pip //the pip value of the remaining cards
		v[3] = hand[1].pip //descending order
		return v
	}
	return v
}

func valueStraight(hand Hand) []byte {
	v := make([]byte, 6)
	if hand[0].pip == hand[1].pip+1 &&
		hand[1].pip == hand[2].pip+1 &&
		hand[2].pip == hand[3].pip+1 &&
		hand[3].pip == hand[4].pip+1 {
		v[0] = Straight
		v[1] = hand[0].pip
		return v
	}
	//special case, low straight A-5.
	if hand[0].pip == 14 &&
		hand[1].pip == 5 &&
		hand[2].pip == 4 &&
		hand[3].pip == 3 &&
		hand[4].pip == 2 {
		v[0] = Straight
		v[1] = 5
		return v
	}
	return v
}

func valueFlush(hand Hand) []byte {
	v := make([]byte, 6)
	if hand[0].suit == hand[1].suit &&
		hand[1].suit == hand[2].suit &&
		hand[2].suit == hand[3].suit &&
		hand[3].suit == hand[4].suit {
		v[0] = Flush
		for i := range hand {
			v[i+1] = byte(hand[i].pip)
		}
	}
	return v
}

func valueFullHouse(hand Hand) []byte {
	v := make([]byte, 6)
	if hand[0].pip == hand[2].pip && hand[3].pip == hand[4].pip {
		v[0] = FullHouse
		v[1] = hand[0].pip
		v[2] = hand[3].pip
		return v
	}
	if hand[0].pip == hand[1].pip && hand[2].pip == hand[4].pip {
		v[0] = FullHouse
		v[1] = hand[4].pip
		v[2] = hand[0].pip
		return v
	}
	return v
}
func valueFour(hand Hand) []byte {
	v := make([]byte, 6)
	if hand[0].pip == hand[3].pip {
		v[0] = Four
		v[1] = hand[0].pip
		v[2] = hand[4].pip
		return v
	}
	if hand[1].pip == hand[4].pip {
		v[0] = Four
		v[1] = hand[1].pip
		v[2] = hand[0].pip
		return v
	}
	return v
}

func valueStraightFlush(hand Hand) []byte {
	v := valueFlush(hand)
	if v[0] == Flush {
		v = valueStraight(hand)
		if v[0] == Straight {
			v[0] = StraightFlush
			return v
		}
	}
	return make([]byte, 6)
}

// Value takes a sorted hand and returns a slice of int that values it
// the first number represents the type of hand (eg Pair) the remaining
// numbers are the pip values used to discriminate different pairs
func Value(hand Hand) []byte {
	var v []byte
	v = valueStraightFlush(hand)
	if v[0] == StraightFlush {
		return v
	}
	v = valueFour(hand)
	if v[0] == Four {
		return v
	}
	v = valueFullHouse(hand)
	if v[0] == FullHouse {
		return v
	}
	v = valueFlush(hand)
	if v[0] == Flush {
		return v
	}
	v = valueStraight(hand)
	if v[0] == Straight {
		return v
	}
	v = valueThree(hand)
	if v[0] == Three {
		return v
	}
	v = valueTwoPair(hand)
	if v[0] == TwoPair {
		return v
	}
	v = valuePair(hand)
	if v[0] == Pair {
		return v
	}
	v = valueHigh(hand)
	return v
}

type handValue struct {
	h string
	v string
}

// cream the top off the ordered list of valued hands
func cream(vh []handValue) []string {
	besthands := []string{}
	for _, v := range vh {
		if v.v == vh[0].v {
			besthands = append(besthands, v.h)
		}
	}
	return besthands
}

// BestHand Finds the best poker hand from a list
func BestHand(hands []string) ([]string, error) {
	//We assign a value to each hand in the form of a string.
	//we store each hand with its value in an array
	valuedHands := make([]handValue, len(hands))
	for i, h := range hands {
		//parse the hand into an array of Cards, with a suit and a pip
		hand, err := Parse(h)
		if err != nil {
			return []string{}, err
		}
		//Evaluate each hand
		valuedHands[i] = handValue{h: h, v: string(Value(hand))}
	}
	//sort the hands by their value. Best hands first
	sort.Slice(valuedHands, func(i, j int) bool { return valuedHands[i].v > valuedHands[j].v })
	//return the best hands
	return cream(valuedHands), nil
}
