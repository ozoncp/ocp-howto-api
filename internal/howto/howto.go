package howto

import "fmt"

type Howto struct {
	Id       uint64
	CourseId uint64
	Question string
	Answer   string
}

func (h Howto) String() string {
	return fmt.Sprintf("Q: %v\r\nA: %v", h.Question, h.Answer)
}
