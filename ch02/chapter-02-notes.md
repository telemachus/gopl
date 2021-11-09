# Chapter 2: Program Structure

## Names

Names in Go (for functions, variables, constants, types, statement labels, and packages) (1) must begin with a letter or underscore and (2) may have any number of additional letters, digits, and underscores. In Go, names are case sensitive: `heapSort` and `Heapsort` are different names.

Go has 25 keywords that cannot be used as names:

+ `break` `case` `chan` `const` `continue` `default` `defer` `else`
+ `fallthrough` `for` `func` `go` `goto` `if` `import` `interface` `map`
+ `package` `range` `return` `select` `struct` `switch` `type` `var`

Go also has around three dozen predeclared names. These are not reserved: you can use them in declaration, but you should be careful not to let them become confusing.

+ Constants: `true` `false` `iota` `nil`
+ Types: `int` `int8` `int16` `int32` `int64`
+ Types: `uint` `uint8` `uint16` `uint32` `uint64` `uintptr`
+ Types: `float32` `float64` `complex128` `complex64`
+ Types: `bool` `byte` `rune` `string` `error`
+ Functions: `make` `len` `cap` `new` `append` `copy` `close` `delete`
+ Functions: `complex` `real` `imag` `panic` `recover`

What is the scope of a name? If you declare a name inside of a function, then it is local to the function. If you declare a name outside of a function, then it is visible in all files of the package that it belongs to. If a name begins with a capital letter, then it is exported and will be visible when people import its package. Otherwise, the name is not exported, and it is private to its own package. Package names themselves always begin with a lower-case letter.

Although Go does not limit the length of names, Go idiom prefers short names. Especially when the scope of a variable is small, you should prefer short names. And if a variable is traditional (e.g., `i` for an index in a loop or `r` for a reader), you should use a single-letter variable.

Go programmers prefer camel case for names made up of more than one word. Don’t use underscores. E.g., `quoteRune` or `QuoteRune` not `quote_rune`. Acronyms and initialisms should be all in the same case. E.g. `escapeHTML` or `htmlEscape` not `HtmlEscape` or `escapeHtml`.

## Declarations

“A *delcaration* names a program entity and specifies some or all of its properties” (28). The four most common types of declarations are for variables, constants, types, and functions.

A Go program consists of one ore more files whose names end in `.go`. Every file begins with a `package` declaration that says what package the file is part of. After the `package` declaration, you should put any `import` declarations and then package-level declarations of types, variables, constants, and functions in any order.

Package-level declarations appear outside of any function in a package, and the name of a package-level entity is visible in any file that belongs to the same package as where the package-level entity is declared. Local declarations are visible only inside the function in which they are declared—and maybe only within a part of that function.

### Variables

“A `var` declaration creates a variable of a particular type, attaches a name to it, and sets its initial value” (30). The basic form of a variable declaration is `var name type = expression`. You can omit the type or the expression but not both. If the type is left out, the type is inferred from the expression. If the expression is left out, the initial value will be the type’s zero value. For example, the zero value for numbers is `0`, the zero value for strings is the empty string, the zero value for booleans is `false`, and the zero value for reference types (maps, slices, pointers, channels, and functions) is `nil`. A composite type like an array or struct has the zero value for each of its members.

You can declare and initialize multiple variables in a single declaration. If you omit the type, you can even declare multiple variables of different types. For example, `var b, f, s = true, 2.3, "four"`. You can use literal values or arbitrary expressions to initialize variables. You can initialize multiple variables with a single function that returns multiple values. (For example, functions often return one or more variables followed by a last item that may be an error or nil.) Package-level variables are initialized before `main` begins, and local variables are initialized when their declarations appear during function execution.

### Short Variable Declarations

Within functions, you can use a short variable declaration. These look like `name := expression`. The return type of *expression* determines the type of *name*. Most local variables use short variable declarations. Go programmers reserve local `var` declarations for two cases: (1) when the type needs to differ from the initializer expression or (2) when the variable does not need an initial value.

A short variable declaration must have *some* new variables on the left-hand side, but not all of the variables on the left-hand side need to be new. For example:

```go
in, err := os.Open(infile)
// ...
out, err := os.Create(outfile)
```

In the first statement, both `in` and `err` are declared and initialized. In the second statement, `out` is declared, but `err` already exists and has type `error`. Thus, in the second statement one variable (`out`) is declared, but the other (`err`) is assigned a new value. If you have only existing values on the left-hand side, then you need an assignment (`=`) instead of a short variable declaration (`:=`).

