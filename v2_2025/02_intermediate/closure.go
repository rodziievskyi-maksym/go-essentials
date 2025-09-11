package main

import (
	"errors"
	"fmt"
	"math/rand"
)

// Closure - Enclosed unit of work

func main2() {
	giftCard := activateGiftCard()

	// I've got information from db - how many times user redeem gift card
	userGiftCardRecords := DBResponseImitation(10)

	for i := range userGiftCardRecords {
		amountLeft, err := giftCard(userGiftCardRecords[i].Redeemed)
		if err != nil {
			userGiftCardRecords[i].Valid = false
			fmt.Printf("amount left: %d - redeemed amount: %d\n", amountLeft, userGiftCardRecords[i].Redeemed)
			continue
		}

		userGiftCardRecords[i].Valid = true
		fmt.Printf("amount left: %d\n", amountLeft)
	}

	// Save Records Imitation
	//Save(userGiftCardRecords)
}

func activateGiftCard() func(uint64) (uint64, error) {
	var amount uint64 = 100

	debitFunc := func(debitAmount uint64) (uint64, error) {
		if amount < debitAmount {
			return amount, errors.New("gift card out of money")
		}
		amount -= debitAmount
		return amount, nil
	}

	return debitFunc
}

type UserGiftCard struct {
	ID       uint64
	Redeemed uint64
	Valid    bool
}

func DBResponseImitation(records int) []UserGiftCard {
	userGiftCards := make([]UserGiftCard, records)

	for i := 0; i < records; i++ {
		userGiftCards[i] = UserGiftCard{
			ID:       uint64(i),
			Redeemed: uint64(rand.Int63n(40)),
		}
	}

	return userGiftCards
}
