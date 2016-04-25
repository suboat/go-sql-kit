package gosql

type IParser interface {
	ParseJSONString(string) error
	Parse(map[string]interface{}) error
}
