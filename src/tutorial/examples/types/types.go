package types

import (
  "fmt"
  fcontrol "tutorial/examples/flowcontrol"
)

// We can declare structured types
type Point struct {
  X, Y int
}

// Or only types
type myInt int


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

func ArrayAndSliceUsageExample() []string {
  var myArray [4]string
  myArray[0] = "val0"
  myArray[1] = "val1"
  myArray[2] = "val2"
  myArray[3] = "val3"

  myArrayBetter := [4]int{1, 2, 3, 4}

  var mySlice []int = myArrayBetter[1:3] //slices are only references to the underlying array. could be a[:], a[1:], a[:4]...
  mySlice2 := []int{1, 2, 3} // create an array then build a slice that references it.
  fmt.Println("Slice with length : %d and Capacity : %d, %s", len(mySlice), cap(mySlice), mySlice2)

  var emptySlice []int
  if(emptySlice == nil){
    fmt.Println("This slice is empty !")
  }

  // The only way to allocate dynamically sized arrays :
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

func MapsExample() map[string]Point {
  var myMap map[string]Point
  myMap = make(map[string]Point) //This is a way to create & initialize maps
  myMap["amacan"] = Point{12, 53}

  mySecMap := map[string]Point{ "amausa" : {12, 13}, "amafra" : {1, 3}} // this is the other way, when we already have values.

  delete(mySecMap, "amafra")

  value, isPresent := mySecMap["amafra"] // in this case value = {0 0} because it is the "zero" element of Point. But isPresent = false.

  fmt.Println(myMap["amacan"], value, isPresent)

  return myMap
}

func FunctionsExample(functiontoApply func(int, int) string) string {
  var dummyVar string = "dummy appending value but outside of the function"
  innerFunction := func(txt string) string {
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
  fcontrol.MyExportedFunction("test call from another package")
}
