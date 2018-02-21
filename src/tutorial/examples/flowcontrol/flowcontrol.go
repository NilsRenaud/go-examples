package flowcontrol

//Another way to import things
import (
	"fmt"
	"math"
  "time"
)

var ExportedVariable string = "I'm visible outside the package, because I start with a capital letter"
var packageOnlyVariable string = "life is sad..." // Not exported because does not start with capital letter.

// Same thing for the functions, when they start with a capital letter they are "public"
func MyExportedFunction(name string) string {
	return "Ola " + name + "!"
}

// ================== LOOPS ================

func MyforLoop(count int) int {
	result := 0

	for i := 0; i < count; i++ { //For loop example
		result += i
	}

	return result
}

func MyForLikeWhileLoop(count int) (result int) {

	for result < count { 	// There is no "while", this is the way to do a while
		result++						// forever loop looks like : "for { ... }"
	}

	return // Wait... it should return something ! YES, BUT the output has been names in the signature,
				 // so the variable is defined in the whole function, no need to repeat ourself. look at the function again
}

// ================== CONDITIONS =============

func MyCondition(toPrint string, shouldBe bool) {

	if shouldBe && toPrint != "" {                   // Nothing special here...
		fmt.Println(toPrint)
	} else if !shouldBe && toPrint != "" {
		fmt.Println("I should not but... ", toPrint)
	} else {
		fmt.Println("Nothing to print !")
	}

}

func MySpecialGoCondition(limit int, seed int) {

	if innerVar := int(math.Pow(float64(seed), 3)); innerVar < limit { //Yes we can define some scoped variables
		fmt.Printf("Value allowed : %d", innerVar)
	} else {
		fmt.Println("Value not allowed : %d", innerVar) //accessible here too !
	}
	// here innerVar is no more defined !
}

func MySwitchStatement() {
	fmt.Println("When's Saturday?")

	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0: // we can put variable, functions or constants here. Even a condition like "switch {case i == 5 ...}"
		fmt.Println("Today.") // no need to add break, Go will only execute the first matching case.
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:                      // We can also use a switch to discriminate depending on a type
		fmt.Println("Too far away.")
	}
}

func MyDeferingStatement() {
	defer fmt.Println("!!")    // will be executed when the surrounding function returns. But its args are executed immediately.
	defer fmt.Println("world") // will be printed before "!!" because defer statements are stacking

	fmt.Println("hello")
}
