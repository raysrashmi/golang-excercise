package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    switch card{
        case "ace":
        	return 11
        case "two":
            return 2
        case "three":
        	return 3
        case "four":
        	return 4
        case "five":
        	return 5
        case "six":
        	return 6
        case "seven":
        	return 7
        case "eight":
        	return 8
        case "nine":
            return 9
        case "ten", "jack", "queen", "king":
        	return 10
        default:
        	return 0
    }
	panic("Please implement the ParseCard function")
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	if card1 == "ace" && card2 == "ace" {
		return "P"
	}

    value1 := ParseCard(card1)
	value2 := ParseCard(card2)
	sum := value1 + value2
	dealer := ParseCard(dealerCard)
    
	switch {
	case sum == 21:
		if dealer >= 10 { // ace, king, queen, ten
			return "S"
		}
		return "W"

	case sum >= 17:
		return "S"

	case sum >= 12 && dealer >= 7:
		return "H"

	case sum >= 12:
		return "S"

	default: // sum <= 11
		return "H"
	}
}