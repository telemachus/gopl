# Chapter 7: Interfaces

Interfaces are types that gather together what other types can do. In other words, if a type satisfies the `Stringer` or `ReadCloser` interface, then we know various things about what objects of that type can and will do.

The most important thing about Go interfaces is “that they are *satisfied implicitly*” (171, their emphasis). Many programming languages have explicit interfaces (e.g., Java), but Go makes things easier. If a type has all the methods in an interface, then that type satisfies the interface. As a result, you have a great deal of freedom. You can create a new interface that collects together existing types from other people’s packages, and you can satisfy existing interfaces, all with a reasonable amount of code. Both of these are useful when using other libraries.

## Interfaces as Contracts

There are two kinds of types in Go: concrete and abstract types. Concrete types, like slices or numbers, specify “the exact representation of its values and…the intrinsic operations of that representation…A concrete type may also provide additional behaviors through its methods. When you have a value of a concrete type, you know exactly what it *is* and what you can *do* with it” (171). An abstract type, or interface type, doesn’t reveal “the representation or internal structure of its values” (171). And it doesn’t necessarily reveal the entire (or even basic) operations of its values. An abstract type specifies only some methods of its values. If you know something satisfies an interface, you don’t necessarily know anything about the nature of that object, and you don’t necessarily know it’s basic operations. But you know at least some of what it can do, how it will behave.

Here’s an example. The function `Fprintf` has `io.Writer` as its first parameter. An `io.Writer` is an interface type, and it’s declaration serves as a contract for any method that will satisfy the interface.

```go
// Writer is the interface that wraps the basic Write method.
//
// Write writes len(p) bytes from p to the underlying data stream.
// It returns the number of bytes written from p (0 <= n <= len(p))
// and any error encountered that caused the write to stop early.
// Write must return a non-nil error if it returns n < len(p).
// Write must not modify the slice data, even temporarily.
//
// Implementations must not retain p.
type Writer interface {
	Write(p []byte) (n int, err error)
}
```

Any type that provides a `Write` method that meets this contract satisfies the `Writer` interface. Thus, any `Writer` can be the first argument in a call to `Fprintf`. That method, in turn, cannot assume that it is writing to a file. All it can assume is that it has hold of something that it can call `Write` on. This gives us enormous freedom to use such interfaces for our own purposes and in our own ways.

Here’s an example from the authors.

```go
type ByteCounter int
func (c *ByteCounter) Write(p []byte) (int, error) {
    *c += ByteCounter(len(p))
    return len(p), nil
}

var c ByteCounter
fmt.Println(c) // "0": the zero value for an int is zero, so this is safe.
c.Write([]byte("hello"))
fmt.Println(c) // "5" = len("hello")

c = 0          // Since the underlying type is an int, we can reset bc.
var name = "Dolly"
fmt.Fprintf(&c, "hello, %s", name)
fmt.Println(c) // "12" = len("hello, Dolly")
```

## Interface Types

An interface type specifies the methods that a concrete type must have to belong to that interface. Many interfaces only specify a single method. For example, to be a `Reader`, a type must implement a `Read` method, and to be a `Closer`, a type must implement a `Close` method.

You can also create an interface by embedding other interfaces. For example:

```go
package io

type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

type Closer interface {
    Close() error
}

type ReadWriter interface {
    Reader
    Writer
}

type ReadWriteCloser interface {
    Reader
    Writer
    Closer
}
```

## Interface Satisfaction

A type that has all the methods in an interface satisfies the interface. “As a shorthand, Go programmers often say that a concrete type ‘is a’ particular interface type, meaning that it satisfies the interface” (175).

An expression may be assigned to an interface if its type satisfies the interface. The authors give this example.

```go
var w io.Writer
w = os.Stdout           // OK: *os.File has Write method
w = new(bytes.Buffer)   // OK: *bytes.Buffer has Write method
w = time.Second         // compile error: time.Duration lacks Write method

var rwc io.ReadWriteCloser
rwc = os.Stdout         // OK: *os.File has Read, Write, Close methods
rwc = new(bytes.Buffer) // compile error: *bytes.Buffer lacks Close method
```

