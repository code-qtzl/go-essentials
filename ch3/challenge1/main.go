package main

import "fmt"


func calculateFinalBill(costPerMessage float64, numMessages int) float64 {
	baseBill := calculateBaseBill(costPerMessage, numMessages)
	discountRate := calculateDiscountRate(numMessages)

	fmt.Println("Base bill:", baseBill)
	fmt.Println("discountRate:", discountRate)

	return baseBill * (1 - discountRate)
}

func calculateDiscountRate(messagesSent int) float64 {
	if messagesSent > 5000 {
		return 0.20
	} else if messagesSent > 1000 {
		return 0.10
	} else {
		return 0
	}
}

// don't touch below this line

func calculateBaseBill(costPerMessage float64, messagesSent int) float64 {
	return costPerMessage * float64(messagesSent)
}
