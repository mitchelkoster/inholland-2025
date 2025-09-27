# InHolland Golang 2025
An interactive workshop for InHolland to get familiar with Go. The focus is on hands-on exercises that highlight Go’s strengths and uniqueness.

## 0x00 - Introduction
- Whoami
- What will we learn?
- What is Golang
- What makes Golang unique?

## 0x01 - The Golang Environment
Students configure their environment, get familiar with the CLI, and build their first program.

### Theory
- Setting up the Go environment
- Go CLI tools
- Go modules & packages

### Practice
- Program entry point _(main)_
- Create a _"Hello World!"_ program
- Use tooling: `go run`, `go build`, `go doc`
- Cross-compile for another platform

## 0x02 - The (boring) basics
Minimal building blocks for writing simple programs. Ends with a character substitution cipher challenge.

### Theory
- **Varibles, constants & data types**
  - Declaring variables and arrays
  - Initializing variables and arrays
  - Shorthand (Walrus notation `:=`) and scope
  - Using constants
- **Control flows: if/else, switch, for, range, defer**
  - A quick note on if/else
  - Switches: Basic switches & behavior
  - Switches: Multiple Expressions
  - Switches: No conditions needed
  - Loops: For
  - Loops: while / Infinite
  - Loops: Ranges
  - Defer(ring execution)
- **Functions**
  - Function header notation
  - Exporting functions
  - Consecutive parameters
  - Mutiple return values
  - Named return values and when to use
  - Variadic functions
  - Anonymous functions

### Practice
- Build a hashmap (`a–zA–Z` substitution)
- Write a decode function that returns both the decoded string and a checksum validation (`bool`)

## 0x03 - Working with files & Error handling
### Theory
- Reading from a file
- Writing to a file
- Error handling vs exceptions
- The idiomatic if err != nil

### Practice
- Handle user input
- Open a file
- Read and display file contents
- Close file and handle errors

## 0x04 - Advanced data types
Work with more advanced structures by reading and parsing webserver logs into data models.

### Theory
- Slices, maps, enums
- Structs and custom types

### Practice
- Read log file from disk
- Parse data into struct models

## 0x05 - Computer Architecture

### Theory
- Stack vs Heap
- Pointers in Go
- Pass-by-value vs pass-by-reference
- Struct methods and pointer receivers

### Practice
- Extend the log parser with enrichments (geo-IP lookup, user agent parsing)
- Compare pass-by-value vs pass-by-reference implementations
- Measure performance differences (benchmark or timing)

## 0x06 - Golang != OOP
No classes, but composition and interfaces. Applied to a website uptime monitoring tool.

### Theory
- Inheritance v.s. Composition (no classes)
- Interfaces
- Packages & Program Structure

### Practice
Let's build an application that monitors the top 100 websites in the Netherlands, whether they are reachable and what their status code is. In order for the application to process the data let's use a "Observer Pattern" to:

- Print the status to the console for debugging
- Store the current state in a CSV file
- etc

## 0x07 - Concurrency!
The crown jewel of Golang, let's speed up our website monitoring tool! 

### Theory
- Goroutines
- Channels
- Worker pools
- Mutexes (when needed, vs channels)
- Detecting and preventing issues:
- Deadlocks & Race Conditions (Dirty Cow Exploit)
- Race detector (go run -race)

### Practice
Update the uptime monitor to use concurrency:

- Implement a worker pool to check 10 websites in parallel
- Send results back via channels
- Use a dedicated writer goroutine for CSV (no mutex needed)
- Print status to console for debugging

_(Optional Experiment):_ Introduce a race condition and use `-race`
_(Optional Experiment):_ Introduce a deadlock 
