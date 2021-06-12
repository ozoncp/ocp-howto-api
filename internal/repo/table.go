package repo

type howtoColumns struct {
	id       string
	courseId string
	question string
	answer   string
}

func (c *howtoColumns) ordered() []string {
	return []string{c.id, c.courseId, c.question, c.answer}
}

type howtoTable struct {
	name    string
	columns howtoColumns
}
