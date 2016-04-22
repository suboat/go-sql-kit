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

type LimitValueSet struct {
	Limit int
	Skip  int
	Page  int
}

func (l *LimitValueSet) IsLimited() bool {
	return l.Limit > 0 && (l.Skip+l.Limit*l.Page) >= 0
}

type IRuleLimit interface {
	SetDefaultLimit(int) IRuleLimit
	SetMaxLimit(int) IRuleLimit
	GetLimit(int) int
}

type RuleLimit struct {
	defaultLimit int
	maxLimit     int
}

func (l *RuleLimit) SetDefaultLimit(lmt int) IRuleLimit {
	if lmt >= 0 {
		if lmt > l.maxLimit {
			l.SetMaxLimit(lmt)
		}
		l.defaultLimit = lmt
	}
	return l
}

func (l *RuleLimit) SetMaxLimit(lmt int) IRuleLimit {
	if lmt >= 0 {
		if lmt < l.defaultLimit {
			l.SetDefaultLimit(lmt)
		}
		l.maxLimit = lmt
	}
	return l
}

func (l *RuleLimit) GetLimit(lmt int) int {
	if lmt >= 1 {
		if lmt > l.maxLimit {
			return l.maxLimit
		}
		return lmt
	}
	return l.defaultLimit
}
