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
- Panic & Recover

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
 - _Parse_ the log into a data model (`struct`)

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
### Theory
- Inheritance vs Composition (no classes)
- Interfaces
- Packages & Program Structure

### Practice
- Allow the user to provide data to populate structs with web crawler info