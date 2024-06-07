// cheatsheet.go

package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

// Introduction to Go
// See: A tour of Go (tour.golang.org), Go repl (repl.it), Golang wiki (go.dev)

// Hello world
func greetMe(name string) string {
	return "Hello, " + name + "!"
}

func helloWorld() {
	message := greetMe("world")
	fmt.Println(message)
}

// Variables
func variables() {
	var msg string
	var msg2 = "Hello, world!"
	var msg3 string = "Hello, world!"
	var x, y int
	var x2, y2 int = 1, 2
	var x3, msg4 = 1, "Hello, world!"
	msg5 := "Hello"

	var (
		a    int
		b        = 20
		c    int = 30
		d, e     = 40, "Hello"
		f, g string
	)
}

// Constants
func constants() {
	const Phi = 1.618
	const Size int64 = 1024
	const x, y = 1, 2
	const (
		Pi = 3.14
		E  = 2.718
	)
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
}

// Basic types
func basicTypes() {
	str := "Hello"
	str2 := `Multiline
string`

	num := 3          // int
	num2 := 3.        // float64
	num3 := 3 + 4i    // complex128
	num4 := byte('a') // byte (alias for uint8)

	var u uint = 7       // uint (unsigned)
	var p float32 = 22.7 // 32-bit float

	numbers := [...]int{0, 0, 0, 0, 0}
	slice := []int{2, 3, 4}
	slice2 := []byte("Hello")

	a := new(int)
	*a = 234
}

// Type conversions
func typeConversions() {
	i := 2
	f := float64(i)
	u := uint(i)
}

// Flow control
func flowControl(day string, isTired func() bool, doThing func() (int, error), guess func() int) {
	if day == "sunday" || day == "saturday" {
		rest()
	} else if day == "monday" && isTired() {
		groan()
	} else {
		work()
	}

	if _, err := doThing(); err != nil {
		fmt.Println("Uh oh")
	}

	switch day {
	case "sunday":
		// cases don't "fall through" by default!
		fallthrough
	case "saturday":
		rest()
	default:
		work()
	}

	for count := 0; count <= 10; count++ {
		fmt.Println("My counter is at", count)
	}

	entry := []string{"Jack", "John", "Jones"}
	for i, val := range entry {
		fmt.Printf("At position %d, the character %s is present\n", i, val)
	}

	n := 0
	x := 42
	for n != x {
		n = guess()
	}
}

// Functions
func lambdas(x int) {
	myfunc := func() bool {
		return x > 10000
	}
	_ = myfunc()
}

func getMessage() (string, string) {
	return "Hello", "World"
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// Packages
func importing() {
	fmt.Println("Importing packages example")
	_ = rand.Intn(10)
}

func exportingNames() {
	Hello()
}

func Hello() {
	fmt.Println("Hello from exported function!")
}

// Concurrency
func concurrencyExample() {
	ch := make(chan string)

	go push("Moe", ch)
	go push("Larry", ch)
	go push("Curly", ch)

	fmt.Println(<-ch, <-ch, <-ch)
}

func push(name string, ch chan string) {
	msg := "Hey, " + name
	ch <- msg
}

func bufferedChannels() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	ch <- 3
}

func closingChannels() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	for i := range ch {
		fmt.Println(i)
	}

	v, ok := <-ch
	if !ok {
		fmt.Println("Channel closed")
	}
}

func waitGroupExample(itemList []string) {
	var wg sync.WaitGroup

	for _, item := range itemList {
		wg.Add(1)
		go doOperation(&wg, item)
	}
	wg.Wait()
}

func doOperation(wg *sync.WaitGroup, item string) {
	defer wg.Done()
}

// Error control
func errorControl() {
	defer fmt.Println("Done")
	fmt.Println("Working...")

	defer func() {
		fmt.Println("Done")
	}()
	fmt.Println("Working...")

	var d int64
	defer func(d *int64) {
		fmt.Printf("& %v Unix Sec\n", *d)
	}(&d)
	fmt.Print("Done ")
	d = time.Now().Unix()
}

// Structs
type Vertex struct {
	X int
	Y int
}

func structsExample() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X, v.Y)

	v2 := Vertex{X: 1, Y: 2}
	v3 := Vertex{1, 2}
	v4 := Vertex{X: 1}
	_ = v2
	_ = v3
	_ = v4

	v5 := &Vertex{1, 2}
	v5.X = 2
}

// Methods
func (v Vertex) Abs() float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
}

func methodsExample() {
	v := Vertex{1, 2}
	fmt.Println(v.Abs())

	v2 := Vertex{6, 12}
	v2.Scale(0.5)
	fmt.Println(v2)
}

func (v *Vertex) Scale(f float64) {
	v.X = int(float64(v.X) * f)
	v.Y = int(float64(v.Y) * f)
}

// Interfaces
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Length, Width float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

func interfaceExample() {
	var r Shape = Rectangle{Length: 3, Width: 4}
	fmt.Printf("Type of r: %T, Area: %v, Perimeter: %v.\n", r, r.Area(), r.Perimeter())
}

// Main function
func main() {
	helloWorld()
	variables()
	constants()
	basicTypes()
	typeConversions()
	flowControl("sunday", func() bool { return true }, func() (int, error) { return 0, nil }, func() int { return 42 })
	lambdas(10001)
	getMessage()
	split(10)
	importing()
	exportingNames()
	concurrencyExample()
	bufferedChannels()
	closingChannels()
	waitGroupExample([]string{"item1", "item2"})
	errorControl()
	structsExample()
	methodsExample()
	interfaceExample()
}

func rest() {
	fmt.Println("Rest")
}

func groan() {
	fmt.Println("Groan")
}

func work() {
	fmt.Println("Work")
}
