# Chapter 4: Composite Types

In Chapter 3, we looked at basic types in Go. In this chapter, we will look at composite types. In particular, we will look at arrays, slices, maps, and structs.

First, some general points. Arrays and structs are aggregate types: this means that “their values are concatenations of other values in memory” (81). Arrays are homogeneous: all their elements must be of the same type. (The same goes for slices and maps, I think, though the authors don’t say so here.) Structs are heterogeneous: their elements can be of different types. Arrays and structs have fixed size, but slices and maps are dynamic. They grow when you add values to them.

## Arrays

An array in Go is a fixed-length sequence of zero or more elements of one type. You will rarely use arrays directly in Go because they cannot grow. However, arrays form the basis for slices, so you still need to understand them well. Arrays work as you would expect.

```go
var a [3]int            // An array of three integers
fmt.Println(a[0])       // Print the first element of the array
fmt.Println(len(a)-1]   // Print the last element of the array
```

Elements in an array are initialized to their zero value. Thus, all the items in `a` are initially zero. You can iterate over an array using explicit indexing (start at 0 and end at `< len(array)`), or you can use `range` for automatic iteration. Each time through the loop, `range` yields two values: an index number and the item at that index of the array. If you don’t need one of the two values, use `_` to tell the compiler that you won’t be using it anywhere in the loop.

You can initialize an array with an array literal containing a list of values. E.g., `var q [3]int = [3]int{1,2,3}`. You can provide a literal in several ways.

```go
a := [3]{1,2,3}         // explicit declaration of size and all values
a := [3]{1,3}           // values left out are initialized to zero  
a := [...]{1,2,3}       // use ... and Go will infer the size from values
a := [3]{2:3, 0:1, 1:2} // give explicit indexes and add items in any order
a := [3]{2:2, 1:1}      // values left out are initialized to zero
```

The size of an array is part of its type. If two arrays have the same size and type of elements, then they are comparable. However, you cannot compare arrays of different sizes or that contain different elements because they have different types.

Go is a pass by value language. When a function takes an array as a parameter, you face two problems. First, if the array is large, the copying may be costly. Second, any changes you make to the copy are not reflected in the original array. However, if you want to modify an array in a function, you can do so by making the parameter a pointer to the array instead of an array value.

## Slices

Slices are variable-length sequences of elements of one type. You write slices `[]T`, where T is the type. A slice looks like an array without a size.

Slices are based on arrays, though you generally don’t see the connection explicitly. A slice is a pointer, a length, and a capacity. The pointer refers to the an item in an array that underlies the slice. The pointer is the first element in the sequence of the slice, but in the array itself, that same element may or may not be first. The length of the slice says how many elements the slice contains, and the capacity says how many elements it can contain. The length cannot be greater than the capacity. There are built-in functions (`len` and `cap`) to get at the length and capacity of slices.

More than one slice can refer to the same array; these multiple slices can share parts of the array or not. You can use the slice operator to create a new slice: `s[i:j]`. Some rules: 0 <= i <= j <= cap; if you omit i, Go assumes 0 for i; if you omit j, Go assumes `len(s)` for j. You can slice beyond the length of a slice, but not beyond the capacity of a slice. (However, if you slice beyond the length of a slice, you may be surprised at what ends up in those elements.)

When you pass a slice to a function, a copy of the pointer is passed. Therefore, you *can* modify the slices elements—and by association you can modify the elements in the underlying array. Thus, one way to modify an array in place is to pass it to a function as a slice: `a[:]`.

Go does not provide an equality test for slices. The authors explain why two different types of equality check are not provided. A deep equality test would say whether every item in the slice is the same as in another slice. This is difficult because “the elements of a slice are indirect, making it possible for a slice to contain itself. Although there are ways to deal with such cases, none is simple, efficient, and most importantly, obvious” (87). A shallow equality test would check whether two slices refer to the same thing. This would be simple and useful, but it would make `==` ambiguous between arrays and slices. Hence, if you want to compare two slices in Go, you need to do it yourself. (Except, not entirely because Go does provide a deep equal function in the `reflect` package, I think.)

