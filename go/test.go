package main
import "strings"
import "fmt"

func isAlpha(s string) bool {
	if len(s) > 1 {
		panic("panjangnya harus 1")
	}
	const alpha = "abcdefghijklmnopqrstuvwxyz"
	lowered := strings.ToLower(string(s))
	return strings.Contains(alpha, lowered) 
}

func isNumeric(s string) bool {
	if len(s) > 1 {
		panic("panjangnya harus 1")
	}
	const numbers = "1234567890"
	return strings.Contains(numbers, s) 
 }

func isalnum(s string) bool {
	return isAlpha(s) || isNumeric(s) 
}

func stringToMap(word string, ignoreCase bool) map[string]int {
	dict := make(map[string]int)

	for _, char := range word {
		if ignoreCase {		// kecil gede diperhatiin
			lowered := strings.ToLower(string(char))
			dict[lowered]++
		} else {
			dict[string(char)]++
		}
	}

	return dict

}

func main() {
	/*
	fmt.Println(stringToMap("TtTEe", false))
	fmt.Println(strings.ToLower("1"))
	fmt.Println(strings.ToLower("\\"))

	var rate float64 = 90

	fmt.Println(float64(rate * 45))

	
	var y int = 8
	var x float32 = float32(y)
	var z float64 = float64(y)
	fmt.Println(x)
	fmt.Println(z) 
	*/
	fmt.Println(string(3))

	/*
	value, _ := dict['test']

	if value == nil {
		dict['test'] = true
	} else {
		dict['test']++
	}

	fmt.Println(dict)
	*/
	
}