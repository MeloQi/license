package crypto

import (
	"testing"
	"math/rand"
	"time"
)

func TestSm4Pading(t *testing.T) {
	t.Log("\n")
	data := "0123456789abcdefgh"
	t.Logf("data: %x", []byte(data))
	encPadingData := NewSM4Utils().encPading([]byte(data))
	t.Logf("encPadingData: %x", encPadingData)
	decPadingData := NewSM4Utils().decPading(encPadingData)
	t.Logf("decPadingData: %x", decPadingData)
}
const numberArray string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIZKLMNOPQRSTUVWXYZ0123456789"
func GetRandStr(length int) string {
	var str []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		num := r.Int() % len(numberArray)
		str = append(str, numberArray[num])
	}
	return string(str)
}
func TestSm4encdec(t *testing.T) {
	t.Log("\n")
	data := "爱迪生水电费全额wef很温柔合同儿童画委托函kqwow我我二哥 二哥eeoig" +
		"安慰覅机器噢诶人安慰热欧冠IQ噢 而我挺好玩热狗" +
		"加热管are我给枸杞 玩儿个问题日过期我二个人沟通和维护玩儿二哥日韩国认" +
		"为该奥尔我给人工清热阿二哥二额地方感动死了喀" +
		"纳斯的的热给哦噶维尔个人股发清热违法而发热发  缺乏的妇女IE你" +
		"论文的分工仍按时发起威风威风请问 窝头会馆沃尔提货人沙企鹅我热饭龙的发送到你哦啊诶让我" +
		"给狗水电费 打发我  我玩儿我让他给我"
	t.Logf("data: %s", data)
	key := GetRandStr(16)
	t.Log("key:", key)
	encData, err := NewSM4Utils().GetEncStr(&data, &key)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	t.Logf("encData: %s", encData)
	decData, err := NewSM4Utils().GetDecStr(&encData, &key)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	t.Logf("decData: %s", decData)
}

func TestSM4Dec(t *testing.T) {
	t.Log("\n")
	data := "eIs7l4hSWOUWxRrRheO7nlNxbZjRAWZ3f0IpJagL3n5eBMZs+7O4gHF4lnfMRSuN7um1VrDGpPszVtQQ/Mb/n3d0k29cm+CpQaq84kMogyRSdnBX/BIepaMkTA7lRjlRrtVNxfUnveBKta4d8W6WvA=="
	key := "6cEUghiR2pP6U151"
	if decData, err := NewSM4Utils().GetDecStr(&data, &key); err != nil {
		t.Log(err)
		t.Fail()
	} else {
		t.Log("decData； ", decData)
	}
}
