# Chapter 5: Functions

## Function Declarations

Function declarations have a name, list of parameters, an optional list of results, and a body.

```go
func name(parameter-list) (result-list) {
    // body
}
```

If a function returns nothing or one (unnamed) type, then you don't need the parentheses. You can name results exactly as you can name parameters. If you have one or more results, the function must specify a return except in unusual cases where “execution clearly cannot reach the end of the function, perhaps because the function ends with a call to `panic` or an infinite `for` loop with no `break`. (But, yeah, don’t do that.)

If you have multiple parameters or results of the same type, you can group them together before specifying the type. The following are equivalent.

```go
func f(i, j, k int, s, t string) { /* ... */ }
func f(i int, j int, k int, s string, t string) { /* ... */ }
```

The type, or signature, of a function is its formal layout of parameter types and results. The names of parameters and results don’t matter.

When you call a function in Go, you must provide an argument for each parameter in the order they appear in the signature. There are no default values, and you can’t swap order using keywords or names.

Within the body of a function, parameters are local variables. Arguments are passed by value. Non-reference types are copied, and reference types copy the address of the reference type. (That is, reference types are the same as passing a pointer to a non-reference type.)

Sometimes you will see a function declaration without a body. In such cases, the function is implemented in some language other than Go (e.g., assembly). The declaration in such cases defines the function’s type and signature for users.

## Recursion

Unsurprisingly, Go handles recursion well. The authors provide examples.

## Multiple Return Values

Functions in Go can return multiple results. You have to wrap multiple results in parentheses, but other than that, multiple results work no differently than single results.

Many examples of multiple-result functions involve a result to indicate failure or success. By convention, the (possible) failure value goes last. In some cases (the authors say when “the failure has only one possible cause”), the convention is to use a boolean as the failure result. For example with maps: `value, ok := map[key]`. In general, however, the failure value will be of `error` type. Also, in general, if the failure result is not nil, then you should expect all other values to be nil or useless.

## Anonymous Functions

You can use a function literal as a function in an expression. A function literal looks like a function declaration without a name for the function after the `func` keyword. For example:

```go
strings.Map(func(r rune) rune { return r + 1 }, "HAL-9000")
```

Anonymous functions are closures. They have access to the variables in their enclosing function and environment. For example:

```go
func squares() func() int {
    var x int
    return func() int {
        x++
        return x * x 
    }
}

func main() {
    f := squares()
    fmt.Println(f()) // 1
    fmt.Println(f()) // 4
    fmt.Println(f()) // 9
    fmt.Println(f()) // 16
}
```

Because they are closures, functions have state as well as behavior. The authors point out that this is why functions are not comparable. Different return values of `squares` will have the same function signature, but the value of their internal `x` variable may differ.

When an anonymous function uses recursion, you must declare a variable first and then assign an anonymous function to that variable. Otherwise, there would be no way to call the function recursively since it isn’t defined yet. For example:

```go
var visitAll func(items []string)
visitAll = func(items []string) {
    for _, item := range items {
        if !seen[item] {
            seen[item] = true
            visitAll(m[item])
            order = append(order, item)
        }
    }
}
```

### Caveat: Capturing Iteration Variables

You have to be careful when you (intentionally or otherwise) use the address of a loop variable. For example, what do you expect the following to print?

```go
func main() {
	original := []string{"foo", "bar", "fizz", "buzz", "fizzbuzz"}
	copies := make([]*string, 0, cap(original))

	for _, item := range original {
		copies = append(copies, &item)
	}

	for _, item := range copies {
		fmt.Println(*item)
	}
}
```

Answer: the program will print “fizzbuzz” five times. Why? In the assignment at the top of the loop (`for _, item := range original`), Go uses a single address to store successive values of `item`. This means that what you get later is the value of the last item stored at that address, namely “fizzbuzz.”

There are two ways to avoid this problem. First, you can create a new variable and take its address. Second, you can rewrite your code to take the address of the actual item from the original slice. Which you choose depends on what you want in a larger sense. If you care about the content of the variable, you probably want the first answer. If you care about the original address (regardless of what ends up stored there), you probably want the second answer. Here are examples of both solutions.

```go
// First solution:
for _, v := range values {
    copied := v
    output = append(output, &copied)
}

// Second solution:
for i := range values {
    output = append(output, &values[i])
}
```

