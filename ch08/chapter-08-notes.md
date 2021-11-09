# Goroutines and Channels

## Channels

Channels provide the connections between goroutines. “A channel is a communication mechanism that lets one goroutine send values to another goroutine” (225).

Channels have types that designates what values are shared through that channel. For example, a channel that transmits items of type `int` is `chan int`. You create a channel using `make`.
    ```go
    ch := make(chan int)
    ```
Channels are reference types—like maps. Channels are passed by copying the reference, “so calledr and callee refer to the same data structure” (225). The zero value of a channel is `nil`. You can compare channels with `==`, and you can compare channels with `nil`.

Channels can send and receive, and sending and receiving are collectively known as *communication*. Both sending and receiving are written using the `<-` operator. If you are sending, `<-` separates channel and value operands. If you are receiving, `<-` precedes the channel operand. You can use a receive expression without assigning the value that is received. In such a case, the value is discarded, but the statement is valid.
    ```go
    ch <- x     // a send statement
    x = <-ch    // a receive expression in an assignment statement
    <-ch        // a receive statement where the result is discarded
    ```

You can also close a channel. Once a channel is closed, any attempt to send to this channel will panic. However, you can receive from a closed channel. Initially, the channel will send any values that are backed up in the reception queue; after that, receive operations will get the zero value of the channel’s element type. You close a channel using `close`: `close(ch)`.

Channels can be buffered or unbuffered. When you create a channel, `make` accepts an integer as an optional second argument. If you pass any argument other than `0`, the channel has a capacity of that number, and it is buffered.
    ```go
    ch = make(chan int)     // unbuffered channel
    ch = make(chan int, 0)  // unbuffered channel
    ch = make(chan int, 3)  // buffered channel with capacity of 3
    ```

### Unbuffered Channels

When you send or receive on an unbuffered channel, the goroutine is blocked until another goroutine accepts or sends on the same channel. Thus, “unbuffered channels are sometimes called *synchronous* channels” because “[c]ommunication over an unbuffered channel causes the sending and receiving goroutines to *synchronize*” (226).

With that in mind, we can use a channel to signal when a program is finished. Imagine a server that would quit when its input stream closes “even if the background goroutine is still working” (226). A channel can prevent such a server from quitting too soon.
    ```go
    func main() {
        conn, err := net.Dial("tcp", "localhost:8000")
        if err != nil {
            log.Fatal(err)
        }
        done := make(chan struct{})
        go func() {
            io.Copy(os.Stdout, conn)
            log.Println("done")
            done <- struct{}{}
        }()
        mustCopy(conn, os.Stdin)
        conn.Close()
        <-done // Wait for the background goroutine to finish.
    }
    ```
Notice that in this case, we don’t care about the content of the message sent through this channel. We only care that *communication has occurred*. Thus, we use an empty struct to signal that the only reason for this channel is synchronization. (Such messages are often called *events*.) Instead of an empty struct, you can use a channel of `bool` or `int` type.

### Pipelines

In a *pipeline*, the output of one goroutine provides the input of another. Here’s a program that prints squares of numbers for as long as it’s running.
    ```go
    func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	for {
		fmt.Println(<-squares)
	}
    }
    ```

What if we want to stop this program after a certain range of squares? We can close the naturals channel. Since we won’t send to that channel again, the program won’t panic. However, after all the values sent to the channel have been seen, the program will print an infinite number of zeros. Why? Because “after the last sent element has been received, all subsequent receive operations will proceed without blocking but will yield a zero value” (229). 

How can we fix this second problem? We can use the two-result variant of a receive statement. When you receive two values from a channel, the second is a boolean that signals whether the receive comes from a closed and drained channel. (A drained channel is one that has sent all its real values.) This boolean is often called *ok*, like the second argument to a map-indexing expression.

Here is how that looks.
    ```go
    go func() {
        for {
            x, ok := <-naturals
            if !ok {
                break
            }
            squares <- x * x
        }
        close(squares)
    }()
    ```
Because this is a bit noisy, Go provides a shortcut. You can use `range` with a channel. The `range` built-in will handle looping until the channel is closed and drained.

With all that said, here’s a limited squaring program.

    ```go
    func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 1; x <= 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	for x := range squares {
		fmt.Println(x)
	}
    }
    ```

A few more notes about closing channels.
    + You don’t need to close every channel. You should explicitly close a channel only if you want to signal the receiving goroutines that all data has been sent. The garbage collector will automatically handle any channel that is unreachable. (Contrast this with open files. You *should* always close those!)
    + If you call `close` on a closed channel, Go will panic.
    + If you call `close` on a nil channel, Go will panic.

### Unidirectional Channel Types

