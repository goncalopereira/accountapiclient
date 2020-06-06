package data

//interface for any return types like
//Account, Accounts and ErrorMessage
type IOutput interface {
	String() string
}
