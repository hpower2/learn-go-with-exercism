package triangle


type Kind string

const (
    NaT = "NaT" 
    Equ = "Equ"
    Iso = "Iso"
    Sca = "Sca"
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	k = NaT
	if a > 0 && b > 0 && c > 0{
		if (a > b + c) || (b > a + c) || (c > a + b){
			k = NaT
		}else if a == b && b == c{
			k = Equ
		}else if (a == b && b != c) || (a == c && b != c) || (b == c && b != a){
			k = Iso 
		}else if a != b && b != c && a != c{
			k = Sca
		}
	}
	return k
}
