package meta

import "testing"

var mq *Meta

func TestNewMeta(t *testing.T) {
	example := `[{"%and":{"%eq":{"key1":"A12","key2":"B23"},"%ne":{"key3":"C34","key4":"D45"}}},` +
		`{"%o":["+key1","-key2"]},` +
		`{"%l":5,"%s":12,"%p":1}]`
	mq = NewMeta(example)
}

func TestMeta_ParseJSONValue(t *testing.T) {
	for i := 0; i < 100; i++ {
		if _, err := mq.ParseJSONValue(); err != nil {
			t.Fatal(err)
		}
	}
}

func TestMeta_GetData(t *testing.T) {
	for i := 0; i < 100; i++ {
		if v, ok := mq.ParseData("key1"); ok {
			if str, ok := v.(string); ok && str == "A12" {
				return
			}
			t.Errorf("%#v", v)
		}
	}
}

func TestMeta_UniqueID(t *testing.T) {
	for i := 0; i < 100; i++ {
		if v, ok := mq.UniqueID(); !ok || len(v) == 0 {
			t.Fatal("Error unique id")
		}
	}
}
