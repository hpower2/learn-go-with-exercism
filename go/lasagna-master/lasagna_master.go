package lasagna
// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int{
	if time != 0{
		return time * len(layers)
	}
	return 2 * len(layers)
} 
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64){
	var temp int
	var tempS float64
	for _, val := range layers{
		if string(val) == "noodles"{
			temp = temp + 50
		}
		if string(val) == "sauce"{
			tempS = tempS + 0.2
		}
	}
	return temp, tempS
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string){
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, amounts int) []float64{
	var value []float64
	for _, val := range quantities{
		value = append(value, val*float64(amounts)/2.0)
	}
	return value
}