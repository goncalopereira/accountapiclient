package data

//Implementing Null Object Pattern to be able to String() any result
//Also for Delete to have a non nil valid result
type NoOp struct {
	IOutput
}

func (o *NoOp) String() string {
	return ""
}

type NoContent struct {
	NoOp
}
