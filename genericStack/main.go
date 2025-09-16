// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type stack[T any] struct {
	data []T
	size int
}

func (st *stack[T]) push(el T) {
	st.data = append(st.data, el)
	st.size++
}
func (st *stack[T]) top() T {
	if st.size <= 0 {
		panic(0)
	}

	return st.data[st.size-1]

}

func (st *stack[T]) pop() {
	if st.size <= 0 {
		panic(0)
	}
	st.data = st.data[:st.size-1]
	st.size--
}
func (st *stack[T]) empty() bool {
	return st.size <= 0
}

func (st *stack[T]) sizes() int {

	return st.size
}

func newStack[T any]() *stack[T] {
	return &stack[T]{}
}
func main() {
	fmt.Println("generic stack implementaion")
	stackINT := newStack[int]()
	stackINT.push(1)

	fmt.Println(stackINT.top())
	stackINT.push(88)
	fmt.Println(stackINT.top())
	fmt.Printf("size: %d\n", stackINT.sizes())
	stackINT.push(33)
	stackINT.push(44)
	stackINT.push(55)
	fmt.Println(stackINT.top())
	fmt.Printf("size: %d\n", stackINT.sizes())
	stackINT.pop()
	fmt.Println(stackINT.top())
	fmt.Printf("size: %d\n", stackINT.sizes())

	// string stack
	stackString := newStack[string]()
	stackString.push("prasang")
	stackString.push("mani")
	fmt.Printf("top: %s, size: %d", stackString.top(), stackString.sizes())

}