### Pointers

They define *variable* as “a piece of storage containing a value” (32), and they remind us that some variables have names (e.g., `x := 0`) while others are only identified by expressions (e.g., `x[i]`). A pointer value, however, is the address of a variable. The pointer refers to “the location at which a value is stored” (32). You can use pointers to indirectly read or update the value of a variable.

To work with pointers, you use two operators: `&` and `*`. The `&` yields a pointer to a variable; you can read `&x` as “address of x.” The `*` yields the value of the variable at a specific address. Thus, in `p := &x; fmt.Println(*p)`, `*p` returns the value at the address of `x`. Since `*p` is a variable, you can also assign to it. Thus `x := 2; p := &x; *p = 1` updates the value of `x`.

The zero value of all pointers is `nil`. You can compare pointers: two pointers are equal if and only if they point to the same variable or they are both `nil`.

You can pass pointers to functions and return pointers to functions. You will want to pass a pointer to a function in cases where you want to change a variable. In Go, all variables are passed by value not by reference. For most types, this means that variables are copied. And that, in turn, means that if you change the value in a function, the calling context will not see the change.

### The `new` Function

You can also create a variable with the built-in `new` function. When you call `new(T)`, the following happens.

+ Go creates an unnamed variable of type T.
+ Go initializes the unnamed variable to the zero value of T.
+ Go returns the address of the unnamed variable of type T, namely a value of type `*T`.

All of this said, you rarely need to use `new`. The most common unnamed variables are structs, and you can use struct literals instead of `new` to create a new struct variable.

Since `new` is a predeclared function and not a keyword, you can redefine the name *new*. Don’t do this, however.

### Lifetime of Variables

The lifetime of a variable is the time in the execution of the program that the variable exists. Briefly, a package-level variable has a lifetime of the entire program execution, but local variables have dynamic lifetimes. They are created when their declaration is reached, and their lifetime extends until they can no longer be reached. When they become unreachable, variables can be garbage collected.

## Assignments

You can update the value of a variable by an assignment statement. An assignment statement is a variable (or indirect variable) on the left, an `=` sign, and an expression on the right.

```go
x = 1
*p = 2
person.name = "bob"
count[x] = count[x] * scale
```

All of the arithmetic and bitwise binary operators has a corresponding assignment operator. For example, `+` has `+=`. We can rewrite our last example above as `count[x] *= scale`. You can also increment and decrement numeric variables with `++` and `--`.

### Tuple Assignment

You can have several variables on either side of `=`. All of the right-hand side expressions are evaluated before any variable is updated. You can use tuple assignment to swap values. E.g., `x, y = y, x`. Tuple assignment is also common when a function returns multiple values. In particular, this happens where a function may return an error. E.g., `f, err = os.Open("foo.txt")`. `err` will be `nil` if there is no error; otherwise, it has a specific error that occurred when `os.Open` tried to open the given file. In addition to functions, map lookups, type assertions, and channel receives return a second value that can be useful. (See later chapters for details.)

```go
v, ok = m[key]
v, ok = x.(T)
v, ok = <-ch
```

You can use tuple assignment for several simple assignments if you want brevity. E.g., `i, j, k = 2, 3, 5`. But the authors recommend that you don’t do this often.

### Assignability

In addition to explicit assignment, assignment often happens implicitly. Here’s one example:

```go
medals := []string{"gold", "silver", "bronze"}
// Here’s an equivalent
medals := make([]string, 3)
medals[0] = "gold"
medals[1] = "silver"
medals[2] = "bronze"
```

Maps and channels also involve implicit assignment.

Go enforces various rules about assignability. In particular, type rules are enforced. Assignability also affects whether two values can be compared with `==` or `!=`. The authors will say more about assignability and comparability as new types are introduced.

## Type Declarations

You can use variables of the same type (e.g., `int`) to represent different concepts (e.g., a loop index, a month, a day of the week, money, a grade). In order to make your intentions more clear—and to enforce type safety—you can declare a new type. When you declare a type, you create a new named type using some existing underlying type. “The named type provides a way to separate different and perhaps incompatible uses of the underlying type so that they can’t be mixed unintentionally.”

Type declarations look like this: `type name underlying-type`. For example:

```go
type Celsius float64
type Fahrenheit float64
//...

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c * 9 / 5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9 }
```

