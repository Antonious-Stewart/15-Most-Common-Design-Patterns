package composite

import "fmt"

//Composite is a structural design pattern that lets you compose objects into tree structures and then work with these structures as if they were individual objects.
//Composite became a pretty popular solution for the most problems that require building a tree structure.
//Composite’s great feature is the ability to run methods recursively over the whole tree structure and sum up the results.

//How to Implement
//
//Make sure that the core model of your app can be represented as a tree structure.
//Try to break it down into simple elements and containers.
//Remember that containers must be able to contain both simple elements and other containers.
//
//Declare the component interface with a list of methods that make sense for both simple and complex components.
//
//Create a leaf class to represent simple elements. A program may have multiple different leaf classes.
//
//Create a container class to represent complex elements. In this class, provide an array field for storing references to sub-elements.
//The array must be able to store both leaves and containers, so make sure it’s declared with the component interface type.
//
//While implementing the methods of the component interface, remember that a container is supposed to be delegating most of the work to sub-elements.
//
//Finally, define the methods for adding and removal of child elements in the container.
//
//Keep in mind that these operations can be declared in the component interface.
//This would violate the Interface Segregation Principle because the methods will be empty in the leaf class.
//However, the client will be able to treat all the elements equally, even when composing the tree.

type Soldier interface {
	Brief(orders string)
	Add(component ...Soldier)
}

type Division struct {
	name     string
	brigades []Soldier
}

func NewDivision(name string) *Division {
	return &Division{
		name:     name,
		brigades: make([]Soldier, 0),
	}
}

func (d *Division) Brief(orders string) {
	// should call each brigade and give them order
	message := fmt.Sprintf("Briefing %d Brigades", len(d.brigades))
	for _, brigade := range d.brigades {
		brigade.Brief(orders)
	}
	fmt.Println(message)
}

func (d *Division) Add(brigades ...Soldier) {
	d.brigades = append(d.brigades, brigades...)
}

type Brigade struct {
	name     string
	platoons []Soldier
}

func NewBrigade(name string) *Brigade {
	return &Brigade{
		name:     name,
		platoons: make([]Soldier, 0),
	}
}

func (b *Brigade) Brief(orders string) {
	message := fmt.Sprintf("Briefing %d Platoons", len(b.platoons))

	// should call each platoon and give them order
	for _, platoon := range b.platoons {
		platoon.Brief(orders)
	}
	fmt.Println(message)
}

func (b *Brigade) Add(platoons ...Soldier) {
	b.platoons = append(b.platoons, platoons...)
}

type Platoon struct {
	name   string
	squads []Soldier
}

func NewPlatoon(name string) *Platoon {
	return &Platoon{
		name:   name,
		squads: make([]Soldier, 0),
	}
}

func (p *Platoon) Brief(orders string) {
	message := fmt.Sprintf("Briefing %d Squads", len(p.squads))

	// should call each squad and give them order
	for _, squad := range p.squads {
		squad.Brief(orders)
	}
	fmt.Println(message)
}

func (p *Platoon) Add(squads ...Soldier) {
	p.squads = append(p.squads, squads...)
}

type Squad struct {
	name      string
	enlistees []Soldier
}

func NewSquad(name string) *Squad {
	return &Squad{
		name:      name,
		enlistees: make([]Soldier, 0),
	}
}

func (s *Squad) Brief(orders string) {
	message := fmt.Sprintf("Briefing %d Enlistees", len(s.enlistees))
	// should call each enlistee and give them order
	for _, enlistee := range s.enlistees {
		enlistee.Brief(orders)
	}
	fmt.Println(message)
}

func (s *Squad) Add(enlistees ...Soldier) {
	s.enlistees = append(s.enlistees, enlistees...)
}

type Enlisted struct {
	name string
}

func NewEnlisted(name string) *Enlisted {
	return &Enlisted{
		name: name,
	}
}

func (e *Enlisted) Brief(orders string) {
	// should call each enlistee and give them order
	fmt.Println(orders)
}

func (e *Enlisted) Add(enlistees ...Soldier) {}

//Pros and Cons
//
//You can work with complex tree structures more conveniently: use polymorphism and recursion to your advantage.
//Open/Closed Principle. You can introduce new element types into the app without breaking the existing code, which now works with the object tree.
//
//It might be difficult to provide a common interface for classes whose functionality differs too much.
//In certain scenarios, you’d need to overgeneralize the component interface, making it harder to comprehend.
