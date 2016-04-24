package gosql

const (
	OrderKey string = "%o"
)

func IsOrderKey(str string) bool {
	switch str {
	case OrderKey:
		return true
	}
	return false
}

const (
	OrderKeyASC  string = "+" // 正序
	OrderKeyDESC        = "-" // 反序
)

type IOrder interface {
	IsASC() bool
	IsDESC() bool
}

type OrderRoot struct {
	Values []IOrder
}

func (o *OrderRoot) IsASC() bool {
	return false
}

func (o *OrderRoot) IsDESC() bool {
	return false
}

type OrderValue struct {
	Key   string
	Field string
}

func (o *OrderValue) IsASC() bool {
	return o.Key == OrderKeyASC
}

func (o *OrderValue) IsDESC() bool {
	return !o.IsASC()
}
