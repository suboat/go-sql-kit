package gosql

const (
	LimitKeyLimit string = "%l"
	LimitKeySkip         = "%s"
	LimitKeyPage         = "%p"
)

func IsLimitKey(str string) bool {
	switch str {
	case LimitKeyLimit, LimitKeySkip, LimitKeyPage:
		return true
	}
	return false
}

type ILimit interface {
	IsLimited() bool
}

type LimitRoot struct {
	Values []ILimit
}

func (l *LimitRoot) IsLimited() bool {
	return false
}

type LimitValue struct {
	Key   string
	Value int
}

func (l *LimitValue) IsLimited() bool {
	switch l.Key {
	case LimitKeyLimit:
		return l.Value > 0
	case LimitKeySkip:
		return true
	case LimitKeyPage:
		return l.Value >= 0
	}
	return false
}

type LimitValues struct {
	Limit int
	Skip  int
	Page  int
}

func (l *LimitValues) IsLimited() bool {
	return l.Limit > 0 && (l.Skip+l.Limit*l.Page) >= 0
}

type IRuleLimit interface {
	SetMaxLimit(int) IRuleLimit
	GetLimit(int) int
}

type RuleLimit struct {
	maxLimit int
}

func (l *RuleLimit) SetMaxLimit(lmt int) IRuleLimit {
	if lmt >= 0 {
		l.maxLimit = lmt
	} else {
		l.maxLimit = 0
	}
	return l
}

func (l *RuleLimit) GetLimit(lmt int) int {
	if lmt <= 0 {
		return 0
	} else if l.maxLimit <= 0 {
	} else if lmt > l.maxLimit {
		return l.maxLimit
	}
	return lmt
}
