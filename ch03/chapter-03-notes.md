# Chapter 3: Basic Data Types

“Go’s types fall into four categories: *basic types*, *aggregate types*, *reference types*, and *interface types*” (51). Basic types are numbers, strings, and booleans. Aggregate types are arrays and structs. Reference types are pointers, slices, maps, functions, and channels. Interface types provide a kind of duck typing by creating a space for “any type that satisfies an interface.”

## Integers

Go has four sizes of integers (8, 16, 32, and 64 bits) in both signed and unsigned versions. You can also (and generally should) ask for simply `int` or `uint`. The compiler will choose the most natural or efficient size for a given platform. (It will be either 32 or 64 bits, but we don’t need to chare which in general.)

There are also several alias types based on integers. The `rune` type is an `int32`: it generally indicates that a variable is a Unicode code point. You can use `rune` or `int32` interchangeably. The `byte` type is a synonym for `uint8`. You should use `byte` to indicate that the value is “a piece of raw data rather than a small numeric quantity” (52). Last, there is `uintptr`, which is of an unspecified size, but which is guaranteed to be large enough “to hold all the bits of a pointer value” (52).

Binary operators for math, logic, and comparison have five levels of precedence. The list below goes in decreasing order of precedence.

+ `*`, `/`, `%`, `<<`, `>>`, `&`, `&^`
+ `+`, `-`, `|`, `^`
+ `==`, `!=`, `<`, `<=`, `>`, `>=`
+ `&&`
+ `||`

Operators at the same level of precedence associate left to right. Use parentheses as needed for clarity or to get things to do what you want. (E.g., `mask & (1 << 28)`.)

Every operator in the first two lines has a corresponding assignment operator. E.g., `*` and `*=`.

The modulus operator only applies to integers. The sign of the result follows the sign of the dividend: `-5 % 3` returns -2. Integer division truncates towards zero.

The following are bitwise binary operators:

+ `&`   bitwise AND
+ `|`   bitwise OR
+ `^`   bitwise XOR
+ `&^`  bitwise AND NOT (= bit clear)
+ `<<`  left shift
+ `>>`  right shift

The AND, OR, and XOR (= exclusive or) operators provide logical operations. AND returns 1 if both its operands are 1; otherwise it returns 0. OR returns 1 if either or both of its operands are 1; it returns 0 if both of its operands are 0. XOR returns 1 if exactly one of its operands is 1, but if neither or both of its operands are 1, it returns 0.

When used as a unary operator `^` is bitwise negation or complement; it returns a value with each bit in its operand flipped.

The AND NOT operator works as follows: “in the expression `z = x &^ y`, each bit of z is 0 if the corresponding bit of y is 1; otherwise it equals the corresponding bit of x” (53).

When you use the two bit shift operators, the right-hand operand must be positive. It determines how many times to shift the bit positions. The left-hand operand can be positive or negative. However, if you want to use an integer as a bit pattern, you should always use an unsigned integer. Otherwise, right shifting will fill vacated bits with copies of the sign bit.

The authors mention that Go programmers generally use integers rather than unsigned integers, except when bitwise math is involved or in specialized domains such as binary files, hashing, and cryptography.

You can write integer literals as decimal, octal, or hexadecimal numbers. Octal numbers begin with 0, and hexadecimal numbers begin with 0x or 0X. Hex digits can be upper or lower case. People tend to use octal numbers only to represent file permissions on POSIX systems. People use hexadecimal numbers “to emphasize the bit pattern of a number over its numeric value” (55).

Rune literals are a single Unicode code point within single quotes. You print runes with the `%c` or `%q` verbs. The `%q` verb quotes the result.

### Some `fmt` Tricks

```go
o := 0666
fmt.Printf("%d %[1]o %#[1]o %[1]b\n", o)

x := int64(0xdeadbeef)
fmt.Printf("%d %[1]x %#[1]x %#[1]X %[1]b\n", x)
```

You can print an integer in decimal, octal, or hexadecimal by choosing different verbs (`d`, `o`, and `x` or `X` respectively). If you add the adverb `#` after `%`, then Go will print `0` before octal numbers and `0x` or `0X` before hexadecimal numbers (depending on whether you use the lower-case or upper-case `x` verb). Finally, instead of repeating the variable over and over here, we use the `[n]` adverb after `%` to tell `Printf` to use the first operand repeatedly.

## Floating-Point Numbers

