package gosql

type RuleMapping struct {
	allowKey   map[string]bool
	keyMapping map[string]string
}

func (r *RuleMapping) Allow(key string) {
	if len(key) == 0 {
		return
	} else if r.allowKey == nil {
		r.allowKey = make(map[string]bool)
	}
	r.allowKey[key] = true
}

func (r *RuleMapping) Disallow(key string) {
	if len(key) == 0 {
		return
	} else if r.allowKey == nil {
		r.allowKey = make(map[string]bool)
	}
	r.allowKey[key] = false
}

func (r *RuleMapping) IsAllowed(key string) bool {
	if len(key) == 0 {
		return false
	} else if r.allowKey == nil {
		return true
	} else if allow, ok := r.allowKey[key]; ok {
		return allow
	}
	return true
}

func (r *RuleMapping) SetMapping(key, mapping string) {
	if len(key) == 0 {
		return
	} else if r.keyMapping == nil {
		r.keyMapping = make(map[string]string)
	}
	r.keyMapping[key] = mapping
}

func (r *RuleMapping) GetMapping(key string) string {
	if len(key) == 0 {
		return key
	} else if r.keyMapping == nil {
		return key
	} else if mapping, ok := r.keyMapping[key]; ok {
		return mapping
	}
	return key
}
