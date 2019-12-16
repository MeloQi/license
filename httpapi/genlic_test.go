package httpapi

import (
	"testing"
	"encoding/json"
	"net/http"
	"strings"
	"io/ioutil"
	"github.com/MeloQi/license/license"
)

func TestGenLicHttpApi(t *testing.T) {

	HttpApi := GetGenLicHttpApiInst(":8081")
	HttpApi.Start()

	body, err := clientPost("topsci", "bdqi", "c1", "im", "2018-06-02 15:04:05", "88bf74f6-423a-4fa1-8251-330d5e01b791", "0123456789abcdef", t)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	t.Log("body:", body)
	if licInfo, err := license.LicCheck(body, "0123456789abcdef"); err != nil {
		t.Error(err)
		t.Fail()
		return
	} else {
		t.Log("LicCheck OK, license Info:", licInfo)
	}
}

func clientPost(org, applicant, user, appname, exp, machineid, key string, t *testing.T) (string, error) {
	type LicInfo struct {
		Org       string `json:"org"`
		Applicant string `json:"applicant"`
		User      string `json:"user"`
		Appname   string `json:"appname"`
		Exp       string `json:"exp"`
		Machineid string `json:"machineid"`
	}
	req := &LicInfo{Org: org, Applicant: applicant, User: user, Appname: appname, Exp: exp, Machineid: machineid}
	reqStr, err := json.Marshal(req)
	if err != nil {
		t.Error(err)
		t.Fail()
		return "", err
	}
	url := "http://127.0.0.1:8081/lic/getlic/" + key
	resp, err := http.Post(url, "application/json", strings.NewReader(string(reqStr)))
	if err != nil {
		t.Error(err)
		t.Fail()
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		t.Fail()
		return "", err
	}
	t.Log(string(body))
	return string(body), nil
}
