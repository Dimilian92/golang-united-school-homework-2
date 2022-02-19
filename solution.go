package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type amountOfSides int

const SidesTriangle amountOfSides = 3
const SidesSquare amountOfSides = 4
const SidesCircle amountOfSides = 0

func CalcSquare(sideLen float64, sidesNum amountOfSides) float64 {

	if sidesNum == SidesTriangle {
		return (math.Pow(sideLen, 2.0) * math.Sqrt(3.0)) / 4
	} else if sidesNum == SidesSquare {
		return math.Pow(sideLen, 2.0)
	} else if sidesNum == SidesCircle {
		return math.Pi * math.Pow(sideLen, 2.0)
	}
	return 0
}
