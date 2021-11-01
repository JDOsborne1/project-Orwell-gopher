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

type Entity struct {
	position [2]float32
	velocity [2]float32
}

type Entity_interface interface {
	get_position()
	get_velocity()
	iterate_position()
}

func (e Entity) get_position() {
	fmt.Println(e.position)
}

func (e Entity) get_velocity() {
	fmt.Println(e.velocity)
}

func (e *Entity) iterate_position() {
	for index, value := range e.position {
		e.position[index] = value + e.velocity[index]
	}

}

func main() {

	//var test_f Field_interface = Field{size: 4, vals: [][]int{{1, 2, 3, 4}, {4, 3, 2, 1}, {3, 4, 5, 6}, {6, 5, 4, 3}}, desc: "a testing field"}
	//test_f.print_vals()
	//test_f.call_desc()
	//test_f.call_size()

	var test_e Entity_interface = &Entity{position: [2]float32{2.2, 1.2}, velocity: [2]float32{0.2, -0.1}}

	test_e.get_position()
	test_e.get_velocity()

	test_e.iterate_position()
	test_e.get_position()
}
