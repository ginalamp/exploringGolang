// methods and interfaces
package main

import (
	"fmt"
	"image"
	"io" // read data
	"math"
	"strings"
	"time"
)

// non-struct type declaration (can only have methods of types within the same package)
type MyFloat float64

// methods have a special receiver argument
// func <receiver> <functionName>() <returnType> {}
func methods() {
	basicsMethods()
	pointerReceivers()
}

// vertex with (x,y) coordinates
type Vertex struct {
	X, Y float64
}

// calling methods
func basicsMethods() {
	// create vertex and compute value
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	// absolute value of non-struct type
	f := MyFloat(-math.Sqrt2) // convert float to MyFloat
	fmt.Println(f.Abs())
}

// get Vertex square root of (X^2 + Y^2)
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// get the absolute value of f
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// better performance (avoid copying values)
// methods can receive pointers (not for pointers to pointers)
// can modify the value to which the receiver points
// can send value or pointer to method (automatically converts)
func pointerReceivers() {
	// send Vertex value
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())

	// send pointer to Vertex
	p := &Vertex{4, 3}
	p.Scale(3)
	fmt.Println(p.Abs())
}

// receives pointer and modifies values to multiply with f
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// set of method signatures
// type <interfaceName>er interface {}
func interfaces() {
	basicsInterfaces()
	interfaceValues()
	types()
	keyExamples()
	readers()
	images()
}

// all values of type float64 automatically implements Abs()
type Abser interface {
	Abs() float64
}

// creating of interfaces
func basicsInterfaces() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f // a MyFloat implements Abser
	fmt.Println(a.Abs())
	a = &v // a *Vertex implements Abser
	fmt.Println(a.Abs())
	// TODO: why does the below run? Find out why it shouldn't.
	a = v // a Vertex does NOT implement Abser
	fmt.Println(a.Abs())

	// implicit interface implementation
	var i I = X{"hello"}
	i.M()
}

// all types can implement M()
type I interface {
	M()
}

// struct with type string S
type X struct {
	S string
}

// struct with type string S
type T struct {
	S string
}

// X implements interface I and thus implements M() (don't need to declare explicitly)
func (x X) M() {
	fmt.Println(x.S)
}

// T implements interface I and thus implements M()
func (t *T) M() {
	// should cater for underlying nil values
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

// MyFloat implements interface I and thus implements M()
func (f MyFloat) M() {
	fmt.Println(f)
}

// interface values tuple (<value>, <type>)
func interfaceValues() {
	// Calling a method on an interface value executes the method of the same name on its underlying type.
	var i I
	i = &T{"Hello"}
	describe(i)
	i.M()

	i = MyFloat(math.Pi)
	describe(i)
	i.M()

	// nil underlying values (note nil interface value causes runtime exception)
	var t *T
	i = t
	describe(i)
	i.M() // outputs <nil> due to explicit if statement catering for it

	// empty interface holds any value type (for unknown types)
	var emptyInterface interface{}
	describeUnknown(emptyInterface)
	emptyInterface = 42
	describeUnknown(emptyInterface)
	emptyInterface = "hello"
	describeUnknown(emptyInterface)
}

// interface type assertions and switches
// type assertion: access interface's underlying concrete value
// type switch: permits several type assertions in series
func types() {
	// type assertion
	var i interface{} = "hello"
	s := i.(string) // assert that i holds the concrete type string
	fmt.Println(s)
	// test if interface holds a specific type
	s, ok := i.(string) // (<underlyingValue>, <boolean>)
	fmt.Println(s, ok)
	f, ok := i.(float64)
	fmt.Println(f, ok)
	// f = i.(float64) // error since i does not hold a float64
	// fmt.Println(f)

	// type switches
	typeSwitch(21)
	typeSwitch("hello")
	typeSwitch(true)
}

// type switch - do different things depending on the type (unsure which type hence the empty interface parameter)
func typeSwitch(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

// stringer and error interfaces
func keyExamples() {
	stringers()
	errors()
}

// ubiquotous interface Stringer defined by fmt
// Stringer: type that can describe itself as a string
func stringers() {
	h := Person{"Harry Potter", 22}
	v := Person{"Tom Riddle", 90}
	fmt.Println(h, v)
}

// person with name and age
type Person struct {
	Name string
	Age  int
}

// implements interface Stringer (defined by fmt)
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

// ubiquotous interface Error defined by fmt
// express error state
func errors() {
	// if the value returned by run is not nil, then print the value
	if val := run(); val != nil {
		fmt.Println(val)
	}

	// sqrt complex number error check
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

// error with when and what
type MyError struct {
	When time.Time
	What string
}

// output formatted error
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

// running function returns a MyError struct
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

// sqrt error type
type ErrNegativeSqrt float64

// sqrt error string output
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// Newton's method: compute sqrt using loop through guesses with error case
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := x / 2
	z_prev := -z
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)

		// break if value not changing
		if z == z_prev {
			fmt.Println("Exiting loop")
			break
		}
		z_prev = z
		fmt.Println(z)
	}
	fmt.Printf("The sqrt of %g is ~%g\n", x, z)
	return z, nil
}

// io.Reader interface (read stream of data)
func readers() {
	r := strings.NewReader("Hello, Reader")
	b := make([]byte, 8)

	for {
		// Read returns the number of bytes populated and error value (nil if no err)
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		// output value of bytes read
		fmt.Printf("b[:n] = %q\n", b[:n])

		// EOF error when the stream ends
		if err == io.EOF {
			break
		}
	}
}

func images() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}

func main() {
	fmt.Println("Running methods & interfaces")
	methods()
	interfaces()
}

// output value and type for values of type I
func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// output value and type for values of unknown type
func describeUnknown(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
