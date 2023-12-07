package tasks

import "fmt"

func HelloUser(year int) string {
	switch {
	case 1964 >= year && year >= 1946:
		return "Привет, бумер!"
	case 1965 <= year && year <= 1980:
		return "Привет, представитель Х!"
	case 1981 <= year && year <= 1996:
		return "Привет, миллениал!"
	case 1997 <= year && year <= 2012:
		return "Привет, зумер!"
	case 2013 < year:
		return "Привет, альфа!"
	default:
		return "Unknown"
	}
}

func IncB(b int) string {
	switch b {
	case 1:
		fmt.Println("Один")
		fallthrough
	case 3:
		fmt.Println("Три")
	default:
		fmt.Println("unqualified")
	}
	return "OK"
}

func Cycle(a int, b int) bool {
	v := a
	for i := 1; i < b; i++ {
		v++
		fmt.Println(v)
	}

	array := [3]int{1, 2, 3}
	for arrayIndex, arrayVaalue := range array {
		fmt.Printf("array[%d]: %d\n", arrayIndex, arrayVaalue)
	}

	return true
}

func GoTo() bool {
outerLoopLabel:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("[%d, %d]\n", i, j)
			break outerLoopLabel
		}
	}
outerLoopLabel2:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("[%d, %d]\n", i, j)
			continue outerLoopLabel2
		}
	}
	fmt.Println("End")
	return true
}

func FizzBuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
