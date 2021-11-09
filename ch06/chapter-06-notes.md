# Chapter 6: Methods

Go supports object-oriented programming, but maybe not as I would have expected. Go doesn’t use classes and objects, as found in Ruby or Python. Instead, Go considers an object “a value or variable that has methods, and a method is a function associated with a particular type. An object-oriented program is one that uses methods to express the properties and operations of each data structure so that clients need not access the object’s representation directly” (155).

## Method Declarations

You declare a method by adding an extra parameter to an ordinary function definition. The extra parameter precedes the function’s name, and this parameter connects the function with the parameter’s type. Here’s an example:

```go
type Point struct { X, Y float64 }

func (p Point) Distance(q Point) float64 {
    return math.Hypot(q.X-p.X, q.Y-p.Y)
}
```

The parameter `p` is known as the method’s *receiver* because “early object-oriented languages...described calling a method as ‘sending a message to an object’” (156). By convention, Go does not use names like `this` or `self` for the receiver. Instead, Go programmers use ordinary (and usually short) variable names: “[a] common choice is the first letter of the type name, like `p` for `Point`” (156). The name of a method like this is `Point.Distance`. That is `Type.MethodName`.

When you call the method, the receiver argument appears before the method name, and the method name is added after a dot: `p.Distance(q)`. An expression like `p.Distance` is known as a *selector* “because it selects the appropriate `Distance` method for the receiver `p` of type `Point`” (156). We also call dot notation for fields on struct types *selectors*. Since struct fields and methods live in the same name space, you cannot name a field and a method the same thing for the same struct. However, you can safely have a `Distance` method on two different types without a problem. Each type maintains its own name space. (Note that not all named types are structs. You can name other underlying types and then define methods on those names.)

## Methods with a Pointer Receiver

When you call a function in Go, the runtime makes a copy of each argument value. Therefore, if a function needs to change a value in the calling context, or if an argument is large enough that we want to avoid making a copy, we should pass the argument as a pointer instead of as a copied value. The same logic applies to methods and receivers. For example:

```go
func (p *Point) ScaleBy(factor float64) {
    p.X *= factor
    p.Y *= factor
}
```

If the receiver here was `p Point` rather than `p *Point`, the method would scale a copy—and be useless. Note that this method’s name is `(*Point).ScaleBy`. The parentheses matter: without them, the expression would be read as `*(Point.ScaleBy)`. That is, Go would see it as a pointer to a method  called `ScaleBy` with a value receiver `Pointer`.

By convention, if any method on a type needs a pointer receiver, then *all* methods for that type should take a pointer receiver, even those that don’t strictly need one.

Only named types `(Point)` and pointers to named types `(*Point)` can be the receiver in a method declaration. You cannot create a method on named types that are pointers themselves. The following won’t compile.

```go
type P *int
func (p P) f() { /* boom */ }
```

A method like `(*Point).ScaleBy` can be called in several ways.

```go
// One way
r := &Point{1, 3}
r.ScaleBy(2)

// A second way
p := Point{1, 3}
pptr := &p
pptr.ScaleBy(2)

// A third way
p := Point{1, 3}
(&p).ScaleBy(2)
```

The second and third way above require extra work or are ugly, so Go simplifies things. As the authors explain:

> If the receiver `p` is a *variable* of type `Point but the method requires a `*Point` receiver, we can use this shorthand: `p.ScaleBy(2)` and the compiler will perform an implicit `&p` on the variable.

However, note that this implicit addition of the "address-of" operator only works on variables “including struct fields like `p.X` and array or slice elements like `perim[0]`” (158). You cannot call a pointer method on a struct literal because a struct literal is “a non-addressable...reciever” (158). The following will not compile: `Point{1, 3}.ScaleBy(2)`.

On the other hand, Go will implicitly get the value from a pointer if you call a value method on a pointer receiver. For example:

```go
p := Point{1, 3}
pptr := &p
p.Distance(q) // Fine
(*pptr).Distance(q) // Also fine: explicitly get pptr’s value for Distance
pptr.Distance(q) // Also also fine: Go implicitly gets the value of pptr
```

Because Go does several kinds of implicit work for you, people become confused about what method calls are valid. The authors summarize the valid cases as follows.

1. The call is valid if the receiver argument has the same type as the receiver parameter. Both can be type `T` or type `*T`.

        ```go
        Point{1, 2}.Distance(q) // Parameter and argument are of type Point
        pptr.ScaleBy(2) // Parameter and argument are of type *Point
        ```

2. The receiver argument is a variable of type `T`, and the receiver parameter has type `*T`.

        ```go
        p.ScaleBy(2) // Implicitly, Go runs this as (&p).ScaleBy(2)
        ```

3. The receiver argument is a variable of type `*T`, and the receiver parameter has type `T`.

        ```go
        pptr.Distance(q) // Implicitly, Go runs this as (*pptr).Distance(q)
        ```

Finally, if a method has only value types, you can freely and safely copy instances of that type. However, if a method has a pointer receiver, you should not copy instances of that type or receiver. As examples, they say you can copy `time.Duration` instances all you like, but you should not copy `bytes.Buffer` instances.

### `Nil` is a Valid Receiver Value

Since pointers can be `nil`, it often makes sense to have methods handle `nil` gracefully and meaningfully. For example, consider this linked list of integers and its `Sum` method.

```go
type IntList struct {
    Value int
    Tail *IntList
}

func (il *IntList) Sum() int {
    if il == nil {
        return 0
    }
    return il.Value + il.Tail.Sum()
}

