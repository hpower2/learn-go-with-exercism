package pascal

func calculateRow(n int) []int{
	if n == 1{
		return []int{1}
	}else if n== 2{
		return []int{1 ,1}
	}
	tempRow := []int{1, 1}
	for i := 0; i <= n-3 ; i++{
		var newRow []int
		newRow = append(newRow, 1)
		l := len(tempRow)
		for j := 0 ; j < l-1; j++{
			newRow = append(newRow, tempRow[j]+tempRow[j+1])
		}
		newRow = append(newRow, 1)
		tempRow = newRow
	}
	return tempRow
}

func Triangle(n int) (out [][]int) {
	for i:= 1; i <= n ; i++{
		out = append(out, calculateRow(i))
	}
	return out
}