If we consider the program above, we can see that in some cases we only send to a channel, and in other cases we only receive from a channel. This pattern turns out to be typical: “When a channel is supplied as a function parameter, it is nearly always with the intent that it be used exclusively for sending or exclusively for receiving” (230).

In order to “document this intent and prevent misuse, the Go type system provides *unidirectional* channel types that expose only one or the other of the send and receive operations” (230). You indicate a send-only channel by placing `<-` immediately after `chan`, and you indicate a receive-only channel by placing `<-` immediately before `chan`. For example `chan<- int` indicates a send-only channel, and `<-chan int` indicates a receive-only channel

You can only call `close` on a bidirectional channel or a send-only channel. If you call `close` on a receive-only channel, Go throws a compile-time error.

Here’s an example program that uses unidirectional channels.
    ```go
    func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
    }

    func counter(out chan<- int) {
        for x := 0; x < 101; x++ {
            out <- x
        }
        close(out)
    }

    func squarer(out chan<- int, in <-chan int) {
        for v := range in {
            out <- v * v
        }
        close(out)
    }

    func printer(in <-chan int) {
        for v := range in {
            fmt.Println(v)
        }
    }
    ```
When you create the channel, you don’t specify whether that it is unidirectional. The function calls implicitly convert the channels as necessary. Thus, in `go counter(naturals)`, `naturals` is implicitly converted to a receive-only channel. By contrast, in `go squarer(squares, naturals)`, `naturals` is implicitly converted to a send-only channel. Similarly `squares` is used as a receive-only channel by `squarer`, but `printer` uses the same channel as a send-only channel.

## Looping in Parallel

Let’s start with a fun phrase: *embarrassingly parallel*. The authors define an embarrassingly parallel problem as one that consists “entirely of subproblems that are completely independent of each other” (235). As an example, they describe a program that makes thumbnails of an arbitrary number of specified image files. Here’s a non-concurrent version of such a program:
    ```go
    func makeThumbnails(filenames []string) {
        for _, f := range filenames {
            if _, err := thumbnail.ImageFile(f); err != nil {
                log.Println(err)
            }
        }
    }
    ```
You might try to make it concurrent by adding just one `go`, but the result is broken.
    ```go
    func makeThumbnails(filenames []string) {
        for _, f := range filenames {
            go thumbnail.ImageFile(f) // Ignores errors, but has bigger problems
        }
    }
    ```
Why does the second function do wrong? It returns as soon as it has started all the thumbnail actions, but it doesn’t wait for all (or any!) of them to finish. That’s bad.

How can we fix this? We can ask the inner goroutine to report when it is finished.
    ```go
    func makeThumbnails(filenames []string) {
        ch := make(chan struct{})
        for _, f := range filenames {
            go func(f string) {
                thumbnail.ImageFile(f) // Still ignoring errors…
                ch <- struct{}{}
            }(f)
        }

        // Wait for goroutines to complete.
        for range filenames {
            <-ch
        }
    }
    ```
This works since we know the exact number of tasks up front. (The number is always identical to `len(filenames)`.) We will need other solutions if we don’t know the number of tasks up front.

Before I go on, notice that we must explicitly pass in the filename to the anonymous `go` function. The following does *not* work.
    ```go
    for _, f := range filenames {
        go func() {
            thumbnail.ImageFile(f)
        }()
    }
    ```
This version does not work because `f` is the same variable, but it’s value will change over time. Beware of this sort of error. The GOPL authors suggest two ways to fix this problem. First, you can “use a buffered channel with sufficient capacity that no worker goroutine will block when it sends a message” (237). Second, you can “create another gouroutine to drain the channel while the main goroutine returns the first error without delay” (237). Unfortunately, they don’t give code for either solution, so I am not entirely sure what these fixes should look like.

Last, the authors show whow to use a `sync.WaitGroup` to handle situations where you don’t know in advance how many iterations the loop will run.
    ```go
    func makeThumbnails(filenames <-chan string) int64 {
        sizes := make(chan int64)
        var wg sync.WaitGroup
        for f := range filenames {
            wg.Add(1)
            go func(f string) {
                defer wg.Done()
                thumb, err := thumbnail.ImageFile(f)
                if err != nil {
                    log.Println(err)
                    return
                }
                info, _ := os.Stat(thumb)
                sizes <- info.Size()
            }(f)
        }

        go func() {
            wg.Wait()
            close(sizes)
        }()

        var total int64
        for size := range sizes {
            total += size
        }
        return total
    }
    ```
They point out a few things for special attention.
    + The `Add` method must be called before the goroutine starts.
    + The `Add` method takes a parameter.
    + The `Done` method is called inside the goroutine, and it takes no parameter. `Done()` is equivalent to `Add(-1)`.
    + They use `defer wg.Done()` to make sure that `Done` is called even when there is an error. This is a common pattern for looping in parallel when you don’t know the number of iterations up front.
