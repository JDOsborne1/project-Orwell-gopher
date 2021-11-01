package main

import "fmt"

type Field struct {
	size int
	vals [][]int
	desc string
}

type Field_interface interface {
	call_desc()
	print_vals()
	call_size()
}

func (f Field) call_desc() {
	fmt.Println(f.desc)
}

func (f Field) print_vals() {
	fmt.Println(f.vals)
}

func (f Field) call_size() {
	fmt.Println(f.size)
}

func main() {

	var test_f Field_interface = Field{size: 4, vals: [][]int{{1, 2, 3, 4}, {4, 3, 2, 1}, {3, 4, 5, 6}, {6, 5, 4, 3}}, desc: "a testing field"}
	test_f.print_vals()
	test_f.call_desc()
	test_f.call_size()
}
