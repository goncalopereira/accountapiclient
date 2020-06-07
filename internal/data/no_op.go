package data

//NoOp implements Null Object Pattern, non nil result.
type NoOp struct {
	IOutput
}

//Deleted represents a valid deletion.
type Deleted struct {
	NoOp
}
