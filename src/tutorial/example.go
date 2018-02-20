package main

import "fmt"
import "math"
// import "github.com/go-chef/chef" after a "go get" on it.

// Package level var & multiple initializer & type-inference
var packageLevelVar string
var i, name = 1, "nils"

var (
		myBool bool = false
		myInt int = -1000 // also existing : int  int8  int16  int32  int64 and rune=int32. Default : 0
		myUnsignedInt uint = 1000 // also existing uint uint8 uint16 uint32 uint64 uintptr and byte=uint8. Default : 0
		myFloat float64 = 3.2 //also existing float32 float64. Default : 0
		myString string = "string"// Default : ""
		myComplex complex64 = -5 + 12i //also existing complex64, complex128. Default : 0
)

func addAndMult(x, y int) (int, int) {
	return x+y, x*y
}

func getStrings() (firstString, secString string) {
	firstString = "[ Yes it is "
	secString = "me !]"
	return
}

// ================ MAIN PROGRAM because function "main" in package "main"
func main() {
	var functionLevelVar = "testFunction"
	packageLevelVar = "testPackage"
	const unicodeString = "世界" // constant declaration

	// inside a function ":=" is the shorter form of "var ... [type] ="
	addition, multiplication := addAndMult(2,3)
	fmt.Println("Resultat de l'addition : ", addition)
	fmt.Println("Resultat de la multiplication : ", multiplication)

	fmt.Println(getStrings())
	fmt.Println(functionLevelVar,packageLevelVar, unicodeString)
	fmt.Println(math.Sqrt(float64(addition))) // Type conversion example
}
