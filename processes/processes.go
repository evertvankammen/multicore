package processes

import (
	"errors"
	"fmt"
	"math"
	"time"
	"unicode/utf8"
)

type person struct {
	name string
	age  int
}

// Interfaces are named collections of method signatures.
type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 { // rect has function area()
	return r.width * r.height
}
func (r rect) perim() float64 { // rect has function perim()
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 { // circle has function area()
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 { // circle has function perim()
	return 2 * math.Pi * c.radius
}

func measure(g geometry) { // input of type geometry must have the declared functions.
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

// hello returns a greeting for the named person.
func hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

func values() {

	fmt.Println("go" + "lang")

	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

func variables() {

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple" // := declaring and assigning a variable
	fmt.Println(f)

	const g = 123
	const h = true
	//arrays
	var i [5]int // array size 5, default value of int is 0
	fmt.Println("emp:", i)
	i[4] = 100
	fmt.Println("set:", i)
	fmt.Println("get:", i[4])

	fmt.Println("len:", len(i))

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
	/*
		Unlike arrays, slices are typed only by the elements they contain (not the number of elements). To create an empty
		slice with non-zero length, use the builtin make. Here we make a slice of strings of length 3 (initially zero-valued).
	*/

	fmt.Println("Slices")

	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c2 := make([]string, len(s))
	copy(c2, s)
	fmt.Println("cpy:", c2)

	l := s[2:5] // s[2] to s[4]
	fmt.Println("sl1:", l)

	twoD2 := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD2[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD2[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD2)

	// map
	fmt.Println("Maps")
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)
	fmt.Println("v1: ", m["k1"])
	delete(m, "k2")
	fmt.Println("map:", m)

	_, prs := m["k2"] // checks existence key, prs is bool _ is blank identifier
	fmt.Println("prs:", prs)

	// range

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}

	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "go" { // iterates over unicode code points
		fmt.Println(i, c)
	}
}

func loopIfElseSWitch() {
	j := 1
	for j <= 9 {
		fmt.Println(j)
		j = j + 1

	}

	for i := 1; i <= 9; i++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	var x = 1

	if x == 1 {
		fmt.Println(x)
	}

	if num := 9; num < 0 { // statement in if
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num)
		fmt.Println(num, "has multiple digits")
	}

	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true) // executes a function
	whatAmI(1)
	whatAmI("hey")

}

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

// multi return
func vals() (int, int) {
	return 3, 7
}

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

// closures
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// recursion
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func zeroVal(ival int) { // passed by value
	ival = 0
}

func zeroPtr(iptr *int) { // int pointer as input, is a memory address
	// *iptr dereferences the pointer from its memory address to the current value at that address.
	// the value on memory address iptr changes on 0
	*iptr = 0 // changes value at the memory
}

func stringsAndRunes() {
	const s = "สวัสดี" // slice of bytes, each char or rune consists of the bytes from Utf-8
	//ส == E0 B8 AA
	//ดี 2 codepoints ด = E0 B8 94 en ี  = E0 B8 B5
	fmt.Println("Len:", len(s))
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:]) //from i to end
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}

	x, y := 0, 0

	print(x, y)
	println()

	word := "Hoi pipeloi"

	println(word[2:])

}

func examineRune(r rune) { // type rune
	if r == 't' { //rune literal
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil

}

func ManyLanguageElements() {
	// Get a greeting message and print it.
	message := hello("Gladys")
	fmt.Println(message)

	values()
	variables()
	loopIfElseSWitch()

	nextInt := intSeq() // sets variable i to 0 and passes this to the enclosed function
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	newInts := intSeq()
	fmt.Println(newInts())

	sum(1, 2)

	fmt.Println(fact(7))

	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(7))
	i := 1
	fmt.Println("initial:", i)
	zeroVal(i)
	fmt.Println("zeroval:", i) // no change to i
	zeroPtr(&i)                // passes the memory address of i.
	fmt.Println("zeroptr:", i)
	fmt.Println("pointer:", &i)

	stringsAndRunes()

	fmt.Println(person{name: "Bob", age: 20})
	fmt.Println(newPerson("Sheldon")) // pointer to struct

	//Go automatically handles conversion between values and pointers for method calls.
	r := rect{width: 10, height: 5}
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())

	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}

	var x person
	x.name = "aaa"
	x.age = 25

}
