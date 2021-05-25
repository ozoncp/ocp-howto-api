package howto

import "fmt"

type Howto struct {
	Id       uint64
	CourseId uint64
	Question string
	Answer   string
}

var nextId uint64

func init() {
	// Get initial value from database or whatever else
	nextId = 42
}

func getNextId() uint64 {
	current := nextId
	nextId++
	return current
}

func New(courseId uint64, question string, answer string) *Howto {
	return &Howto{
		getNextId(),
		courseId,
		question,
		answer,
	}
}

func (h Howto) String() string {
	return fmt.Sprintf("Q: %v\r\nA: %v", h.Question, h.Answer)
}
