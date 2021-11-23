package list

type List interface {
	PushFront(val int)
	PushBack(val int)
	Reverse()
	Remove(val int)
	String() string
}
