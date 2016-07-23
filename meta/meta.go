package meta

import (
	"encoding/json"
	"fmt"
	"sort"
	"sync"
)

type Meta struct {
	mux    sync.RWMutex
	value  string
	array  []map[string]interface{}
	dict   map[string]interface{}
	unique string
}

func NewMeta(v string) *Meta {
	return &Meta{
		value: v,
	}
}

func (mq *Meta) Valid() bool {
	return len(mq.value) != 0
}

func (mq *Meta) GetValue() string {
	mq.mux.RLock()
	defer mq.mux.RUnlock()
	return mq.value
}

func (mq *Meta) IsParsed() bool {
	mq.mux.RLock()
	defer mq.mux.RUnlock()
	return mq.array != nil
}

func (mq *Meta) ParseJSONValue() ([]map[string]interface{}, error) {
	if mq.IsParsed() {
		return mq.array, nil
	}
	mq.mux.Lock()
	defer mq.mux.Unlock()
	var objs []interface{}
	if err := json.Unmarshal([]byte(mq.value), &objs); err != nil {
		return nil, err
	} else if mq.array = make([]map[string]interface{}, 0, len(objs)); len(objs) != 0 {
		for _, obj := range objs {
			if obj == nil {
			} else if m, ok := obj.(map[string]interface{}); ok {
				mq.array = append(mq.array, m)
			}
		}
	}
	return mq.array, nil
}

func (mq *Meta) parse(key string, value interface{}) {
	if value == nil {
	} else if m, ok := value.(map[string]interface{}); ok {
		arr := make([]string, 0, len(m))
		for k, _ := range m {
			arr = append(arr, k)
		}
		sort.Strings(arr)
		for _, k := range arr {
			mq.parse(k, m[k])
		}
	} else if len(key) != 0 {
		mq.dict[key] = value
		mq.unique += fmt.Sprintf("%v:%v|", key, value)
	}
}

func (mq *Meta) ParseData(key string) (interface{}, bool) {
	if mq.IsParsed() {
		mq.mux.RLock()
		if mq.dict == nil {
			mq.mux.RUnlock()
			mq.mux.Lock()
			mq.dict = make(map[string]interface{})
			for _, data := range mq.array {
				mq.parse("", data)
			}
			mq.mux.Unlock()
			mq.mux.RLock()
		}
		if len(key) != 0 && len(mq.dict) != 0 {
			v, ok := mq.dict[key]
			return v, ok
		}
		mq.mux.RUnlock()
	}
	return nil, false
}

func (mq *Meta) UniqueID() (string, bool) {
	mq.mux.RLock()
	if len(mq.unique) != 0 {
		mq.mux.RUnlock()
		return mq.unique, true
	}
	mq.mux.RUnlock()
	mq.ParseData("")
	return mq.unique, len(mq.unique) != 0
}
