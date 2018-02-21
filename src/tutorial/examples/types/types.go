package types

import (
  "fmt"
  fcontrol "tutorial/examples/flowcontrol" // non SDK package import defining an alias.
)

// ===================== STRUCTURES ===================

// We can declare structured types
type Point struct {
  X, Y int
}

// Or only types
type myInt int

// ==================== POINTERS ====================

func PointerExample(valueVar int, pointerTest *uint) *int { // using a pointer as parameter allow to modify the inner value
  i := valueVar // assign i to the value of valueVar
  p := &i   // point to i (& generates a pointer to its operand.)
  fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

  return p
}

func StructUsageExample(xIn, yIn int) Point{

  var myPoint Point = Point{xIn, yIn} // or Point{} or Point{X : xIn, Y : yIn} or Point{Y : Yin} ...
  myPoint.X = myPoint.X +2

  // "pointerToPoint := &Point{}" is a convenient way to declare a pointer to a newly created object
  p := &myPoint // pointer to a struct, or "var p *Point = ..."
  p.X = p.X +3 // we should use (*p).X but this is ugly and useless, so there is a shortcut.

  return myPoint
}

// ================= ARRAYS and SLICES ===============

func ArrayAndSliceUsageExample() []string { //the return type is a slice = kind of pointer to an array
  var myArray [4]string           // This is an array. THEY HAVE FIXED SIZE
  myArray[0] = "val0"
  myArray[1] = "val1"
  myArray[2] = "val2"
  myArray[3] = "val3"

  myArrayBetter := [4]int{1, 2, 3, 4}  // This is a more compact way to define an array

  var mySlice []int = myArrayBetter[1:3] //slices are only references to the underlying array. could be a[:], a[1:], a[:4]...
  mySlice2 := []int{1, 2, 3} // create an array then build a slice that references it.
  fmt.Println("Slice with length : %d and Capacity : %d, %s", len(mySlice), cap(mySlice), mySlice2)

  var emptySlice []int
  if(emptySlice == nil){        // The default value for a pointer is nil too.
    fmt.Println("This slice is empty !")
  }

  // The only way to allocate dynamically sized arrays (but not resizeable) :
  dynamicArray := make([]int, 5)
  fmt.Println(dynamicArray)

  // How append elements to the slice :
  var sliceToAppend []int
  sliceWithOneElem := append(sliceToAppend, 42)
  fmt.Println(sliceWithOneElem)

  //Multi dimentional slice
  ticTacToeBoard := [][]string{
		[]string{"X", "O", "X"},
		[]string{"_", "X", "0"},
		[]string{"_", "_", "O"},
	}
  fmt.Println(ticTacToeBoard)

  //iterate over a slice :
  for index, value := range mySlice { // if we didn't care about the index : "for _, value := range mySlice {...}"
    fmt.Printf("Index %d, value : %d", index, value)
  }

  return myArray[:]
}



// ================ MAPS ======================


func MapsExample() map[string]Point { // will return a map having a string as key and a Point as value
  var myMap map[string]Point     // Map declaration
  myMap = make(map[string]Point) //This is a way to create & initialize maps. maps are automatically sized
  myMap["amacan"] = Point{12, 53}

  mySecMap := map[string]Point{ "amausa" : {12, 13}, "amafra" : {1, 3}} // this is the other way, when we already have values.

  delete(mySecMap, "amafra") // delete the value at this index in the map.

  value, isPresent := mySecMap["amafra"] // in this case value = {0 0} because it is the "zero" element of Point. But isPresent = false.

  fmt.Println(myMap["amacan"], value, isPresent)

  return myMap
}


// ================== FUNCTIONS AS TYPE =============

func FunctionsExample(functiontoApply func(int, int) string) string { // Functional programming coming... :p
  var dummyVar string = "dummy appending value but outside of the function"
  innerFunction := func(txt string) string {          // it is possible to declare inner annonymous functions
    return "["+txt+"]"+dummyVar
  }

  return innerFunction(functiontoApply(2, 3)) //we could have returned innerFunction too
}

// We can define functions on types (structured or not!)!
func (p Point) sum() int { // We can also do that EXACT EQUIVALENT : func (p *Point) sum(){...}
  return p.X + p.Y
}

func useTypeFunctionExample(){
  p := Point{1, 2}
  fmt.Println(p.sum()) //we are using the type function here.
}

func UselessFunction(){
  fcontrol.MyExportedFunction("test call from another package") // would be called with flowcontrol.[...] without defined alias
}