Note also that this problem is not specific to `range` loops with `for`. You can create the same problem with a manual `for` loop. Here’s an example:

```go
var rmdirs []func()
dirs := tempDirs()
for i := 0; i < len(dirs); i++ {
    os.MkdirAll(dirs[i], 0755) // Okay
    rmdirs = append(rmdirs, func() {
        os.RemoveAll(dirs[i])  // Not okay
    })
}
```

The first use of `dirs[i]` is fine because it immediately uses the value stored in the single address that holds successive values of `i`. The second use of `dirs[i]` is not fine because by the time the functions stored in `rmdirs` are called, `i` equals its last value (`len(dirs) - 1`) for every call to `os.RemoveAll`. Long story short, the manual loop with `i` has also captured the last value of `i` for all members of `rmdirs`.

For more on this, see the following links:

+ [Pointer/Value Subtleties](http://jmoiron.net/blog/pointer-value-subtleties)
+ [Go gotcha #0: Why taking the address of an iterated variable is wrong](https://developmentality.wordpress.com/2014/02/25/go-gotcha-0-why-taking-the-address-of-an-iterated-variable-is-wrong)
+ [Go Gotcha: Don’t take the address of loop variables](https://www.evanjones.ca/go-gotcha-loop-variables.html)
+ [Go Common Mistakes: Using reference to loop iterator variable](https://github.com/golang/go/wiki/CommonMistakes#using-reference-to-loop-iterator-variable)

If you want to know more about *why* this happens, see this article: [Go internals: capturing loop variables in closures](https://eli.thegreenplace.net/2019/go-internals-capturing-loop-variables-in-closures).

## Deferred Function Calls

You can schedule functions to run when the containing function finishes. A deferred function call will happen whether the enclosing function ends normally or panics. You can ask for any number of functions to be deferred: they will run in last-in first-run order. Note that the deferred function call will be *evaluated* where you put it but run at the end of the enclosing function. This may matter in some cases. Here’s what a `defer` statement looks like.

```go
resp, err := http.Get(url)
if err != nil {
    return err
}
defer resp.Body.Close()
```

You should note two things. First, the defer statement goes *after* error handling. If the error is not nil, then there’s no body to close. Second, the defer statement needs parentheses. (See above about evaluation: you need to be able to pass arguments to the deferred function.)

You have to be careful when using defer statements inside loops. For example, the following is potentially dangerous. (I was caught by this gotcha my first month using Go!)

```go
for _, filename := range filenames {
    f, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer f.Close() // You may run out of file descriptors!
```

The deferred `f.Close()` calls will stack up until the entire enclosing function ends. As a result, you may run out of file descriptors.

You can manually close these files or you can wrap the loop work into another function. For example:

```go
for _, filename := range filenames {
    if err := doFile(filename); err != nil {
        return err
    }
}

func doFile(filename string) error {
    f, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer f.Close()
    // Do whatever you need to do with f.
    // ...
    return nil
}
```

In passing, the authors note here that you shouldn’t ignore the possible error from `f.Close()` because some file systems only report write errors when a file is closed.

## Panic

A panic occurs when Go detects (certain?) runtime errors. (For example, a panic occurs when you try to access an array out of bounds or dereference a nil pointer.)

You can also call for a panic within your program if you think your program has reached an unworkable state. You should almost never call `panic` within your program. However, there is a certain class of cases where a panic makes sense. In those case, Go often names functions `MustWhatever` and calls for an automatic panic if the `Must` call fails. The authors give the example of regex compilation.

```go
package regexp

func Compile(expr string) (*Regexp, error) { /* ... */ }
func MustCompile(expr string) *Regexp {
    re, err := Compile(expr)
    if err != nil {
        panic(err)
    }
    return re
}
```

## Recover

Usually a panic will (and should) lead to a crash. However, sometimes it makes sense to recover from a panic. You can use `recover` for such cases. Mostly, you shouldn’t expect to do this, but the authors demonstrate a complex way that you can test for specific kinds of panic and then either return a simple error or continue panicking.

```go
func soleTitle(doc *html.Node) (title string, err error) {
    type bailout struct{}

    defer func() {
        switch p := recover(); p {
        case nil:
            // no panic; do nothing
        case bailout{}:
            // "expected" panic; make it an error instead
            err = fmt.Errorf("multiple title elements")
        default:
            // an unexpected panic; panic away!
            panic(p)
        }
    }()

    // Rest of the function...
}
```
