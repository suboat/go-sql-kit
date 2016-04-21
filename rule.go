package gosql

type IRuleMapping interface {
	Allow(string) IRuleMapping
	Disallow(string) IRuleMapping
	IsAllowed(string) bool
	SetMapping(string, string) IRuleMapping
	GetMapping(string) string
}

type RuleMapping struct {
	allowKey map[string]bool
	mappings map[string]string
}

func (r *RuleMapping) rule() *RuleMapping {
	if r.allowKey == nil {
		r.allowKey = make(map[string]bool)
	}
	if r.mappings == nil {
		r.mappings = make(map[string]string)
	}
	return r
}

func (r *RuleMapping) Allow(key string) IRuleMapping {
	if len(key) != 0 {
		r.rule().allowKey[key] = true
	}
	return r
}

func (r *RuleMapping) Disallow(key string) IRuleMapping {
	if len(key) != 0 {
		r.rule().allowKey[key] = false
	}
	return r
}

func (r *RuleMapping) IsAllowed(key string) bool {
	if len(key) == 0 {
		return false
	} else if allow, ok := r.rule().allowKey[key]; ok {
		return allow
	}
	return true
}

func (r *RuleMapping) SetMapping(value, mapping string) IRuleMapping {
	if len(value) != 0 {
		r.rule().mappings[value] = mapping
	}
	return r
}

func (r *RuleMapping) GetMapping(value string) string {
	if len(value) == 0 {
	} else if mapping, ok := r.rule().mappings[value]; ok {
		return mapping
	}
	return value
}
