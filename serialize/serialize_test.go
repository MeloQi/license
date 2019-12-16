package serialize

import (
	"testing"
	"github.com/MeloQi/license/crypto"
)

func TestSerialize(t *testing.T) {
	str, _, err := crypto.GenRsaKey(1024)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	str = crypto.GetStrByPem(str)
	serializeStr, _ := Serialize(str)
	t.Log("序列化后：", serializeStr)
	deSerializeStr, _ := Deserialize(serializeStr)
	t.Log("反序列化后：", deSerializeStr)
	if deSerializeStr != str {
		t.Error("测试失败")
		t.Fail()
	} else {
		t.Log("OK")
	}
}