The same rule applies when the right-hand side is itself an interface

```go
// Code as above, minus the compile errors
w = rwc                 // OK: io.ReadWriteCloser has Write method
rwc = c                 // compile error: io.Writer lacks Close method
```

Once you assign a concrete type to a variable that’s been set to the type of a specific interface, the variable may seem to lose methods.

```go
os.Stdout.Write([]byte("hello"))    // OK: *os.File has Write method
os.Stout.Close()                    // OK: *os.File has Close method
var w io.Writer
w = os.Stout
w.Write([]byte("hello"))            // OK: io.Writer has Write method
w.Close()                           // compile error: io.Writer lacks Close
```

The type `interface{}`, aka *the empty interface*, tells us nothing about the concrete types that satisfy it, and it places no demands on those concrete types. You can assign any value to the empty interface. This may seem useless, but is allows generic functions using type assertions. (We’ll see more about type assertions later in the chapter.)

It’s common to see interfaces satisfied by pointers to structs, but other types can satisfy interfaces. The authors mention slice types (e.g., `geometry.Path`), map types (e.g., `url.Values`), and function types (e.g., `http.HandlerFunc`) as reference types that sometimes satisfy interfaces. In addition, basic types sometimes satisfy interfaces (e.g., `time.Duration`).

One concrete type may satisfy more than one interfaces, and the interfaces may have nothing to do with one another.

## Parsing Flags with `flag.Value`

