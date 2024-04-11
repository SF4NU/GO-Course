package main

import "log"

//Define Interface
type Human interface {
	UserName() string
	Profession() string
}

//Create user structs
type User1 struct {
	Age    int
	Gender string
}

type User2 struct {
	Age      int
	Gender   string
	Location string
}

//Methods

func (j User1) Profession() string {
	return "Front-End Developer"
}

func (j User1) UserName() string {
	return "John"
}

func (j User2) Profession() string {
	return "Back-End Developer"
}

func (j User2) UserName() string {
	return "Luigi"
}

//Users Information
func PrintInfo(h Human) {
	log.Println("Hi my name is ", h.UserName(), " and I'm a professional ", h.Profession())
}

func main() {
	//instance of structs

	john := User1{
		Age:    20,
		Gender: "male",
	}

	PrintInfo(john)

	luigi := User2{
		Gender:   "male",
		Age:      21,
		Location: "Italy",
	}
	PrintInfo(luigi)
}
