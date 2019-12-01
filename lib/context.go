package lib

type ContextKey string

func (c ContextKey) String() string {
	return "rbs.todos." + string(c)
}
