package main

import (
    "fmt"
    "math/rand"
)

// Functions - parameters type
func add(x, y int) int {
    return x + y;
}

// Return multiple values
func swap(x, y string) (string, string) {
    return y, x;
}

// Named Return Values
func split(number int) (x, y int) {
    x = number * 4/9
    y = number - x
    return // Naked return
}

// Variables
var c, python, java bool

// Varibale Declaration
var learn string = "Learn"

// Data Types
/*
bool

string

int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64

float32 float64

complex64 complex128

Use int for numbers unless you have a specific reason to use int16 or uint

Implicit types are derived based on the highest precision
*/

// const - Implicit type is of highest precision - ie 3 - int, 3.14 - float64, 2 + 3i - complex128
const Pi = 3.14

//struct
type Vertex struct {
    X int
    Y int
}


func main() {
    fmt.Println(rand.Intn(2));
    fmt.Println(add(34,57));
    a, b := swap("Hello", "World")
    fmt.Println(a, b);
    fmt.Println(split(17))

    // Variables
    var i int
    fmt.Println(i, c, python, java)

    // Variable declaration - can omit type, type inferred from declaration
    var j, k, isFun = 1, 2, true
    fmt.Println(j, k, learn, isFun)

    // Short Variable Declaration - using :=, cannot be used outside functions
    short, num := "Short Variable", 7
    fmt.Println(short, num);

    fmt.Println(Pi);
    fmt.Printf("Type of %v is %T\n", Pi, Pi);

    // Loops - only has for loop
    sum := 0
    for i := 0; i <= 5; i++ {
        sum += i;
    }
    fmt.Println(sum);

    // init and post statements are optional
    sum = 1;
    for ; sum < 100 ; {
        sum += sum
    }
    fmt.Println(sum);

    //while loop - just drop the init and post statements
    sum = 1
    for sum < 10 {
        sum += sum;
    }
    fmt.Println(sum);

    // Pointers
    f := 42;
    p := &f;
    fmt.Println(p);

    // struct
    v := Vertex{1,2}
    fmt.Println(v);
    v.X = 4
    fmt.Println(v);

    // you could also have pointers to struct
    structPointer := &v;
    structPointer.X = 5;   // Note that we are not using (*structPointer).X - ease of use
    fmt.Println(v);
    fmt.Println(structPointer);

    //Arrays
    var array [10]int;
    fmt.Println(array);
;
}