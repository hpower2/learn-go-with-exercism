package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	var interest float32
	if balance < 0{
		interest = 3.213
		
	}
	if balance >= 0{
		interest = 0.5
		
	}
	if balance >= 1000{
		interest = 1.621
		
	}
	if balance >= 5000{
		interest = 2.475
		
	}
	return interest
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	pctInterest := InterestRate(balance)
	currentInterest := float64(pctInterest)/100*balance
	return currentInterest
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	annualInterest := Interest(balance)
	totalBalance := annualInterest + balance
	return totalBalance
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance:
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	var start float64
	var count int
	for start = balance ; start <= targetBalance ; start = start + Interest(start){
		count ++
	}
	return count
}
