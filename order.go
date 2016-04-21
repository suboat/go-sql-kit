package gosql

const (
	OrderKey_asc  string = "+"
	OrderKey_desc        = "-"
)

type IOrder interface {
	IsASC() bool
	IsDESC() bool
	String() string
}

type OrderRoot struct {
	Value []IOrder
}

func (o *OrderRoot) IsASC() bool {
	return false
}

func (o *OrderRoot) IsDESC() bool {
	return false
}

func (o *OrderRoot) String() string {
	return ""
}

type OrderValue struct {
	ASC   bool
	Field string
}

func (o *OrderValue) IsASC() bool {
	return o.ASC
}

func (o *OrderValue) IsDESC() bool {
	return !o.ASC
}

func (o *OrderValue) String() string {
	return o.Field
}