func main() {
	var il IntList
	fmt.Println(il.Sum()) // Or fmt.Println(IntList(nil).Sum())
}
```

Even before `il` is initialized, the call to `Sum` is safe, and it returns a reasonable result, namely zero.

## Composing Types by Struct Embedding

In the same way that struct embedding allows you to get at the fields of an embedded struct more easily, you can also use struct embedding to get at methods more easily. The methods of the embedded struct are said to be *promoted* to the embedding struct, and if both embedded and embedding types have methods of their own, we end up with *composition* of the total set of methods.

The anonymous field can be a named type or a pointer to a named type. That is, the following are both valid.

```go
type Point struct { X, Y float64 }
red := color.RGBA{255, 0, 0, 255}

// Named type as anonymous field
type ColoredPoint struct {
    Point
    Color color.RGBA
}
p := ColoredPoint{Point{1, 1}, red}

// Pointer to named type as anonymous field
type ColoredPoint struct {
    *Point
    Color color.RGBA
}
p := ColoredPoint{&Point{1, 1}, red}
```

The authors also demonstrate a way to use an unnamed struct type with methods declared on it.

```go
// A first way to implement a cache using package-level variables.
var (
    mu sync.Mutex
    mapping = make(map[string]string)
)

func Lookup(key string) string {
    mu.Lock()
    v := mapping[key]
    mu.Unlock()
    return v
}

// A second way to implement the cache using an unnamed struct
var cache = struct {
    sync.Mutex // An anonymous field: its methods will be promoted
    mapping map[string]string
}{
    mapping: make(map[string]string)
}

func Lookup(key string) string {
    cache.Lock()
    v := cache.mapping[key]
    cache.Unlock()
    return v
}
```

## Method Values and Expressions

In general, you select (`T.Method`) and call (`T.Method()`) a method at the same time. But these are distinct operations, and sometimes it’s useful to separate them. A selector yields a *method value*, a function bound to a specific receiver value. You can then invoke the method later with any necessary non-receiver arguments. Here’s an example:

```go
p := Point{1, 3}
q := Point{1, 4}
distanceFromP := p.Distance
fmt.Println(distanceFromP(q))
```

Method values come in handy when a package requires a function value as an argument, and the client code wwants to call a method with a specific receiver. They use `time.AfterFunc`, which calls a function (it’s last argument) after a specified delay.

```go
type Rocket struct { /* ... */ }
func (r *Rocket) Launch() { /* ... */ }
r := new(Rocket)
time.AfterFunc(10 * time.Second, r.Launch)
// The alternative is ugly and complicated
// time.AfterFunc(10 * time.Second, func() { r.Launch() })
```

You can also capture the *method expression*. A method expression is `T.f` or `(*T).f`. Using a method expression, you can pass in the receiver as the first argument. As a result, a method expression looks more like a function call rather than a method call. Here’s an example.

```go
p := Point{1, 3}
q := Point{1, 4}
distance := Point.Distance
fmt.Println(distance(p, q) // Replaces p.Distance(q)

scale := (*Point).ScaleBy
scale(&p, 2) // Replaces p.ScaleBy(2)
```

The authors explain that “[m]ethod expressions can be helpful when you need a value to represent a choice among several methods belonging to the same type so that you can call the chosen method with many different receivers” (164-165). They give the following example.

```go
func (path Path) TranslateBy(offset Point, add bool) {
    var op func(p, q Point) Point
    if add {
        op = Point.Add
    } else {
        op = Point.Sub
    }
    for i := range path {
        // Call either path[i].Add(offset) or path[i].Sub(offset)
        path[i] = op(path[i], offset)
    }
}
```

## Encapsulation

Encapsulation, also known as *information hiding* is an important part of object-oriented programming, and Go supports it in a characteristically direct manner. As the authors explain, “Go has only one mechanism to control the visibility of names: capitalized identifiers are exported from the package in which they are defined, and uncapitalized names are not” (168). The same rule applies to functions, methods, and the fields of structs.

The authors list three benefits of encapsulation.

+ You need only consider a limited number of statements “to understand the possible values” of variables “because clients cannot directly modify the object’s variables” (168). In other words, because you limit the places where an object can change, you simplify understanding of that object in use.
+ Clients need not and cannot depend on implementation details, and designers therefore can change those implementation details as much as they like.
+ Clients cannot arbitrarily set an object’s variables, so the designer of a package can guarantee various forms of control and coherence. They give the counter code below as an example.

        ```go
        type Counter struct { n int }
        func (c *Counter) N() int { return c.n }
        func (c *Counter) Increment() { c.n++ }
        func (c *Counter) Reset() { c.n = 0 }
        ```

Given this design, a client can increase a counter or set one back to zero, but they cannot change the value at will. This makes it easier to reason about what a counter object will do. The encapsulation also guarantees that a counter can never be negative.

If you didn’t need encapsulation, you wouldn’t need to create a struct with a single field. But if you exported the object as an `int`, you would lose encapsulation and its benefits. Thus, you will often see single-field structs for exactly this reason.

Methods that access or change the internal values of a type are traditionally called *getters* and *setters*. By convention in Go, you should omit *get*, *fetch*, *find*, *lookup*, or the equivalent in the names of methods that access internal values. They give the following example from the standard library’s `log` package.

        ```go
        package log

        type Logger struct {
            flags int
            prefix string
            // ...
        }

        func (l *Logger) Flags() int                // getter
        func (l *Logger) SetFlags(flag int)         // setter
        func (l *Logger) Prefix() string            // getter
        func (l *Logger) SetPrefix(prefix string)   // getter
        ```

Go does not prevent you from exporting fields: give one a capital letter, and it is exported. However, you should think carefully about when to export fields. In addition, you should not change your choice once you release something.

As an example of a usefully exported field, they mention `time.Duration`. Because durations are exported as `int64` numbers, you can perform math on them and compare them easily in code. For example:

        ```go
        const day = 24 * time.Hour
        fmt.Println(day.Seconds()) // prints 86400
        ```
