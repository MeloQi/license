package machineid

import "testing"

func TestGetMachineid(t *testing.T) {
	id, err := GetMachineid()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	t.Log("OK :", id)
}
