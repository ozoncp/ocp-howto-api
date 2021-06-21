package howto

import "fmt"

// Howto - структура, описывающая сущность Howto для Ozon Code Platform
type Howto struct {
	Id       uint64
	CourseId uint64
	Question string
	Answer   string
}

// String преобразует Howto в строку
func (h Howto) String() string {
	return fmt.Sprintf("Q: %v\r\nA: %v", h.Question, h.Answer)
}
