package gosql

type IRuleMapping interface {
	Allow(...string) IRuleMapping
	Disallow(...string) IRuleMapping
	IsAllowed(string) bool
	SetMapping(string, string) IRuleMapping
	GetMapping(string) string
	SetMappingFunc(string, RuleMappingFunc) IRuleMapping
	GetMappingFunc(string) (RuleMappingFunc, bool)
	SetMappingValue(string, string, interface{}) IRuleMapping
	SetRuleMappingResult(string, RuleMappingResult) IRuleMapping
	GetRuleMappingResult(string) (RuleMappingResult, bool)
}

type RuleMappingFunc func(string, interface{}) (string, interface{}, bool)
type RuleMappingResult func(string, interface{}, ...string) (interface{}, bool)

type RuleMapping struct {
	allowKey       map[string]bool
	mappings       map[string]string
	mappingFuncs   map[string]RuleMappingFunc
	mappingResults map[string]RuleMappingResult
}

func (r *RuleMapping) rule() *RuleMapping {
	if r.allowKey == nil {
		r.allowKey = make(map[string]bool)
	}
	if r.mappings == nil {
		r.mappings = make(map[string]string)
	}
	if r.mappingFuncs == nil {
		r.mappingFuncs = make(map[string]RuleMappingFunc)
	}
	if r.mappingResults == nil {
		r.mappingResults = make(map[string]RuleMappingResult)
	}
	return r
}

func (r *RuleMapping) Allow(keys ...string) IRuleMapping {
	if len(keys) != 0 {
		for _, key := range keys {
			if len(key) == 0 {
				continue
			}
			r.rule().allowKey[key] = true
		}
	}
	return r
}

func (r *RuleMapping) Disallow(keys ...string) IRuleMapping {
	if len(keys) != 0 {
		for _, key := range keys {
			if len(key) == 0 {
				continue
			}
			r.rule().allowKey[key] = false
		}
	}
	return r
}

func (r *RuleMapping) IsAllowed(key string) bool {
	if len(key) == 0 {
		return false
	} else if allow, ok := r.rule().allowKey[key]; ok {
		return allow
	}
	return false
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

func (r *RuleMapping) SetMappingFunc(value string, f RuleMappingFunc) IRuleMapping {
	if len(value) != 0 && f != nil {
		r.rule().mappingFuncs[value] = f
	}
	return r
}

func (r *RuleMapping) GetMappingFunc(value string) (RuleMappingFunc, bool) {
	if len(value) == 0 {
	} else if f, ok := r.rule().mappingFuncs[value]; ok && f != nil {
		return f, true
	}
	return nil, false
}

func (r *RuleMapping) SetMappingValue(value string, mapping string, v interface{}) IRuleMapping {
	return r.SetMappingFunc(value, func(string, interface{}) (string, interface{}, bool) {
		return mapping, v, true
	})
}

func (r *RuleMapping) SetRuleMappingResult(value string, f RuleMappingResult) IRuleMapping {
	if len(value) != 0 && f != nil {
		r.rule().mappingResults[value] = f
	}
	return r
}

func (r *RuleMapping) GetRuleMappingResult(value string) (RuleMappingResult, bool) {
	if len(value) == 0 {
	} else if f, ok := r.rule().mappingResults[value]; ok && f != nil {
		return f, true
	}
	return nil, false
}
