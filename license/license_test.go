package license

import (
	"testing"
)

func TestGenLic(t *testing.T) {
	key := `0123!@#$%%^*abcd`
	licInfo := &LicInfo{
		Type:      "trial",
		Org:       "topsic",
		Applicant: "bdqi",
		User:      "topsci",
		Appname:   "fims",
		Exp:       "2018-09-27 00:00:00",
		Machineid: "88bf74f6-423a-4fa1-8251-330d5e01b791",
	}

	lic, err := GenLic(licInfo, key)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	} else {
		t.Log("\n\nOK,LIC:\n", lic)
	}
	if licInfo, err := LicCheck(lic, key); err != nil {
		t.Error("解析错误")
		t.Fail()
		return
	} else {
		t.Log("LicCheck OK, license Info:", licInfo)
	}
}
