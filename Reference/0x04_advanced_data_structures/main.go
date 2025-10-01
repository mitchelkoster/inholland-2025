package main

import "fmt"

// Custom type
type HTTPStatus int

// Enum
const (
	OK                  HTTPStatus = 200
	Created             HTTPStatus = 201
	BadRequest          HTTPStatus = 400
	Unauthorized        HTTPStatus = 401
	NotFound            HTTPStatus = 404
	InternalServerError HTTPStatus = 500
)

func handleRequest(code HTTPStatus) {
	switch code {
	case OK:
		fmt.Println("Everything is fine")
	case Created:
		fmt.Println("Resource created")
	case BadRequest:
		fmt.Println("Client sent a bad request")
	case Unauthorized:
		fmt.Println("Unauthorized access")
	case NotFound:
		fmt.Println("Resource not found")
	case InternalServerError:
		fmt.Println("Server error")
	default:
		fmt.Println("Unknown status:", code)
	}
}

type Response struct {
	URL  string
	Code HTTPStatus
	Body string
}

func (s HTTPStatus) IsError() bool {
	return s >= 400
}

// Print summary only if it’s an error
func (r Response) PrintIfError() {
	if r.Code.IsError() {
		fmt.Printf("[ERROR] %s -> %d\n", r.URL, r.Code)
	}
}

// Recieve by value (does NOT mutate)
func (r Response) Summary() {
	fmt.Printf("URL: %s, Code: %d, Body: %s\n", r.URL, r.Code, r.Body)
}

// Pointer receiver (mutates struct)
func (r *Response) SetCode(code HTTPStatus) {
	r.Code = code
	fmt.Println("Status code updated to", r.Code)
}

func main() {
	responses := []Response{
		{"https://example.com", OK, "Welcome"},
		{"https://example.com/missing", NotFound, "Not here"},
		{"https://example.com/fail", InternalServerError, "Oops"},
	}

	for _, r := range responses {
		r.PrintIfError() // only prints the 404 and 500
	}

	type Person struct{ Age int }
	elizabeth := Person{30}
	pElizabeth := &elizabeth

	// Equivalent to (*pElizabeth).Age = 27
	pElizabeth.Age = 27

	x := 10
	p := &x // p is a pointer to x
	*p = 20 // dereference p and write 20 into x

	resp := Response{
		URL:  "https://example.com",
		Code: OK,
		Body: "Hello World!",
	}

	resp.Summary()         // 200
	resp.SetCode(NotFound) // changes the Code field
	resp.Summary()         // 404

	handleRequest(OK) // Everything is fine

	var numbersSlice []int         // nil
	fmt.Println(len(numbersSlice)) // 0

	// Append to add items
	numbersSlice = append(numbersSlice, 1)
	numbersSlice = append(numbersSlice, 2, 3, 4, 5)

	// If you love the walrus operator
	newSlice := make([]string, 2)

	newSlice[0] = "Hello, "
	newSlice[1] = "World!"

	// Slice selection
	tmpSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(tmpSlice[5:])  // [6 7 8 9 10]
	fmt.Println(tmpSlice[:5])  // [1 2 3 4 5]
	fmt.Println(tmpSlice[2:7]) // [3 4 5 6 7]

	// Deleting slice items a chore!
	nums := []int{10, 20, 30, 40}

	// remove element at index 1 (value 20)
	i := 1
	nums = append(nums[:i], nums[i+1:]...)

	// The other option is target matching
	languages := []string{"go", "rust", "python", "go"}
	target := "go"

	out := []string{}
	for _, lang := range languages {
		if lang != target {
			out = append(out, lang)
		}
	}

	// Maps
	people := map[string]int{
		"Mitchel":   36,
		"Elizabeth": 25,
		"Esmée":     4,
	}

	// Guess the order
	fmt.Println("ORDER:")
	for name, age := range people {
		fmt.Println(name, age)
	}

	// Zero if not found
	fmt.Println("Mitchel", people["Mitchel"])           // 21
	fmt.Println("i-dont-exit:", people["i-dont-exist"]) // 0 (not found!)

	// The "comma ok idiom"
	if age, ok := people["Mitchel"]; ok {
		fmt.Println("Mitchel found:", age)
	} else {
		fmt.Println("Mitchel not found")
	}

	// Alter maps
	people["Dave"] = 40      // Insert new value
	people["Elizabeth"] = 32 // Update

	// Guess the order
	for name, age := range people {
		fmt.Println(name, age)
	}

	// Delete
	delete(people, "Mitchel")
	fmt.Println("After delete:", people)
}
