package gosql

const (
	OrderKeyASC  string = "%asc"  // 正序
	OrderKeyDESC        = "%desc" // 反序
)

func IsOrderKey(str string) bool {
	switch str {
	case OrderKeyASC, OrderKeyDESC:
		return true
	}
	return false
}

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
