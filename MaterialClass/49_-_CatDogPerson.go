package main
import "fmt"

type barker interface {
	bark()
}
type meower interface {
	meow()
}
type runner interface {
	 run()
}

type RunnerBarker interface{
	runner
	barker
}
type RunnerMeower interface{
	runner
	meower
}

type animal struct {
	name string
	price float32
	country string
}

type dog struct {
	animal
	length float32
}
type cat struct {
	animal
	age float32
}


func (c cat)meow(){
	fmt.Println("---------func meow(c cat)-----------------------------")
	fmt.Println("Name:", c.name)
	fmt.Println("Price:", c.price)
	fmt.Println("Country:", c.country)
	fmt.Println("Age: ",c.age)
	fmt.Println("-------------------------------------------------------")
}

func (a animal)run(){
	fmt.Println("---------func run(a animal)-----------------------------")
	fmt.Println("Name:", a.name , " is running")
	fmt.Println("-------------------------------------------------------")
}

func runmeow(pm RunnerMeower) {
	fmt.Println("---------func runmeow----------------------------------")
	pm.run()
	pm.meow()
}

func (d dog)bark(){
	fmt.Println("---------func bark(d dog)-----------------------------")
	fmt.Println("Name:", d.name)
	fmt.Println("Price:", d.price)
	fmt.Println("Country:", d.country)
	fmt.Println("Length: ",d.length)
	fmt.Println("-------------------------------------------------------")
}

func runbark(pb RunnerBarker) {
	fmt.Println("---------func runbark----------------------------------")
	pb.run()
	pb.bark()
}

type cleaner interface {
	clean()
}
type driver interface {
	drive()
}
type killer interface {
	kill()
}
type DriverCleaner interface {
	driver
	cleaner
}

type person struct {
	name string
	yearExperience int
	country string
}


func (p person)clean(){
	fmt.Println(p.name , " is cleaning")
}

func (p person)drive(){
	fmt.Println(p.name , " is driving")
}

func (p person) kill(){
	fmt.Println(p.name , " is killing")
}


func driveclean(dc DriverCleaner) {
	fmt.Println("---------func driveclean----------------------------------")
	dc.drive()
	dc.clean()
}

type DriverKiller interface {
	driver
	killer
}
func drivekill(dk DriverKiller) {
	fmt.Println("---------func drivekill----------------------------------")
	dk.drive()
	dk.kill()
}
type DriverKillerBarker interface {
	DriverKiller
	barker
}
func drivekillbark(dkb DriverKillerBarker) {
	fmt.Println("---------func drivekillbark----------------------------------")
	dkb.drive()
	dkb.bark()
}

func (p person)bark(){
	fmt.Println("---------func bark(p person)-----------------------------")
	fmt.Println("Name:", p.name)
	fmt.Println("Price:", p.yearExperience)
	fmt.Println("Country:", p.country)
	fmt.Println("-------------------------------------------------------")
}

func main() {

	fmt.Println("--------cat section-----------------------------")
	cat1 := cat{animal{"catty",234.45,"iran"},3}
	runmeow(cat1)
    fmt.Println("--------dog section-----------------------------")
	dog1 := dog{animal{"doggy",123.95,"india"},5}
	runbark(dog1)
	fmt.Println("--------driver killer barker section-----------------------------")
	person1 := person{"singh",10,"iran"}

	driveclean(person1)
	drivekill(person1)
	drivekillbark(person1)
}

