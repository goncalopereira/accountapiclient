package data

//Implementing Null Object Pattern, non nil result.
type NoOp struct {
	IOutput
}

type Deleted struct {
	NoOp
}
