package flowcontrol

import (
	"fmt"
	"math"
  "time"
)

var ExportedVariable string = "I'm visible outside the package, because I start with a capital letter"
var packageOnlyVariable string = "life is sad..."

func MyExportedFunction(name string) string {
	return "Ola " + name + "!"
}

func MyforLoop(count int) int {
	result := 0

	for i := 0; i < count; i++ {
		result += i
	}

	return result
}

func MyForLikeWhileLoop(count int) (result int) {

	for result < count { //forever is : "for { ... }"
		result++
	}

	return
}

func MyCondition(toPrint string, shouldBe bool) {
	if shouldBe && toPrint != "" {
		fmt.Println(toPrint)
	} else if !shouldBe && toPrint != "" {
		fmt.Println("I should not but... ", toPrint)
	} else {
		fmt.Println("Nothing to print !")
	}
}

func MySpecialGoCondition(limit int, seed int) {
	if innerVar := int(math.Pow(float64(seed), 3)); innerVar < limit {
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
	default:
		fmt.Println("Too far away.")
	}
}

func MyDeferingStatement() {
	defer fmt.Println("!!")    // will be executed when the surrounding function returns. But its args are executed immediately.
	defer fmt.Println("world") // will be printed before "!!" because defer statements are stacking

	fmt.Println("hello")
}
