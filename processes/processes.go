package processes

import (
	"fmt"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

func Values() {

	fmt.Println("go" + "lang")

	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

func Variables() {

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

func LoopIfElseSWitch() {
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

//closures
func IntSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}