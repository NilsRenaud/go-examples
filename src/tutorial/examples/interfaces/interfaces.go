package interfaces

import "fmt"

type Animal interface {
  speak() string
}

type Pokemon struct {
  trainer string
}

type Dog struct {
  master string
}

func (myPokemon Pokemon) speak() string {
  return "pika !"
}

func (myDog *Dog) speak() string { //On pointer !
  if myDog == nil {
    return "Dog is null !" //Yes, we can call a function on a null object !
  }
  return "Ouaf !"
}

func testInterface(){
  var animal Animal
  pikachu := Pokemon{"Sacha"}
  medor := Dog{"Robert"}

  animal = pikachu //valid
  animal = &medor //also valid, because the pointer implements the function speak.

  fmt.Println(animal)
  // BUT animal = medor is NOT valid although you can use medor.speak() (which is translated as (*medor).speak())
}