Go provides two floating point types: `float32` and `float64`. You can find their limits using the `math` package, but in general a `float32` provides approximately six decimal digits of precision, and a `float64` provides about fifteen decimal digits of precision. The authors recommend using `float64` since the smaller floats “accumulate error rapidly…and the smallest positive integer that cannot be exactly represented as a `float32` is not large” (56).

You can use `math.IsNaN` to check whether something “is a not-a-number value” (57). Do not check for equality with `math.NaN` “because any comparison with NaN *always* yields `false` (except `!=` which is the negation of `==`)” (57). You should use the `result float64, ok bool` pattern when writing functions for floats where the result might fail. Do not return `math.NaN` as a result.

## Complex Numbers

Go provides two complex number types: `complex64` and `complex128`. Nobody uses them much, and I won’t ever want them. (I read somewhere that they may leave the language at some point.)

## Booleans

There are two boolean values: `true` and `false`. Conditions in `if` and `for` statements are boolean tests. Comparison operators (such as `==` and `>`) produce a boolean result. Unary `!` inverts a boolean value.

You can combine boolean values with the `&&` (AND) and `||` (OR) operators. These operators have short-circuit behavior. If the answer to a test is determined by the value of the left operand, Go will not evaluate the right operand. The authors give the following as an example: `s != "" && s[0] == 'x'`. Since `s[0]` will panic at runtime if `s` is the empty string, the short-circuit behavior of `&&` saves you here.

The `&&` operator has higher precedence than `||` (the authors remind us that `&&` is boolean multiplication and `||` is boolean addition). Therefore, you can leave off parentheses in some useful cases. They give this example:

```go
// Test for an ASCII letter or digit
if 'a' <= c && c <= 'z' ||
    'A' <= c && c <= 'Z' ||
    '0' <= c && c <= '9' {
// Something goes here
}
```

## Strings

Strings are immutable sequences of bytes. In general, a string is interpreted as a UTF-8-encoded sequence of Unicode code points (runes).

The built-in `len` function returns how many bytes are in a string, and the index operation (`s[i]`) retrieves the *i*-th byte of string `s` where zero is less than or equal to i, and i is less than `len(s)`.

Beware: `len` and the index operation work on the *bytes* in a string not the runes. But some Unicode code points require multiple bytes. Thus, you can’t use `len` or indexing to get the results you (probably) want in many cases.

You can also use the substring operation `s[i:j]` to get a new string from an existing string. The i index cannot be less than zero, and the result will run from the i index to one less than j. The result will contain j-i bytes. (Go will panic at runtime if either i or j is out of bounds or if j is less than i.) You can omit the i operand, in which case Go assumes an i of 0. You can omit the j operand, in which case Go assumes a j of `len(s)`.

You can concatenate two strings with `+`, and you can compare them with `==`, `!=`, and the other comparison operators. String comparison is byte by byte, and the result is “natural lexicographic ordering” (65).

Strings are immutable. If you assign a new value to a string variable, you create a new string. Consider this example:

```go
s := "left foot"
t := s
s += ", right foot"
// s == "left foot, right foot"
// t == "left foot"
```

Because strings are immutable, you cannot modify a string’s data in place. `s[0] = 'L' // compile error: cannot assign to s[0]`. (Note to self: I’m not sure how this lines up with their in-place slice examples later.)

### String Literals

You can write a string literal in several ways. First, you can put a sequence of bytes between double quote marks. This is a string literal. Inside of a string literal, you can include various escape sequences. For example '\n' and '\t'. You can also include arbitrary bytes written as octal or hexadecimal escapes. Octal escapes have a backslash and three octal digits (0 to 7). Hexadecimal escapes have a backslash, an `x` and two hexadecimal digits. (You can use upper- or lower-case letters.)

You can also use back quote marks instead of double quote marks for a raw string literal. Inside of a raw string literal, escape sequences are not processed, but you can include newlines and backslashes without extra work. Raw string literals are useful for writing regular expressions (because they contain many backslashes) and JSON literals. They are also helpful for writing usage messages for command line applications.

### Unicode

Unicode provides a way to organize characters from languages all over the world, as well as accents, diacritical marks, control codes (tabs, carriage returns), and more. Each item gets a standard number as its Unicode code point. Go calls Unicode code points *runes*.

