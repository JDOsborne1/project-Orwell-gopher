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
	position   [2]float32
	velocity   [2]float32
	magnitudes map[string]float32
}

type Entity_interface interface {
	get_position()
	get_velocity()
	iterate_position()
	iterate_velocity()
}

func (e Entity) get_position() {
	fmt.Println(e.position)
}

func (e Entity) get_field_position(field_name string) ([2]float32, float32) {

	return e.position, e.magnitudes[field_name]

}

func (e Entity) get_velocity() {
	fmt.Println(e.velocity)
}

func (e *Entity) iterate_position() {
	for index, value := range e.position {
		e.position[index] = value + e.velocity[index]
	}

}

func (e *Entity) iterate_velocity(input [2]float32) {
	for index, value := range e.velocity {
		e.velocity[index] = value + input[index]
	}

}

func (e *Entity) step(input [2]float32) {
	e.iterate_velocity(input)
	e.iterate_position()

}

func main() {

	//var test_f Field_interface = Field{size: 4, vals: [][]int{{1, 2, 3, 4}, {4, 3, 2, 1}, {3, 4, 5, 6}, {6, 5, 4, 3}}, desc: "a testing field"}
	//test_f.print_vals()
	//test_f.call_desc()
	//test_f.call_size()

	//var test_e Entity_interface = &Entity{position: [2]float32{2.2, 1.2}, velocity: [2]float32{0.2, -0.1}}
	var test_e2 = Entity{position: [2]float32{2.2, 1.2}, velocity: [2]float32{0.2, -0.1}, magnitudes: map[string]float32{"gravity": 0.4}}

	fmt.Println(test_e2.get_field_position("gravity"))

	//fmt.Println(test_e2.position)
	//fmt.Println(test_e2.velocity)
	//test_e2.iterate_velocity([2]float32{1, -1})
	//fmt.Println(test_e2.velocity)

	test_e2.get_position()
	test_e2.get_velocity()
	test_e2.step([2]float32{10, 10})
	test_e2.get_position()
	test_e2.get_velocity()
}
