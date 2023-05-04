package parser

type ProgramNode struct {
	Elements []interface{}
}

type DiiaNode struct {
	Name       string
	Params     []ParamNode
	ReturnType string
	Body       []interface{}
}

type NumberNode struct {
	Value int
}

type StringNode struct {
	Value string
}

type IdentifierNode struct {
	Name string
}

type ChainNode struct {
	Left  interface{}
	Right interface{}
}

type CallNode struct {
	Value interface{}
	Args  []interface{}
}

type ArithmeticNode struct {
	Left  interface{}
	Right interface{}
	Op    string
}

type AssignNode struct {
	Left  IdentifierNode
	Type  string
	Value interface{}
}

type ParamNode struct {
	Name string
	Type string
}

type BodyNode struct {
	Elements []interface{}
}
