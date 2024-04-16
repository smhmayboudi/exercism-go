package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "eight":
		return 8
	case "five":
		return 5
	case "four":
		return 4
	case "jack":
		fallthrough
	case "king":
		fallthrough
	case "queen":
		fallthrough
	case "ten":
		return 10
	case "nine":
		return 9
	case "seven":
		return 7
	case "six":
		return 6
	case "three":
		return 3
	case "two":
		return 2
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	pCard1 := ParseCard(card1)
	pCard2 := ParseCard(card2)
	pDealerCard := ParseCard(dealerCard)
	switch sum := pCard1 + pCard2; {
	case sum == 22:
		return "P"
	case sum == 21:
		if pDealerCard == 10 || pDealerCard == 11 {
			return "S"
		}
		return "W"
	case 17 <= sum && sum <= 20:
		return "S"
	case 12 <= sum && sum <= 16:
		if 7 <= pDealerCard {
			return "H"
		}
		return "S"
	case sum <= 11:
		return "H"
	default:
		return ""
	}
}
