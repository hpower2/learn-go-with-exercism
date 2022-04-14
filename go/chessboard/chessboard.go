package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	var counter int
	for r, value := range cb{
		if rank == r{
			for _, val := range value{
				if val{
					counter ++
				}
			}
		}
	}
	return counter
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	var counter int
	for _, value := range cb{
		for in, val := range value{
			if in+1 == file{
				if val{
					counter++
				}
			}
		}
	}
	return counter
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	var counter int
	for _, val := range cb{
		for _, stat := range val{
			if stat || !stat{
				counter++
			}
		}
	}
	return counter
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	var counter int
	for _, val := range cb{
		for _, stat := range val{
			if stat{
				counter++
			}
		}
	}
	return counter
}