Go does provide one valid comparison for slices: you can always test whether `slice == nil`. However, be careful:

```go
var s []int     // len(s) == 0, s == nil
s = nil         // len(s) == 0, s == nil
s = []int(nil)  // len(s) == 0, s == nil
s = []int{}     // len(s) == 0, s != nil
```

The authors add: “So, if you want to test whether a slice is empty, use `len(s) == 0` and not `s == nil`. Other than comparing equal to `nil`, a nil slice behaves like any other zero-length slice; `reverse(nil)` is perfectly safe, for example. Unless clearly documented to the contrary, Go functions should treat all zero-length slices the same way, whether nil or non-nil” (87).

Here are two useful patterns for in-place changes to slices: reverse and rotate.

```go
// Reverse a slice, here of ints
func reverse(s []int) {
    for i, j = 0, len(s)-1; i < j; i, j = i+1 j-1 {
        s[i], s[j] = s[j], s[i]
    }
}

// Rotate a slice left by n elements.
// First reverse(s[:n]). Then reverse(s[n:]). Then reverse the whole slice.
func rotateLeftByN(s []int, n int) {
    reverse(s[:n])
    reverse(s[n:])
    reverse(s)
}

// Rotate a slice right by n elements.
// First reverse(s[:n]). Then reverse(s[n:]). Then reverse the whole slice.
func rotateRightByN(s []int, n int) {
    reverse(s)
    reverse(s[:n])
    reverse(s[n:])
}
```

### The `append` Function

When working with slices, you will often need to use two built-in functions: `copy` and `append`.

First, `copy`: `copy(target_slice, destination_slice)`. The `copy` function copies items from the destination slice into the target slice. The number of items copied is the minimum of the length of the two slices. `copy` returns the number of items copied, but I think you will usually ignore the return value.

The `append` function has this signature: `append(slice []Type, elems ...Type) []Type`. You must assign the return value of `append`; you will often want to assign the return value to the same slice you are appending to. `append` will increase the capacity of the slice as needed.

## Maps

Go maps are references to hash tables. The type of a map is written `map[k]v`, where `k` is the type of the keys and `v` is the type of the values stored in the map. Keys and values must have one type each, but the key type can differ from the value type. There are restrictions on the type of keys, but no restrictions on the types of values. The keys of a map must be comparable with `==`. Although you *can* compare floating-point numbers with `==`, you should *not* use floating-point numbers as the keys of a map. 

You can create a new map in several ways. First, you can use `make`. Second, you can use a map literal. Compare the following examples.

```go
// First, here is a map literal.
ages := map[string]int{
    "alice": 31,
    "charlie": 30,
}

// The map literal above is equivalent to this use of make.
ages := make(map[string]int)
ages["alice"] = 31
ages["charlie"] = 30
```
You can update the value of a key by assigning a new value. E.g., after a birthday, `ages["charlie"] = 31`. You can also remove entries from a map using `delete`: `delete(ages, "alice")`.

It is safe to look up a key, assign a value to a key, or delete a key/value pair by key, *even if the key is not yet in the map*. As the authors say: “a map lookup using a key that isn’t present returns the zero value for its type. For example, if "bob" is not yet a key in the `ages` hash, you can still do this: `ages["bob"] += 1`.

However, you cannot take the address of a map element. Thus `foo := &ages["bob"]` is not valid.

You can iterate over a map using `range`. The order of iteration is unspecified. If you want to iterate in a specific order, you need to store the keys in an slice, sort the slice, and then use the slice items as keys to iterate over the map.

