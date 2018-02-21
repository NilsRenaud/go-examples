package main

// ========== Import section
import "fmt"       //Come from the Go SDK
import "math"
// import "github.com/go-chef/chef" after a "go get" on it.

// Package level var & multiple initializer & type-inference
var packageLevelVar string
var i, name = 1, "nils"

// Multi var declaration
var (
		myBool bool = false
		myInt int = -1000 // also existing : int  int8  int16  int32  int64 and rune=int32. Default : 0
		myUnsignedInt uint = 1000 // also existing uint uint8 uint16 uint32 uint64 uintptr and byte=uint8. Default : 0
		myFloat float64 = 3.2 //also existing float32 float64. Default : 0
		myString string = "string"// Default : ""
		myComplex complex64 = -5 + 12i //also existing complex64, complex128. Default : 0
)

// A simple function
func getStrings() (firstString, secString string) {
	firstString = "[ Yes it is "
	secString = "me !]"
	return
}

// A simple function returning 2 values
func addAndMult(x, y int) (int, int) {
	return x+y, x*y
}

// ================ MAIN PROGRAM because function "main" in package "main"
func main() {
	var functionLevelVar = "testFunction" // function level variable
	packageLevelVar = "testPackage"       // This one was previously declared
	const unicodeString = "世界"          // constant declaration

	// inside a function (here main) ":=" is the shorter form of "var ... [type] ="
	addition, multiplication := addAndMult(2,3)
	fmt.Println("Addition result : ", addition)
	fmt.Println("Multiplication result : ", multiplication)

	fmt.Println(getStrings())
	fmt.Println(functionLevelVar,packageLevelVar, unicodeString)
	fmt.Println(math.Sqrt(float64(addition))) // Type conversion example and method from external package usage.
}
