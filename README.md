# ticket-booking-go

Ticket booking app for learning Golang. This is my first time using Golang and will explore numerous features of the language.

It will make use of benefits of Go such as concurrency and parallelism, which can be typically found as problems which need be tackled in a ticket booking system. For example, a user may be able to book a ticket while another user is also booking a ticket. This will require that we make sure no two users can book the same ticket at the same time.


**Features that will be explored include:**
- Common Data Types
- Variables and Constants
- Formatted Output
- User input
- Pointers
- Scope Rules
- Loops
- If-else & Switch
- Functions
- Packages
- Goroutines


## About Go

Go is a statically typed, compiled programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. Go is syntactically similar to C, but with memory safety, garbage collection, structural typing, and CSP-style concurrency.

Go code is organised into packages. A package is a collection of source files in the same directory that are compiled together. Functions, types, variables, and constants defined in one source file are visible to all other source files within the same package. All code must be included in a package. An example of this is `Print` from the `fmt` package. This is used to print to the console.

Go is a compiled language. Source code is compiled into a binary file that can be executed on a machine. This is in contrast to interpreted languages, which are not compiled but rather read and executed by an interpreter.

Go code needs to have a declaration of the main package and main function. The main function is the entry point of the program. The main package tells the Go compiler that the package should compile as an executable program instead of a shared library.


## Initialising a Go project

A Go project must be initialised before any Go code can be ran. This is done by creating a `go.mod` file in the root directory of the project. This can be done by running the following command:

```bash
go mod init <project-name>
```

