package main
import "fmt"
import "strconv"
import "strings"

func Valid(id string) bool {
	joined := strings.Replace(id, " ", "", -1)
	oddDigits := ""
	evenDigits := ""

	for i := 0; i < len(joined); i++ {
		if i % 2 == 1 {
			digit, _ := strconv.Atoi(string(joined[i]))
			doubled := digit * 2
			evenDigits += strconv.Itoa(doubled)
		} else {
			oddDigits += string(joined[i])
		}
	}

	sum := sumOfDigits(oddDigits) + sumOfDigits(evenDigits)

	if sum % 10 == 0 {
		return true
	}

	return false


}

func sumOfDigits(digits string) int {
	
	sum := 0
	for _, digit := range digits {
		digitInt, _ := strconv.Atoi(string(digit))
		sum += digitInt
	}

	return sum
}

func main() {
	digits := 4539319503436467
	digitsString := strconv.Itoa(digits)
	fmt.Println(digitsString)

	fmt.Println(sumOfDigits(1234))

	str2 := "My dog name is Dollar"
	str2Splitted := strings.Split(str2, " ")
	fmt.Println(str2Splitted)
	str := "a space-separated string"
	str = strings.Replace(str, " ", "|", -1)
	fmt.Println(str)
}