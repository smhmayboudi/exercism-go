package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	switch balance := balance; {
	case balance < 0:
		return 3.213
	case 0 <= balance && balance < 1000:
		return 0.5
	case 1000 <= balance && balance < 5000:
		return 1.621
	case 5000 <= balance:
		return 2.475
	default:
		panic("Will not meet!")
	}
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return balance * float64(InterestRate(balance)/100)
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	var i int
	for i = 0; balance < targetBalance; i++ {
		balance += Interest(balance)
	}
	return i
}