Go’s `flag` package provides several built-in types that arguments can take. If you use a built-in type, `flag.Parse()` will automatically check the validity of arguments given. For example, the package provides `Duration` flags that accepts all the specifiers [recognized by `time.Duration`](https://pkg.go.dev/time@latest#ParseDuration). If you use this flag, then you get error checking for free. If a user ask sfor `-period 1 day`, without any code on your part, the user will see `invalid value "1 day" for flag period: time: invalid duration 1 day`. (We can argue about whether this is good UX, but that’s a separate matter.)

You can also easily use the `flag.Value` interface to create new type-specific arguments. Here’s what the interface looks like. I’ll also include code from the book that satisfies the interface.

```
package flag

// Value is the interface to the value stored in a flag.
type Value interface {
    String() string
    Set(string) error
}

type celsiusFlag struct { Celsius }

func (f *celsiusFlag) Set(s string) error {
    var unit string
    var value float64
    fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed; see below
    switch unit {
    case "C", "°C"":
        f.Celsius = Celsius(value)
        return nil
    case "F", "°F":
        f.Celsius = FToC(Fahrenheit(value))
        return nil
    }
    // Normally, you should check the return value of fmt.Sscanf for an error,
    // but in this case an invalid value will fall through and reach this
    // point. So handle errors at the end of the function.
    return fmt.Errorf("invalid temperature %q" s)
}

func (c Celsius) String() string {
    return fmt.Sprintf("%g°C", c)
}
```

## Interface Values

> Conceptually, a value of an interface type, or *interface value*, has two components, a concrete type and a value of that type. These are called the interface’s *dynamic type* and *dynamic value* (181).

Consider the following run of code.

```go
var w io.Writer
w = os.Stdout
w = new(bytes.Buffer)
w = nil
```

The value and type of `w` change four times in four lines, though the fourth is the same as the first. After line 1, the type of `w` is `nil`, and its value is also `nil`. (The `nil` type has only one value: `nil`.) If you try to write to `w` at this point, you get a panic for a “`nil` pointer dereference." After the second line, the type of `w` is `*os.File`, and its value is `fd int = 1(stdout)`. You can now use `w` to write output to standard output. After the third line, `w` has the type `*bytes.Buffer`, and the value is `data []byte`, a newly allocated and empty byte slice. You can now write to the byte slice. Finally, after the fourth line, `w` once again has `nil` for its type and value.

You can compare interface values, but you may not like the result. Two interface values are equal if (1) they have the same dynamic type and (2) they have equal dynamic values, according to the rules for comparing such values. However, if the dynamic values are not comparable (e.g., if they are slices), then Go will panic when you make the comparison. I’m not sure when I would want to compare interface values, but maybe don’t?

If you need to check the dynamic type of an interface (for debugging or while dealing with an error), you can use `%T` with `fmt`. `fmt` uses reflection to determine the type of an interface, and the book will say more about reflection later.

### Caveat: An Interface Containing a `Nil` Pointer is Non-`Nil`

The following two declarations yield subtly different results.

```go
var buf io.Writer
var buf *bytes.Buffer
```

In the first case, the dynamic type and dynamic value of `buf` are `nil`. In the second case, the dynamic value of `buf` is `nil`, but the dynamic type is `*bytes.Buffer`. That can lead to a runtime panic if you’re not careful. Consider the following code that might follow one of the declarations above.

```go
if debug {
    buf = new(bytes.Buffer)
}
f(buf)

func f(out io.Writer) {
    // ...do whatever
    if out != nil {
        out.Write([]byte("did whatever\n"))
    }
}
```

If you declare `buf` to be of type `*bytes.Buffer`, this code will panic at runtime. You’re safe, however, if you declare `buf` to be of type `io.Writer`.

As the authors explain, “altough a nil `*bytes.Buffer` has the methods needed to satisfy the interface, it doesn’t satisfy the *behavioral* requirements of the interface. In particular, the call violates the implicit precondition of `(*bytes.Buffer).Write` that is receiver is not nil, so assigning the nil pointer to the interface was a mistake” (185-186).

## Sorting with `sort.Interface`

> An in-place sort algorithm needs three things—the length of the sequence, a means of comparing to elements, and a way to swap two elements—so they are the three methods of `sort.Interface` (186).

```go
type Interface interface {
    Len() int
    Less(i, j int) bool // i, j are indices of sequence elements
    Swap(i, j int)
}
```

In order to make a sequence sortable, you need to define a type for that sequence and implement the three methods. That is, you need to satisfy the interface. Here is a simple example.

```go
type StringSlice []string

func (ss StringSlice) Len() int {
    return len(ss)
}

func (ss StringSlice) Less(i, j int) bool {
    return ss[i] < ss[j]
}

func (ss StringSlice) Swap(i, j int) {
    return ss[i], ss[j] = ss[j], ss[i]
}

sort.Sort(StringSlice(whatever))
```

You don’t have to write your own sort for strings because Go’s runtime provides `sort.Strings`. However, this gives you a simple example of how the sort interface works.

Here is a more complicated example.

```go
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type byYear []*Track

func (x byYear) Len() int           { return len(x) }
func (x byYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x byYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
```

Go’s `sort` package provides a `Reverse` function that elegantly wraps any defined sort. Here it is.

```go
package sort

// Because reverse is lowercase, it is not exported.
type reverse struct{ Interface } // that is, sort.Interface

func (r reverse) Less(i, j int) bool { return r.Interface.Less(j, i) }

func Reverse(data Interface) Interface { return reverse{data} }
```

The new interface wraps the old one. The old interface supplies the `Len` and `Swap` methods unchanged. The new interface returns the old interface’s `Less` method with `i` and `j` swapped in the call. That’s all it takes to reverse any sort.

## The `http.Handler` Interface

The `http.Handler` interface serves as the basis for HTTP servers in Go. That interface looks like the following.

```go
package http

type Handler interface {
    ServeHTTP(w ResponseWriter, r *Request)
}

func ListenAndServe(address string, h Handler) Error
```

The `ListenAndServe` function runs forever. When the server fails (or fails to start) with an error, the function returns that error, which is always non-nil. The authors use the following as a toy example.

```go
func main() {
    db := database{"shoes": 50, "socks": 5}
    log.Fatal(http.ListenAndServe("localhost:8000", db))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    switch req.URL.Path {
    case "/list":
        for item, price := range db {
            fmt.Fprintf(w, "%s: %s\n", item, price)
        }
    case "/price":
        item := req.URL.Query().Get("item")
        price, ok := db[item]
        if !ok {
            w.WriteHeader(http.StatusNotFound) // 404
            fmt.Fprintf(w, "no such item: %q\n", item)
            return
        }
        fmt.Fprintf(w, "%s\n", price)
    default:
        w.WriteHeader(http.StatusNotFound) // 404
        fmt.Fprintf(w, "no such page: %s\n", req.URL)
    }
}
```

In this case, the database type is a map, and it implements the method necessary to serve as an instance of the `Handler` interface (i.e., `ServeHTTP(w http.ResponseWriter, req *http.Request)`. This (toy) server handles only two URLs: `/list` and `/price?item=<query>`. If you visit any other path, the browser receives a 404 status and an error message.

As an alternative to `WriteHeader`, you can use `http.Error`. Compare the following. (I can’t see an advantage either way.)

```go
// The first way
w.WriteHeader(http.StatusNotFound) // 404
fmt.Fprintf(w, "no such page: %s\n", req.URL)

// The second way
msg := fmt.Sprintf("no such page: %s\n", req.URL)
http.Error(w, msg, http.StatusNotFound) // 404
```

You can define a server for each route, but that will get tedious. Instead, you can use `ServeMux`, “a *request multiplexer*, to simplify the association between URLs and handlers. A `ServeMux` aggregates a collection of `http.Handler`s into a single `http.Handler`” (193).

In the example they give for `ServeMux`, they also demonstrate how to use `http.HandlerFunc` (a type) to convert a method value into something that will satisfy the `http.Handler` interface. Consider the code, and then I’ll try to explain.

```go
func main() {
    db := database{"shoes": 50, "socks": 5}
    mux := http.NewServeMux()
    mux.Handle("/list", http.HandlerFunc(db.list))
    mux.Handle("/price", http.HandlerFunc(db.price))
    log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
    for item, price := range db {
        fmt.Fprintf(w, "%s: %s\n", item, price)
    }
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
    item := req.URL.Query().Get("item")
    price, ok := db[item]
    if !ok {
        w.WriteHeader(http.StatusNotFound) // 404
        fmt.Fprintf(w, "no such item: %q\n", item)
        return
    }
    fmt.Fprintf(w, "%s\n", price)
}
```

The arguments for `mux.Handle` need to be a route and an `http.Handler`. But `db.list` and `db.price` are *method values*. A method value (see Chapter 6.4) does not itself have methods, and therefore “it doesn’t satisfy the `http.Handler` interface and can’t be passed directly to `mux.Handle` (194). Thus, the authors use `http.HandlerFunc` to convert the method value into something that satisfies the underlying `http.Handle` interface. An `http.HandlerFunc` has a generic `ServeHTTP` method that looks like the following:

```go
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
    f(w, r)
}
```

As you can see, that generic method simply calls `f` (the function receiver) on the appropriate arguments. There are ways to simplify this even further, but I think that’s enough for now.

## The `error` Interface

The interface type is surprisingly straightforward.

```go
type error interface {
    Error() string
}
```

As the authors say, “[t]he entire `errors` package is only four lines long” (196).

```go
package errors

func New(text string) error { return &errorString{text} }

type errorString struct { text string }

func (e *errorString) Error() string { return e.text }
```

You can create new errors using `errors.New` and passing in the text of your error message as a string. Why use a pointer rather than a value for the `Error` function? The design guarantees “that every call to `New` allocates a distinc error instance that is equal to no other. We would not want a distinguished error such as `io.EOF` to compare equal to one that merely happened to have the same message” (196).

Instead of using `errors.New`, you can use `fmt.Errorf` instead. It has this signature.

```go
func Errorf(format string, args ...interface{}) error {
    return errors.New(Sprintf(format, args...))
}
```

## Type Assertions

A type assertion, which looks like `x.(T)`, is an operation on an interface value. `x` is some expression of an interface type, and `T` is a type, called the *asserted type*. A type assertion checks whether the dynamic type of `x` matches `T`.

Type assertions come in two flavors, depending on whether the asserted type is a concrete type or an interface type. These two flavors work differently.

1. If the asserted type is a concrete type, the type assertion checks whether the dynamic type of `x` is identical to `T`. If it is, the result of the type assertion is the dynamic value of `x`, which is `T`. If the assertion fails, then Go panics. Here’s an example.

        ```go
        var w io.Writer
        w = os.Stdout
        f := w.(*os.File)       // Joy: f = &os.File
        c := w.(*bytes.Buffer)  // Sadness: this throws a panic
        ```
2. If the asserted type is an interface type, the type assertion checks whether the dynamic type of `x` satisfies that interface. If it does, then the result is the interface type T. That means that you can use this second kind of type assertion to change get a different and potentially larger set of methods.

        ```go
        var w io.Writer
        w = os.Stdout
        rw := w.(io.Reader)     // Joy: rw now has both Read and Write methods

        w = new(ByteCounter)
        rw = w.(io.ReadWriter)  // Sadness: this panics: no Read method found
        ```

If the operand is a nil value, then the type assertion will always fail.

However, if the type assertion assigns two results, then instead of a panic, the statement will safely return `false` as the second value. This makes for a safe way to check whether a variable supports a given interface.

        ```go
        var w io.Writer = os.Stdout
        f, ok := w.(*os.File)       // Joy: ok, f == &os.File
        b, ok := w.(*bytes.Buffer)  // Sadness: !ok, b == nil
        ```

An idiom is Go: use the double-return-value form with an `if` clause, and reuse the same variable name.

        ```go
        if w, ok := w.(*os.File); ok {
            // …use w…
        }
        ```

## Discriminating Errors with Type Assertions

The `os` package provides helper methods (`IsExist`, `IsNotExist`, and `IsPermission`) to help distinguish three different types of file errors. If you want to create a file that already exists, you have the first problem. If you want to read a file that does not yet exist, you have the second problem. If you don’t have proper permission to work with files, you have the third problem. Under the hood, Go uses type assertions to sniff out which error has been thrown. The code looks like the following.

        ```go
        fun IsNotExist(err error) bool {
            if pe, ok := err.(*PathError); ok {
                err = pe.Err
            }
            return err == syscall.ENOENT || err == ErrNotExist
        }
        ```

## Querying Behaviors with Interface Type Assertions

Consider the following code.

        ```go
        func writeHeader(w io.Writer, contentType string) error {
            if _, err := w.Write([]byte("Content-Type: ")); err != nil {
                return err
            }
            if _, err := w.Write([]byte(contentType)); err != nil {
                return err
            }
            // …
        }
        ```

The conversions from strings to bytes “allocates memory and makes a copy, but the copy is thrown away immediately after” (208). Probably this isn’t a problem, but imagine that it leads to a performance hit. How can we try to avoid the allocation and copying? If you guessed *with an interface type assertion*, you’ve been paying attention.

        ```go
        func writeString(w io.Writer, s string) (n int, err error) {
            type stringWriter interface {
                WriteString(string) (n int, err error)
            }
            if sw, ok := w.(stringWriter); ok {
                return sw.WriteString(s)    // avoid a copy
            }
            return w.Write([]byte(s))       // allocate the copy if we must
        }

        func writeHeader(w io.Writer, contentType string) error {
            if _, err := writeString(w, "Content-Type: "); err != nil {
                return err
            }
            if _, err := writeString(w, contentType); err != nil {
                return err
            }
            // …
        }
        ```

Go provides `io.WriteString` for exactly this purpose. As the authors say, “[i]t is the recommended way to write a string to an `io.Writer`” (209).

## Type Switches

You can vary behavior by using type switches. A type switch, like a regular switch, simplifies a sequence of `if…else` statements. For example:

        ```go
        switch x.(type) {
        case nil: // …
        case int, uint: // …
        case bool: // …
        case string: // …
        default: // …
        }
        ```

Since you often need the specific value from the type switch, you can capture it in a variable for use within the branches.

        ```go
        switch x := x.(type) {
        case nil: // …
        case int, uint: // …
        case bool: // …
        case string: // …
        default: // …
        }
        ```
