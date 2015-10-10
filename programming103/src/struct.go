package main

// struct OMIT
type Foo struct {
	Bar string
	Baz uint64
}

// end struct OMIT

// methods OMIT
func (foo Foo) getBar() string {
	return foo.Bar
}

func (foo *Foo) setBar(bar string) {
	foo.Bar = bar
}

// end methods OMIT

func main() {

}