You can create a map that is `nil`, the zero value for reference types. For example, `var ages map[string]int`. That declaration is not the same as using `make` or a map literal to initialize the map. Although most operations (lookup, `delete`, `len`, and `range` loops) are safe, you will cause a panic if you try to store an element in a `nil` map. In general, I can’t see any reason to create a map without `make` or a literal.

When you access an element by subscripting (e.g., ages["whatever-name-here"]), you *always* get some value. Thus, Go presents an ambiguity. The age you get may be zero because the person is zero years old. Or the age you get may be zero because the person is not yet in the map. To distinguish between these two cases, accessing an element in a Go map always returns two values—though you can ignore the second one in many cases.

You cannot compare maps using `==`. If you want to compare maps, you need to write a custom function. For example:

```go
func equal(x, y map[string]int) bool {
    if len(x) != len(y) {
        return false
    }
    for k, xv := range x {
        if yv, ok := y[k]; !ok || yv != xv {
            return false
        }
    }
    return true
}
```

You can use maps with boolean values as sets. Consider the following:

```go
seen := make(map[string]bool)
input := bufio.NewScanner(os.Stdin)
for input.Scan() {
    line := input.Text()
    if !seen[line] {
        seen[line] = true
        fmt.Println(line)
    }
}
if err := input.Err(); err != nil {
    fmt.Fprintf(os.Stderr, "program: %v\n", err)
    os.Exit(1)
}
```

Although you can only use comparable types as keys, you can get around this. For example, if we want to use slices as keys, we can stringify the slices via a function and then use the function to add elements. Here’s an example from the book:

```go
var m = make(map[string]int)

func k(list []string) string {
    return fmt.Sprintf("%q", list)
}

func Add(list []string) {
    m[k(list)]++
}

func Code(list []string) int {
    return m[k(list)]
}
```

In this code, the `Add` function tracks how many times a given list of strings is passed into that method. The `Count` function looks up the number of times a slice has been called with `Add`.

## Structs

A struct is an aggregate data type that bundles together zero or more named values of arbitrary types as a single thing. Each value is called a *field*. You declare a struct type by listing its name and the names and types of its fields. For example, here’s a student struct type.

```go
type student struct {
    email string
    firstName string
    lastName string
    advisor teacher // This field has a type that is itself a struct.
    graduatingClass int
    onProbation bool
    // etc
}
```

You can access the individual fields of a variable of a given struct type using dot notation. For example, `var s student; s.email = "foo@bar.edu"`. Notice that the fields of a struct are also variables: you can assign to a field (`s.advisor = someTeacher`) or take the address of a field and access it through a pointer (`firstName = &s.firstName; *firstName += "foo"`).

You can use dot notation with a pointer to a struct, and Go simplifies things for you.

```go
var employeeOfTheMonth *Employee = &dilbert
employeeOfTheMonth.Position += " (proactive team player)"
// equivalent to (*employeeOfTheMonth).Position += ...
```

The authors explain that sometimes you should return a pointer to a struct from a function in order to make the return value assignable.

```go
func EmployeeByID(id int) *Employee { /* ... */ }
id := someone.ID
EmployeeByID(id).salary += 1000
```

As the authors explain, the assignment would not compile if `EmployeeByID` returned an `Employee` rather than a pointer. Why not? Because a pointer to a struct type is a variable, but an unassigned struct (a raw `Employee` not yet stored in a variable) is not a variable and not assignable.

A named struct cannot declare a field of its own type: “an aggregate value cannot contain itself” (101). However, a struct of type S can have a field of the type `*S` (pointer to S). This makes it easy to create recursive data structures like linked lists or trees.

```go
type tree struct {
    value int
    left, right *tree
}
```

### Struct Literals

You can write the value of a struct type using a struct literal. For example:

```go
type Point struct {
    X int
    Y int
}
p := Point{1, 2}
q := Point{Y: 4}
```

The first form requires the user to provide values for all fields in the order fields appear in the struct type’s declaration. The second type is more flexible. You can put fields in any order, and you can leave fields out. (If you leave a field out, that field takes the zero value for its type.)

