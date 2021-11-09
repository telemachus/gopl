# Chapter 11: Testing

Go takes a “low-tech” approach to testing (301). The `go` tool provides a subcommand `go test`, and it requires a small set of conventions concerning file- and test-names. Other than that, writing tests in Go is simply writing Go code. As the authors put it, “this [that is, writing tests] is the same process as writing ordinary Go code; it needn’t require new notations, conventions, and tools” (302).

## The `go test` Tool

The `go` tool employs simple conventions. First, during regular builds, the tool ignores files whose names end in `_test.go`. Second, `go test` only runs files whose names end in `_test.go`. Third, within test files (i.e., files whose names end in `_test.go`), three types of functions receive special treatment: tests, benchmarks, and examples. Test functions have names that begin with `Test`. `go test` runs these functions and reports whether they pass or fail. (I.e., `go test` is a test runner.) Benchmark functions begin with `Benchmark`. If you call `go test` with the benchmark flag enabled, these functions serve to measure performance. Finally, example functions begin with `Example`. Such functions offer documentation through examples, and the documentation is tested for correctness. As a fourth convention, all `_test.go` files must import the `testing` package.

## `Test` Functions

Test functions have the following signature.

        ```go
        func Test<Name>(t *testing.T) {
            // …
        }
        ```

Test functions must begin with `Test`. The name suffix is optional, but if you include it, the suffix must start with a capital letter.

        ```go
        func TestSin(t *testing.T) { // … }
        func TestCos(t *testing.T) { // … }
        func TestLog(t *testing.T) { // … }
        ```

The `t` parameter gives you a hook to methods to report test failures and log other information during testing. Within tests, for example, you will often use `t.Error` or `t.Errorf` to signal and display failures. Execution will continue after individual errors that call `t.Error` or `t.Errorf`. If you want execution to stop, use `t.Fatal` or `t.Fatalf` instead. (You can also use `t.Skip` if you need to skip something under specified conditions.)

Many Go programmers use a table-driven style of testing. In this style, you create a bundle of tests in an anonymous, relatively simple struct. Then you iterate over the struct using `range`. See the palindrome directory for an example.

### Randomized Testing

The authors demonstrate how to use randomized testing with their `IsPalindrome` function. It’s clever, but for now, I don’t think I’ll use this sort of thing.

### Testing a Command

You can use Go’s testing package to test commands as well as libraries. The authors demonstrate with their `echo` program. See the echo directory for details, but I’ll add notes here too.

They rely on both files declaring `package main`, which probably works fine for small commands. For larger ones,I would rather that the test file imports the command as a library.

### White-Box Testing

In black-box testing, the test files import the code being tested exactly like a client. The test code makes no assumptions about the internals of the code under test, and the test code cannot see or interact directly with private methods or members in the code under test.

In white-box testing, however, the test code has access to everything in the code under test. (Better names might be *opaque* versus *transparent* testing.)

As the authors note, the two styles of testing are complementary, and they can both be useful in different situations. Black-box testing requires less changes (since it makes no assumptions about internals), and it helps you to see your own code like a client would. White-box testing can simplify unimportant details for tests, and it can avoid unwanted side effects such as changing a database or creating and deleting files.

The authors demonstrate with tests that use global variables to redefine core methods. They don’t seem worried by this, but I’m guessing a lot of people would be. I don’t know what to think.

### External Test Packages

In this section, the authors introduce the use of `x_test.go` files. They call these *external tests* because the test files declare a package outside of the package that they are testing. The authors do not focus on black-box testing in their presentation. They focus on avoiding cyclical imports. However, external tests can serve both functions.

However, what if you need an external test to avoid cycles, but you also want to do (some) white-box testing? In that case, the authors explain that you can create test files purely to export some internals. For example, here is a file in the `fmt` package. This is the entire `export_test.go` file.

        ```go
        package fmt

        var IsSpace = isSpace
        ```

This mechanism allows you to selectively export internals for testing.

### Writing Effect Tests

Go approaches testing differently than many other languages. They provide a minimal testing library that lacks the bells and whistles (custom assertions, a DSL, setup and teardown functions, etc.) of other languages and libraries. You may or may not like this, but it isn’t an accident. The Go attitude is that test writers should “do most of this work themselves, defining functions to avoid repetition, just as they would for ordinary programs” (316). They continue, “The key to a good test is to start by implementing the concrete behavior you want and only then use functions to simplify the code and eliminate repetition. Best results are rarely obtained by starting with a library of abstract, generic testing functions” (317).

### Avoiding Brittle Tests

The authors define a *brittle test* as one that “spuriously fails when a sound change was made to the program” (317). How do they recommend that you avoid this?

+ Only check the properties you care about.
+ Test simpler and more stable interfaces.
+ Don’t test internal functions.
+ Be selective in your assertions.
+ Don’t check for exact string matches; look for relevant substrings instead.

## Coverage

The `go` tool provides coverage tools under `go test`. You run it once as `go test -cover`. If you want a detailed HTML report, you can run `go test --coverprofile=coverage.out` and then `go tool cover -html=coverage.out`. For more details, see `go tool cover`. The authors urge you not to aim for 100% coverage.

## Benchmark Functions

Go provides benchmark facilities as well as testing facilities. A basic benchmark looks like the following.

        ```go
        import testing

        func BenchmarkIsPalindrome(b *testing.B) {
            for i := 0; i < b.N; i++ {
                IsPalindrome("A man, a plan, a canal: Panama")
            }
        }
        ```

You have to ask for benchmarks with a flag: `go test -bench=<regex or . for all benchmarks>`.

The authors show that “the fastest program is often the one that makes the fewest memory allocations” (322). First they abbreviate a loop in the program, but that yields on a 4% improvement. Then they initialize an array of the right size so that `append` never needs to reallocate. That yields an improvement of 35%. Since allocations are this important, you can also add the `-benchmem` flag to see how many allocations occur per operation: `go test -bench=. -benchmem`. They also point out that comparative, relative benchmarks are often more informative than single, absolute benchmarks.

        ```go
        func benchmark(b *testing.B, size int) { /* … */ }
        func Benchmark10(b *testing.B) { benchmark(b, 10) }
        func Benchmark100(b *testing.B) { benchmark(b, 100) }
        func Benchmark1000(b *testing.B) { benchmark(b, 1000) }
        ```

A final warning: do not use `b.N` as input size. If you mess with `b.N`, your results will be garbage.

## Profiling

This material is beyond me right now, so I will come back to it later if I need it.

## `Example` Functions

You can include testable example functions in your Go files. They look like this:

        ```go
        func ExampleIsPalindrome() {
            fmt.Println(IsPalindrome("A man, a plan, a canal: Panama"))
            fmt.Println(IsPalindrome("palindrome"))
            // Output:
            // true
            // false
        }
        ```

As the authors explain, example functions offer three benefits.

1. Example functions provide documentation. A single example can be more useful than paragraphs of text. Since example functions are always run when compiled, they don’t go stale the way that examples in comments can.
1. Example functions are executable tests that run when you call `go test`. If an example function fails, the test runner will show the failure and exit with a non-success exit value.
1. Example functions provide hands-on experimentation. If you use `godoc` to examine a file with example functions, you can view, run, edit, and rerun examples. This provides a great way to learn by noodling around with executable examples of real Go code.
