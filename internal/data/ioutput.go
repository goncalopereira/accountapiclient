package data

//IOutput is the interface for all command outputs
//Account, Accounts and ErrorMessage.
type IOutput interface {
	String() string
}
