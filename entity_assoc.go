package mtx

type Association struct {
	One       *Entity
	Onename   string
	Other     *Entity
	Othername string
}