There are over 120,000 code points in Unicode 8. Go stores a single rune as an `int32`. The word `rune` is an alias for `int32`. One obvious way to store sequences of runes is as sequences of `int32` values. This is simple in one sense, but it takes up a lot of excess space because an enormous amount of text requires only 8 bits or 1 byte. This is where UTF-8 comes in.

### UTF-8

UTF-8 encodes Unicode code points as items of variable length. All ASCII characters can be stores in 1 byte, and most runes require at most 2 or 3 bytes.

All Go source files are UTF-8, and Go’s `unicode` package provides functions for all sorts of operations on Unicode text (changing case, detecting whitespace, and so on).

UTF-8 is convenient in many ways (see pages 68-69 for examples), but it presents one potential problem for programmers. You have to work harder to get at individual UTF-8 characters. For example, if you want to process a UTF-8 string character by character, you might assume you need something like `utf8.DecodeRuneInString`.

```go
s := "Hello, 世界"
for i := 0; i < len(s); {
    r, size := utf8.DecodeRuneInString(s[i:])
    fmt.Printf("%d\t%c\n", i, r)
    i += size
}
```

But Go has you covered! When you want to loop over a string, “Go’s `range` loop…performs UTF-8 decoding implicitly” (69).

```go
for i, r := range "Hello, 世界" {
    fmt.Printf("%d\t%q\t%d\n", i, r, r)
}
```

You can convert between a UTF-8-encoded string and a slice of runes with `[]rune`, and you can turn a slice of runes back into a string with `string`. If you convert an integer value using `string`, you get the character: `string(65)` yields 'A' not "65". To get "65" you need `strconv.Itoa`.

### Strings and Byte Slices

Go provides four packages that come in handy when manipulating strings: `bytes`, `strings`, `strconv`, and `unicode`. `strings` provides functions for searching, replacing, comparing, trimming, splitting, and joining strings. `bytes` has similar functions that operate on byte slices (type `[]byte`). Byte slices come in handy for programs that need to handle string data and don’t want too many allocations and copies from immutable strings. `strconv` provides conversions between strings and other values and for quoting or unquoting strings. The `unicode` package has functions to test and convert runes—i.e., single Unicode characters.

### Conversions between Strings and Numbers

If you want to convert an integer to its string equivalent, you can use `fmt.Sprintf` with the `%d` verb or `strconv.Itoa` (= "integer to ASCII"). If you want to format integers into a specific base, you can use `strconv.FormatInt` or `strconv.FormatUint` or you can use `fmt` functions with the `%b`, `%d`, `%o`, or `%x` verbs to control the representation of the integer.

If you need to parse a string representing an integer, you should probably use `strconv.Atoi` or `strconv.ParseInt`. You can also use `fmt.Scanf`, but the authors warn that “it can be inflexible, especially when handling incomplete or irregular input” (75).

## Constants

> Constants are expressions whose value is known to the compiler and whose evaluation is guaranteed to occur at compile time, not at run time. The underlying type of every constant is a basic type: boolean, string, or number (75).

You can include a type in a constant declaration (e.g., `const name type = value`), but the type is often unnecessary. The compiler can often infer the type from the expression on the right-hand side.

Constants can appear in types, in particular in the length of an array type. This can make your code clearer: you get a named item rather than a magic number. For example:

```go
const IPv4Len = 4

func parseIPv4(s string) IP {
    var p [IPv4Len]byte
    // ...
}
```

### The Constant Generator `iota`

You can create a sequence of related values. The value of `iota` begins with zero and increments by one for each item in a sequence. Here’s a simple and then a less simple example.

```go
// Simple
type Weekday int

const (
    Sunday Weekday = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)

// Less simple
type KeySet byte

const (
    Copper Keyset = 1 << iota   // 1
    Jade                        // 2
    Crystal                     // 4
)
```

In the first example, each weekday is an integer, running from 0 to 6. In the second example, we can use bitmasks to store [different magical keys that a player in a game has in a single byte](https://www.ardanlabs.com/blog/2021/04/using-bitmasks-in-go.html).

### Untyped Constants

You can create a constant in any of the basic data types, but there are also six kinds of “uncommitted constants” (78): untyped boolean, untyped integer, untyped rune, untyped floating-point, untyped complex, and untyped string. The numeric untyped constants are useful because they store values more precisely than basic types and arithmetic involving them is more precise: “you can assume at least 256 bits of precision” (78). I am not sure what the point is for the other three types, however.
