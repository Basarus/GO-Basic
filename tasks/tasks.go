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
