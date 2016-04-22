package gosql

const (
	LimitKey string = "%"
)

type ILimit interface {
	IsLimited() bool
}

type LimitRoot struct {
	Value ILimit
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

type IRuleLimit interface {
	SetDefaultLimit(int) IRuleLimit
	SetMaxLimit(int) IRuleLimit
	Limit(int) int
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

func (l *RuleLimit) Limit(lmt int) int {
	if lmt >= 1 {
		if lmt > l.maxLimit {
			return l.maxLimit
		}
		return lmt
	}
	return l.defaultLimit
}
