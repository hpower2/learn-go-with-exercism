package blackjack

var WordToNumber = map[string]int{
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"ten":   10,
	"ace":   11,
	"jack":  10,
	"queen": 10,
	"king":  10,
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	number := WordToNumber[card]
	return number
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	if ParseCard(card1) + ParseCard(card2) == 21 && ParseCard(dealerCard) < 10{
		return "W"
	}
	if ParseCard(card1) + ParseCard(card2) == 21 && ParseCard(dealerCard) >= 10{
		return "S"
	}
	if ParseCard(card1) + ParseCard(card2) <= 11 {
		return "H"
	}
	if (ParseCard(card1) == ParseCard(card2) && ParseCard(card1) == 11) {
		return "P"
	}
	if ParseCard(card1) + ParseCard(card2) >= 17 && ParseCard(card1) + ParseCard(card2) <= 20{
		return "S"
	}
	if ParseCard(card1) + ParseCard(card2) >= 12 && ParseCard(card1) + ParseCard(card2) <= 16 && ParseCard(dealerCard) >= 7{
		return "H"
	}else{
		return "S"
	}
}
