package factoryMethod

import "fmt"

// Factory method is a creational design pattern which solves the problem of creating product objects without specifying their concrete classes.

// It’s impossible to implement the classic Factory Method pattern in Go due to lack of OOP features such as classes and inheritance.
//	However, we can still implement the basic version of the pattern, the Simple Factory.

//How to Implement
//
//Make all products follow the same interface. This interface should declare methods that make sense in every product.
//
//Add an empty factory method inside the creator class. The return type of the method should match the common product interface.
//
//In the creator’s code find all references to product constructors. One by one, replace them with calls to the factory method, while extracting the product creation code into the factory method.
//
//You might need to add a temporary parameter to the factory method to control the type of returned product.
//
//At this point, the code of the factory method may look pretty ugly. It may have a large switch statement that picks which product class to instantiate. But don’t worry, we’ll fix it soon enough.
//
//Now, create a set of creator subclasses for each type of product listed in the factory method. Override the factory method in the subclasses and extract the appropriate bits of construction code from the base method.
//
//If there are too many product types and it doesn’t make sense to create subclasses for all of them, you can reuse the control parameter from the base class in subclasses.
//
//For instance, imagine that you have the following hierarchy of classes: the base Mail class with a couple of subclasses: AirMail and GroundMail; the Transport classes are Plane, Truck and Train. While the AirMail class only uses Plane objects, GroundMail may work with both Truck and Train objects. You can create a new subclass (say TrainMail) to handle both cases, but there’s another option. The client code can pass an argument to the factory method of the GroundMail class to control which product it wants to receive.
//
//If, after all of the extractions, the base factory method has become empty, you can make it abstract. If there’s something left, you can make it a default behavior of the method.

type IPhone interface {
	GetOS() string
	TurnOn()
	TurnOff()
}

type Phone struct {
	status string
	os     string
}

func (p *Phone) GetOS() string {
	return p.os
}

func (p *Phone) TurnOn() {
	if p.status == "on" {
		return
	}
	p.status = "on"
	fmt.Println("Turning phone on")
}

func (p *Phone) TurnOff() {
	if p.status == "off" {
		return
	}
	p.status = "off"
	fmt.Println("Turning phone off")
}

type Android struct {
	Phone
}

func NewAndroid() IPhone {
	return &Android{
		Phone: Phone{
			os:     "android",
			status: "off",
		},
	}
}

type Google struct {
	Phone
}

func NewGoogle() IPhone {
	return &Google{
		Phone: Phone{
			os:     "google",
			status: "off",
		},
	}
}
