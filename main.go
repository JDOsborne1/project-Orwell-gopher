package main

import "fmt"

type Field struct {
	vals [][]int
	desc string
}

type Field_interface interface {
	call_desc()
	print_vals()
}

func (f Field) call_desc() {
	fmt.Println(f.desc)
}

func (f Field) print_vals() {
	fmt.Println(f.vals)
}

func main() {

	var test_f Field_interface = Field{vals: [][]int{{1, 2, 3, 4}, {4, 3, 2, 1}}, desc: "a testing field"}
	test_f.print_vals()
	test_f.call_desc()
}