You can return a struct value from a function and you can pass one to a function as an argument. For the sake of efficiency, you should use pointers to pass or receive large structures. You *must* use a pointer if you wish to modify the struct inside the function.

### Comparing Structs

You can compare structs if all the fields are comparable. If a struct type is comparable, then you can use it as the key for a map.

### Struct Embedding and Anonymous Fields

First, you can embed one struct inside another. For example, consider the following.

```go
type Point struct {
    X, Y int
}

type Circle struct {
    Center Point
    Radius int
}

type Wheel struct {
    Circle Circle
    Spokes int
}
```

This is clear in some ways, but you must access fields of a wheel explicitly. E.g., `var w Wheel; w.Circle.Radius = 8`. That’s painful.

Alternatively, you can declare a field in a struct using a type but no name. Thus, these are called *anonymous fields*. Here’s an alternative to the previous shapes code.

```go
type Circle struct {
    Point
    Radius int
}

type Wheel struct {
    Circle
    Spokes int
}
```

Now you can refer to fields without the intervening names. E.g., `var w Wheel; w.Radius = 8`. However, now you cannot use simple struct literals to create instances.

```go
w := Wheel{8,8,5,20} // Will not compile
w := Wheel{X: 8, Y: 8, Radius: 5, Spokes: 20} // Also will not compile
```

Instead, as the authors explain, “the struct literal must follow the shape of the type declaration, so we must use one of the two forms below, which are equivalent to each other” (106).

```go
w := Wheel{Circle{Point{8,8}, 5}, 20}
w := Wheel{
    Circle: Circle{
        Point: Point{X:8, Y:8},
        Radius: 5, // The trailing comma is required
    },
    Spokes: 20, // The trailing comma is required
}
```

## JSON

Go provides rich support for JSON and other encoding standards (e.g., XML and Google’s Protocol Buffers). JSON has these basic types: numbers, booleans, and strings. JSON strings are Unicode code points in double quotes. You can specify characters using `\uhhhh` escapes (UTF-16 codes).

JSON also has arrays and objects. A JSON array is an ordered sequence of values, written as a comma-separated list inside square brackets. JSON arrays encode arrays and slices in Go. A JSON object is a collection of `name:value` pairs, separated by commas and surrounded by braces. JSON objects represent Go maps and structs.

In order to convert JSON to Go or Go to JSON, you need to define structs in Go with the proper fields. Since the naming conventions for JSON and Go are different, you can add field tags that go uses to help convert data between JSON and Go.

```go
type Movie struct {
    Title string
    Year int `json:"released"`
    Color bool `json:"color,omitempty"`
    Actors []string
}
```

The `Year` and `Color` fields use different casing in JSON and Go. In addition, if a movie’s color field is false, we omit it from the JSON. For more information about field tags, see [the Go documentation for the JSON encoding package](https://pkg.go.dev/encoding/json).

To convert Go data structures into JSON, we use `json.Marshal`. This method takes one argument, a Go data structure. The method returns a byte-slice (e.g., a string) and a possible error. The byte-slice that `Marshal` returns is ultra-compact: no newlines and no extra whitespace. This format is great for machines, but not for human readers. If you want something more people-friendly, try `json.MarshalIndent`. This method takes two further arguments (in addition to the data structure). The first argument defines a string prefix for each line of output; you can always use the empty string. The second argument is a string to define what to add for each level of indentation. An actual call might look like this: `json.MarshalIndent(dataStructure, "", "    ")`.

To convert JSON into a Go data structure, we use `json.Unmarshal`. It takes two arguments: a byte-slice (e.g., a string) and a pointer to a Go data structure where the data will be stored. It returns an error, if there was a problem, or `nil`.

## Text and HTML Templates

The authors provide a compact but rich introduction to both text and HTML templating in Go. I’ll come back to this another time.
