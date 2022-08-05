### Read input dari console (stdin)
```
import "fmt"
import "bufio"
import "os"
reader := bufio.NewReader(os.Stdin)
fmt.Print("Enter text: ")
text, _ := reader.ReadString('\n')
fmt.Println(text)
```
### Ubah string ke int
```
import "strconv"
s := "123"
i, err := strconv.Atoi(s)
if err != nil {
    // gagal
    panic(err)
}
fmt.Println(s, i)
```

### Ubah int ke string
```
t := strconv.Itoa(123)
fmt.Println(t)
```

### Print type of variable
```
import "reflect"
tst := "string"
tst2 := 10
tst3 := 1.2
fmt.Println(reflect.TypeOf(tst))
fmt.Println(reflect.TypeOf(tst2))
fmt.Println(reflect.TypeOf(tst3))
```

### Lower & capitalize string
```
import "strings"
s := "INTEGRATED ENGINEERING 5 Year (BSC with a Year in Industry)"
fmt.Println(strings.Title(strings.ToLower(s)))
```

### Iterate string
```
text := "abcdefg"
for _, char := range(text) {
    fmt.Println(string(char))
}
```

### Function with multiple returns
```
a, b := getMessage()
func getMessage() (a string, b string) {
  return "Hello", "World"
}
```

### Return error & return multiple values
```
import "errors"
func Distance(a, b string) (int, error) {
    if len(a) != len(b) {
        return 0, errors.New("Different length")
    }
```

### "In" punyanya python buat string
```
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
```

### "In" punyanya python buat map
```
dict := make(map[string]bool)

value, error = dict['test']

if value == nil {
    dict['test'] = 1
} else {
    dict['test']++
}

```

### Declare map
```
map_1 := make(map[int]string)   // empty
map_2 := map[int]string {       // with init values
        90: "Dog",
        91: "Cat",
        92: "Cow",
        93: "Bird",
        94: "Rabbit",           // --> ada koma ga kek dipython
}
```

### Iterate map
```
    dict := map[string]int {
        "AEIOULNRST" : 1,
        "DG" : 2,
        "BCMP": 3,
        "FHVWY": 4,
        "K": 5,
        "JX": 8,
        "QZ": 10,
    }

    for _, char := range word {           
        for key, value := range dict {                 // kv pake range
            uppered := strings.ToUpper(string(char))
            if contains(key, uppered) {
                score += value
            }
        }
    }
    
```


### isalpha(), isnumeric(), isalnum() punya python
```
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

func main() {
	s := "a"
	f := "1"
	b := "?"
	fmt.Println(isalnum(s))
	fmt.Println(isalnum(f))
	fmt.Println(isalnum(b))
	
}
```

### String to map (Counter), bisa atur ignore case/ga
```
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
```

### Convert float64/32 to int
```
var v float32 = 2.1
var x float64 = 5.7
var y int = int(x)
fmt.Println(y)  // outputs "5"
```

### Convert int to float64/32
```
var y int = 8
var x float32 = float32(y)
var z float64 = float64(y)
fmt.Println(x)
fmt.Println(z)
``` 

### Sum of digits
```
func sumOfDigits(digits int) int {
	digitsString := strconv.Itoa(digits)
	sum := 0
	for _, digit := range digitsString {
		digitInt, _ := strconv.Atoi(string(digit))
		sum += digitInt
	}

	return sum
}
```

### Split string
```
str2 := "My dog name is Dollar"
str2Splitted := strings.Split(str2, " ")
```

### Replace char in string
```
import "strings"
str := "a space-separated string"
str = strings.Replace(str, " ", "|", -1)
fmt.Println(str)
```

### Strings package
```
Contains("test", "es")				true
Count("test", "t")				2
HasPrefix("test", "te")				true
HasSuffix("test", "st")				true
Index("test", "e")				1
Join([]string{"a", "b"}, "-")			a-b
Repeat("a", 5)					aaaaa
Replace("foo", "o", "0", -1)			f00
Replace("foo", "o", "0", 1)			f0o
Split("a-b-c-d-e", "-")				[a b c d e]
ToLower("TEST")					test
ToUpper("test")					TEST
```

### Struct
```
type Vertex struct {
	X int
	Y int
}
v := Vertex{X: 1, Y: 2}
v := Vertex{1, 2}

```
#### Struct method
```
func (v Vertex) Abs() float64 {
  return math.Sqrt(v.X * v.X + v.Y * v.Y)
}
v := Vertex{1, 2}
v.Abs()
```