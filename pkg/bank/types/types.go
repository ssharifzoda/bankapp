package types

type Money int64

type Currency string

const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

type Category string

type PAN string

type Card struct {
	ID         int
	PAN        PAN
	Balance    Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
	MinBalance Money
}
type Payment struct {
	ID       int
	Amount   Money
	Category Category
}
type PaymentSource struct {
	Type    string
	Number  string
	Balance Money
}
