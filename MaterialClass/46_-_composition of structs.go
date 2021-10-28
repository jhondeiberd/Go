/*
Composition is a method in which Larger objects with a wider functionality are embedded
with smaller objects with specific behaviors. in inheritance, objects inheriting features
from a parent class/object. in composition, larger objects are composed of the other objects and
can use their functionality.
writing code in the Go Language is around using structs and interfaces so composition is a key.
*/

//composition of structs : example 1

package main
import "fmt"

// We create a struct details
type details struct {
	genre	 string
	genreRating string
	reviews	 string
}

// We create a struct game
type game struct {
	name string
	price string
	//this is composition of structs
	details
}

// this is a method defined
// on the details struct
func (d details) showDetails() {
	fmt.Println("Genre:", d.genre)
	fmt.Println("Genre Rating:", d.genreRating)
	fmt.Println("Reviews:", d.reviews)
}

// this is a method defined
// on the game struct
// this method has access
// to showDetails() as well since
// the game struct is composed
// of the details struct
func (g game) show() {
	fmt.Println("Name: ", g.name)
	fmt.Println("Price:", g.price)
	fmt.Println("Price:", g.genre)
	g.showDetails()
}

func main() {

	// defining a struct
	// object of Type details
	action := details{"Action","18+", "mostly positive"}

	// defining a struct
	// object of Type game
	newGame := game{"batman","$25", action}

	newGame.show()
}

