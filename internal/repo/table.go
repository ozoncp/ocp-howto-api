package repo

type howtoColumns struct {
	id        string
	course_id string
	question  string
	answer    string
}

func (c *howtoColumns) ordered() []string {
	return []string{c.id, c.course_id, c.question, c.answer}
}

type howtoTable struct {
	name    string
	columns howtoColumns
}
