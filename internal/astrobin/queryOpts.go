package astrobin

import "fmt"

type OperatorEnum string

const (
	EQUAL         OperatorEnum = "="
	LESS          OperatorEnum = "__lt="
	LESS_EQUAL    OperatorEnum = "__lte="
	GREATER       OperatorEnum = "__gt="
	GREATER_EQUAL OperatorEnum = "__gte="
)

type QueryOpts struct {
	Name     string
	Operator OperatorEnum
	Value    interface{}
}

func (q *QueryOpts) String() string {
	return fmt.Sprintf("%s%s%v", q.Name, q.Operator, q.Value)
}
