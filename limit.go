package gosql

const (
	LimitKey string = "%"
)

type ILimit interface {
	IsLimited() bool
}

type LimitRoot struct {
	Value []ILimit
}

func (l *LimitRoot) IsLimited() bool {
	return false
}

type LimitValue struct {
	Limit int
	Skip  int
	Page  int
}

func (l *LimitValue) IsLimited() bool {
	return l.Limit > 0 && (l.Skip+l.Limit*l.Page) >= 0
}
