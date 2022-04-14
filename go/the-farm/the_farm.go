package thefarm

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	numFood, err := weightFodder.FodderAmount()
	var foodPerCow float64
	if err == ErrScaleMalfunction{
		foodPerCow = numFood/float64(cows)*2
		return foodPerCow, err
	}else if err == nil{
		foodPerCow = numFood/float64(cows)
		return foodPerCow, err
	}
	return 0, err
}