Although Celsius and Fahrenheit values share an underlying type (`float64`), the type declarations make the program clearer. The type declarations also protect users of this library from accidentally performing mathematical operations on a mix of Celsius and Fahrenheit values. You must explicitly convert values of one type into the other using a type conversion: `Type(value)`.

You can convert between declared types if both values share an underlying type or if both are pointers to the same underlying type. The authors add “conversions change the type but not the representation of the value,” but I’m not sure what this means (40).

You can also convert between numeric types and between strings and some slice types. However, these conversions may change the representation of the value. If you convert a floating value to an integer, for example, you will lose whatever fractional part the floating value held.

The underlying type of a declared type controls its structure, representation, and what operations the type supports. Thus, you can use any arithmetic operators on Celsius values exactly as you would on float values. You can use comparison operators on values of the same declared types, but not on values of different declared types—even if they share an underlying type.

Once you have named a new type, you can define methods on the type and satisfy interfaces with the new type. Both of these are a big deal in Go. For example, with the following, Celsius values now satisfy the Stringer interface.

```go
func (c Celsius) String() string {
    return fmt.Sprintf("%g°C", c)
}
```

## Packages and Files

Go organizes libraries and applications into packages. A package is one or more `.go` files in a directory whose name provides the import path. For example, the files in `unicode/utf16` contain the code for the `utf16` package. You import a package using its full name (e.g., `unicode/utf16`), but by default you refer to code in a package with only the last part of its name (e.g., `utf16.Whatever`).

You can split a single package into as many files as you like. Each file must declare its package at the top of the file. Everything at the package level is visible to every other file in the package. However, only identifiers that begin with an upper-case letter are exported. All other identifiers are private to the package itself.

### Imports

Every package in a Go program is identified by a unique import path (e.g., `gopl.io/ch2/tempconv`. The name of a package is the short name that appears in its package declaration (e.g., `tempconv`). Package names do not have to be unique, and in the case of a conflict, you can adjust the normal rules for the name of a package. But by default, it will be the last part of the packages import path. Thus `gopl.io/ch2/tempconv` yields a package named `tempconv` by default.

You cannot use a package without importing it. In addition, if you import a package and then don’t use it, the `go` tool won’t run or compile your program.

### Package Initialization

Go enforces a sequence when initializing variables. Briefly, package-level variables are initialized in the order that they appear, but with dependencies resolved first. If there are multiple files, the `go` tool sorts them by names before invoking the compiler. (You can also give files to the compiler in a manual order, I think?)

Sometimes you need a function to initialize a value. Go provides `init` functions for this purpose. (People seem to recommend not using `init` functions in more recent tutorials.)

### Scope

The scope of a declaration is the part of the source code where a variable refers to that declaration. Scope is not lifetime. Lifetime is a run-time property, and it describes an extent of time. Scope is a compile-time property, and it describes an extent of space.

Go uses block scope. A name declared within a block is not visible outside that block. The *universe block* refers to the entire source code. There is also a block for an entire package, a single file, a `for` or `if` statement, and so on.

You can declare multiple variables with the same name in different scopes. The inner declarations will shadow or hide the outer ones. Thus, the outer variables become inaccessible within smaller scopes. Here’s an extreme example:

```go
func main() {
    x := "hello!" // First declaration of x
    for i := 0; i < len(x); i++ {
        x := x[i] // Second declaration of x; first x is now inaccessible
        if x != '!' {
            x := x + 'A' - 'a' // Third declaration of x; first two shadowed
            fmt.Printf("%c", x) // Prints "HELLO", one letter per iteration
        }
    }
}
```

Here’s a more subtle problem:

```go
var cwd string
func init() {
    cwd, err := os.Getwd() // Not what you think: compile error unused cwd!
    if err != nil {
        log.Fatalf("os.Getwd failed: %v", err)
    }
}
```

Even worse, if you use `cwd` inside the `init` function, the program will compile, but it’s still wrong!

```go
var cwd string
func init() {
    cwd, err := os.Getwd() // Not what you think: compile error unused cwd!
    if err != nil {
        log.Fatalf("os.Getwd failed: %v", err)
    }
    log.Printf("Working directory = %s", cwd)
}
```

Inside the `init` function, `cwd` is a different variable than at the package level. Thus, `cwd` is still an empty string for the execution of the program. How would we fix this problem? We can declare `err` in advance and use `=` rather than `:=`.

```go
var cwd string
func init() {
    var err error
    cwd, err = os.Getwd()
    if err != nil {
        log.Fatalf("os.Getwd failed: %v", err)
    }
}

// Now cwd is ready to go...
```
