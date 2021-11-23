package list

type singlyLL struct {
	front, back *node
}

type node struct {
	val  int
	next *node
}

func (l *singlyLL) PushFront(val int) {
}
func (l *singlyLL) PushBack(val int) {

}
func (l *singlyLL) Reverse() {

}
func (l *singlyLL) Remove(val int) {

}
func (l *singlyLL) String() string {
	return ""
}

func newSinglyLL() List {
	return &singlyLL{}
}
