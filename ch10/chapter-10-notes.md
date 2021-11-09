# Chapter 10: Packages and the Go Tool

## Introduction

Package systems help programmers design and maintain large programs by creating modularity. A package joins together related code into a single unit that can be learned and changed separately from other packages.

Each package has its own name space. Thus, you don’t have to worry about duplicating names for functions, types, or interfaces across packages. Even if two packages have functions with the same name (e.g., `UpdateTag`), users will call those functions under different name spaces (e.g., `something.UpdateTag` and `anotherthing.UpdateTag`). This prevents conflicts, but it also frees you to use the clear, obvious, and brief names in your code.

Packages also control what they export and what they keep private (aka, encapsulation). This allows package writers to define a consistent API while still reserving the right to change implementation however they want. They are only responsible for keeping the public API stable. In addition, by keeping some variables and settings private, authors can control who updates those values to preserve any necessary restrictions.

## Import Paths

Go identifies every package via a unique string (aka, the package’s *import path*). The Go language specification does not define how to interpret these strings or how to build an import path. Individual tools may do this in different ways. However, many (most?) Go programmers work with the `go` tool that comes with the Go runtime. That’s the most important standard.

In order to maintain unique names, you should prefix anything outside of the standard library with an internet domain name of the person or organization who owns or hosts the package.

## The Package Declaration

You must give a package declaration at the start of every file. That package name may end up becoming the last part of an import. For example, both `math/rand` and `crypto/rand` have packages named simply *rand*. (Later, the authors will explain how to use two packages with the same name in one program.)

In general, every package in a single directory should have the same package name. There should be only one package per directory. However, there are some exceptions.

1. A package that defines a command must have the name `main`, so that the `go` tool knows to make an executable from that file. (I think that this rule has changed and these `main` files must now be in another directory.)
1. If a file is named `x_test.go`, then the package name can be `x_test`. Thus you can have both `rand` and `rand_test` packages in one directory.
1. Some packages append version numbers to their import paths, but the package name leaves the suffix off. For example `import "gopkg.in/yaml.v2"`, but the package name is still `yaml` not `yaml.v2`.

## Import Declarations

Go files can contain zero or more imports. The imports must go immediately after the `package` declaration and before any non-import code. If you have multiple imports, you can do them one by one or group them within parentheses. The following are equivalent.

        ```go
        import "fmt"
        import "math/rand"

        import (
            "fmt"
            "math/rand"
        )
        ```

Go programmers conventionally group imports by domain and then order them alphabetically. But this is not required. The `gofmt` and `goimports` tools will do this for you automatically.

If you want to import two (or more) packages with the same name, you must rename one (or more). Here’s how a renaming import looks.

        import (
            "crypto/rand"
            mrand "math/rand"
        )
        ```

You can also use renaming imports if you want a shorter, clearer, or more convenient name for a package.

Finally, Go imports are per file, not per package. If you have ten files in one package, and they all need `fmt` functions, then you must import `fmt` ten times.

## Blank Imports

The `go` tool considers it an error if you import a package into a file but you don’t use the package. (The tool also, more obviously, considers it an error to use a package that you do not explicitly import.) On rare occasions, you may want to import a package for side effects rather than for using the package. On those cases, you can use the blank identifier as the name of the package. E.g., `import _ "math/rand"`. The authors give an excellent small example.

        ```go
        // The jpeg command reads a PNG image from the standard input
        // and writes it as a JPEG image to the standard output.
        package main

        import (
            "fmt"
            "image"
            "image/jpeg"
            _ "image/png" // register PNG decoder
            "io"
            "os"
        )

        func main() {
            if err := toJPEG(os.Stdin, os.Stdout); err != nil {
                fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
                os.Exit(1)
            }
        }

        func toJPEG(in io.Reader, out io.Writer) error {
            img, kind, err := image.Decode(in)
            if err != nil {
                return err
            }
            fmt.Fprintln(os.Stderr, "Input format =", kind)
            return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
        }
        ```

Briefly, this works because the `image/png` package loads a decoder for PNGs in an `init` function. You don’t need any of the functions or variables from `image/png`, but you do need that file type loaded for the program to work.

## Packages and Naming

The authors give some tips and advice about names for and in packages.

First, here is some advice for package names.

+ Keep the package name short, but not short enough to be cryptic.
+ Be descriptive and unambiguous. Don’t use `util` when you can say `ioutil` or `imageutil`, both of which are still reasonably short.
+ Avoid names for packages that are common for local variables or functions. E.g., avoid `path`!
+ Most package names should use a singular. Exceptions like `bytes`, `errors`, `strings`, and `go/types` are required to avoid the names of predeclared types or keywords.
+ Avoid package names that have connotations for something else. E.g., don’t use `temp` for a temperature package. The name `temp` will always make people think of “temporary.”

Next, here is some advice for package members.

+ Since the package name will always appear with the name of the package member, you should always think of them together. That will omit some redundancy, and it allows you more flexibility in what you name the members. By itself, `Get` may be unclear or ambiguous, but `http.Get` is perfectly clear.
+ If your package exposes one major data type and its methods, you should provide a `New` function to create your data type. They give the example of `math/rand`, which exposes a struct type `rand.Rand` and a `rand.New` method to create that struct type.

## The Go Tool

The `go` tool does many jobs. It downloads, install, and lists packages (in the manner of `gem`, `cpan`, or `pip`). It compiles go code. It runts tests.

By default, `go` will assume that `GOPATH` is `$HOME/go` on Unix systems. You want to put Go packages in `GOPATH/src` and then import them as the rest of their path. E.g., `GOPATH/src/gopl.io/ch1/helloworld` is imported as `gopl.io/ch1/helloworld`, and the package name is `helloworld`.

I think that their advice about `go get` is out of date. Per [Go’s 1.16 release notes][go1.16], we should use `go install` rather than `go get` in most cases. The only case where you want `go get` is `go get -d` “to adjust the current module's dependencies without building packages” (from [the release notes][go1.16]). In a nutshell, `go get` used to (1) download a package and then (2) build that package. You shouldn’t use `go get` for this anymore.

[go1.16]: https://tip.golang.org/doc/go1.16#modules
