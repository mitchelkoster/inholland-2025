# InHolland Golang 2025
The InHolland Golang Curiculum designed to be a interactive workshop to get familiar with Golang.

## 0x00 - Introduction
- Whoami
- What will we learn?
- What is Golang
- What makes Golang unique?

## 0x01 - The Golang Environment
Have everyone configure their environment, get familiar with the CLI tooling and build the first program.

### Theory
- Setting up a Golang environment
- Golang CLI tools
- Go modules and the Packages

### Practice
- Program entry points
- Create a `"Hello World!"` program
- Golang tooling `run`, `build`, `doc`
- Test the program with `go run`
- Cross-compile to multiple platforms

## 0x02 - The (boring) basics
Explain minimal core concepts to be able to build simple programs. This chapter is ended with a challenge to solve a _character substitution_ cipher which should result in a matching _checksum_.

### Theory
- **Varibles, constants & data types**
  - Declaring variables and arrays
  - Initializing variables and arrays
  - Shorthand (Walrus notation) and scope
  - Using constants
- **Control flows**
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
The cipher text has been encoded with a _character substition_, to solve this create a _legenda_ of substituded characters and _validate_ for correctness by comparing the `checksum`.

 - Build a hashmap of every character (`a-zA-Z`) and their substitions
 - Create a function which retures  the decoded string `string` and the checksum result `bools`

## 0x03 - Working with files & Error handling
### Theory
- Reading from a file
- Writing to a file
- Error handling vs exceptions
- The "if err != nil" idiom

### Practice
- Handle user input
- Open the file
- Read and display file contents
- Close the file

## 0x04 - Advanced data types
Get a good grasp on more advanced data structures and use this to _read webserver logs from disk_ and _parse them into data models_ for use within the application.

### Theory
- **Advanced data types**
  - Slices
  - Maps
  - Enums
  - Structs
  - Custom Data Types

### Practice
 - _Read_ the `file` from disk
 - _Parse_ the log data into a list of models

## 0x05 - Core Computing Concepts

### Theory
- Computer Architecture: Stack & Heap
- Pointers in Golang
- Pass-by-value & pass-by-reference
- Structs Methods
- Struct Methods & Pointers

### Practice
To really show the difference between `pass-by-value` and `pass-by-reference` use the web parsed webserver logs and write a function that attaches a geo-location to IP-address and a parse the user agent to make it more clear.

  - Update the program
  - A pass-by-value & benchmark
  - A pass-by-reference & benchmark

## 0x06 - Golang != OOP
So far we have seen no classes, but we do have interfaces! Let's combine everything we've learned to build a website up-time monitoring tool using a design pattern.

### Theory
- Inheritance vs Composition (no classes)
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
- Mutexes
- Worker Pools
- Packages & Program Structure

### Practice
Update the web monitoring program to work with a worker pool so we can visit all 100 websites a lot faster! (does a channel still work? Should it be mutext when writing CSV, etc)

- Print the status to the console for debugging
- Store the current state in a CSV file
- etc