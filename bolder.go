package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := "Concurrency is a crucial aspect of modern software development as it enables programs to handle multiple tasks simultaneously efficiently. You can write programs that execute multiple operations leading to improved performance, responsiveness, and resource utilization.\n\nConcurrency is one of the features responsible for Go’s rapid adoption. Go’s built-in support for concurrent programming is considered straightforward while helping avoid common pitfalls like race conditions and deadlocks.\n\n## Introduction to Concurrency in Go\n\nGo provides robust support for concurrency through various mechanisms, all available in its standard library and toolchain. Go programs achieve concurrency through goroutines and channels.\n\nGoroutines are lightweight, independently executing functions that run concurrently with other goroutines within the same address space. Goroutines allow multiple tasks to progress concurrently without the need for explicit thread management. Goroutines are lighter than operating system threads, and Go can efficiently run thousands or even millions of goroutines simultaneously.\n\nChannels are the communication mechanism for coordination and data sharing between goroutines. A channel is a typed conduit that allows goroutines to send and receive values. Channels provide synchronization to ensure safe data sharing between goroutines while preventing race conditions and other common concurrency issues.\n\nBy combining goroutines and channels, Go provides a powerful and straightforward concurrency model that simplifies the development of concurrent programs while maintaining safety and efficiency. These mechanisms enable you to easily take advantage of multicore processors and build highly scalable and responsive applications.\n\n## How to Use Goroutines for Concurrent Code Execution\n\nThe Go runtime manages goroutines. Goroutines have their own stack, allowing them to have a lightweight footprint with an initial stack size of a few kilobytes. \n\nGoroutines are multiplexed onto a small number of OS threads by the Go runtime. The Go runtime scheduler schedules them onto available threads by efficiently distributing the workload, allowing concurrent execution of multiple goroutines on fewer OS threads.\n\nCreating goroutines is straightforward. You’ll use the **go** keyword followed by a function call to declare goroutines.\n\n```go\nfunc main() {\n    go function(\n) // Create and execute goroutine for function1\n    go function2() // Create and execute goroutine for function2\n\n    // ...\n}\n\nfunc function1() {\n    // Code for function1\n}\n\nfunc function2() {\n    // Code for function2\n}\n```\n\nWhen the program invokes **function1()** and **function2()** with the **go** keyword, the Go runtime executes the functions concurrently as goroutines. \n\nHere’s an example use of a goroutine that prints text to the console:\n\n```go\npackage main\n\nimport (\n\t\"fmt\"\n\t\"time\"\n)\n\nfunc printText() {\n\tfor i := 1; i <= 5; i++ {\n\t\tfmt.Println(\"Printing text\", i)\n\t\ttime.Sleep(1 * time.Second)\n\t}\n}\n\nfunc main() {\n\tgo printText() // Start a goroutine to execute the printText function concurrently\n\n\t// Perform other tasks in the main goroutine\n\tfor i := 1; i <= 5; i++ {\n\t\tfmt.Println(\"Performing other tasks\", i)\n\t\ttime.Sleep(500 * time.Millisecond)\n\t}\n\n\t// Wait for the goroutine to finish\n\ttime.Sleep(6 * time.Second)\n}\n```\n\nThe **printText** function repeatedly prints some text to the console with a **for** loop that runs five times following a one-second delay between each statement.\n\nThe **main** function starts a goroutine by calling **go printText**, which launches the **printText** function as a separate concurrent goroutine that allows the function to execute concurrently with the rest of the code in the **main** function.\n\nFinally, to ensure that the program doesn't exit before the **printText** goroutine finishes, the **time.Sleep** function pauses the main goroutine for six seconds. In real-world scenarios, you’d use synchronization mechanisms like channels or wait groups to coordinate the execution of goroutines.\n\n## Channels for Communication and Synchronization\n\nGoroutines have built-in support for communication and synchronization through channels, making writing concurrent code easier than traditional threads, which often require manual synchronization mechanisms like locks and semaphores.\n\nYou can think of channels as pipelines for data flow between goroutines. One goroutine can send a value into the channel, and another goroutine can receive that value from the channel. This mechanism ensures that data exchange is safe and synchronized.\n\nYou’ll use the **<-** operator to send and receive data through channels.\n\nHere's an example demonstrating the basic usage of channels for communication between two goroutines:\n\n```go\nfunc main() {\n    // Create an unbuffered channel of type string\n    ch := make(chan string)\n\n    // Goroutine 1: Sends a message into the channel\n    go func() {\n        ch <- \"Hello, Channel!\"\n    }()\n\n    // Goroutine 2: Receives the message from the channel\n    msg := <-ch\n    fmt.Println(msg) // Output: Hello, Channel!\n}\n```\n\nThe channel in the **main** function is an unbuffered channel named **ch** created with the **make()** function. The first goroutine sends the message \"Hello, Channel!\" into the channel using the **<-** operator, and the second goroutine receives the message from the channel using the same operator. Finally, the **main** function prints the received message to the console.\n\nYou can define typed channels. You’ll specify the channel type on creation. Here's an example that demonstrates usage of different channel types:\n\n```go\nfunc main() {\n    // Unbuffered channel\n    ch1 := make(chan int)\n\n    // Buffered channel with a capacity of 3\n    ch2 := make(chan string, 3)\n\n    // Sending and receiving values from channels\n    ch1 <- 42             // Send a value into ch1\n    value1 := <-ch1       // Receive a value from ch1\n\n    ch2 <- \"Hello\"        // Send a value into ch2\n    value2 := <-ch2       // Receive a value from ch2\n}\n```\n\nThe **main** function creates two channels: **ch1** is an unbuffered integer channel, while **ch2** is a buffered string channel with a capacity of 3. You can send and receive values to and from these channels using the `<-` operator (the values have to be of the specified type).\n\nYou can use channels as synchronization mechanisms for coordinating goroutine execution by leveraging the blocking nature of channel operations.\n\n```go\nfunc main() {\n    ch := make(chan bool)\n\n    go func() {\n        fmt.Println(\"Goroutine 1\")\n        ch <- true // Signal completion\n    }()\n\n    go func() {\n        <-ch // Wait for completion signal from Goroutine 1\n        fmt.Println(\"Goroutine 2\")\n    }()\n\n    <-ch // Wait for completion signal from Goroutine 2\n    fmt.Println(\"Main goroutine\")\n}\n```\n\nThe `ch` channel is a boolean channel. Two goroutines run concurrently in the `main` function. Goroutine one signals its completion by sending a `true` value into channel `ch`. Goroutine 2 waits for the completion signal by receiving a value from the channel. Finally, the main goroutine waits for the completion signal from Goroutine two.\n\n## You Can Build Web Apps in Go With Gin\n\nYou can build high-performant web apps in Go with Gin while leveraging Go’s concurrency features. \n\nYou can use Gin to handle HTTP routing and middleware efficiently and capitalize on  Go's built-in concurrency support by employing goroutines and channels for tasks like database queries, API calls, or other blocking operations."
	output := replaceInlineCodeWithBold(input)
	outputFile := "output.md"
	err := os.WriteFile(outputFile, []byte(output), 0644)
	if err != nil {
		fmt.Println("Error writing output file:", err)
	}
	fmt.Println(output)
}

func replaceInlineCodeWithBold(input string) string {
	var result strings.Builder
	inCodeBlock := false

	for i := 0; i < len(input); i++ {
		if input[i] == '`' {
			// Check if there is a backtick before or after the current backtick
			prevBacktick := false
			nextBacktick := false

			if i > 0 && input[i-1] == '`' {
				prevBacktick = true
			}

			if i < len(input)-1 && input[i+1] == '`' {
				nextBacktick = true
			}

			if prevBacktick && nextBacktick {
				// It's a code block
				if inCodeBlock {
					result.WriteString("```")
					inCodeBlock = false
				} else {
					result.WriteString("``")
					inCodeBlock = true
				}
				i++ // Skip the next backtick since we are handling a code block
			} else if prevBacktick || nextBacktick {
				// It's either the first or last backtick of an inline code
				if !inCodeBlock {
					result.WriteByte(input[i])
				}
			} else {
				// It's inline code, so replace backtick with bold markers
				if !inCodeBlock {
					result.WriteString("**")
				} else {
					result.WriteByte(input[i])
				}
			}
		} else {
			result.WriteByte(input[i])
		}
	}

	return result.String()
}
