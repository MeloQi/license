package license

import (
	"encoding/json"
	"github.com/MeloQi/license/crypto"
	"errors"
	"encoding/base64"
	"github.com/MeloQi/license/serialize"
	"time"
	"strings"
	"github.com/MeloQi/license/utils"
)

type LicInfo struct {
	Type       string `json:"type"`
	Org       string `json:"org"`
	Applicant string `json:"applicant"`
	User      string `json:"user"`
	Appname   string `json:"appname"`
	Exp       string `json:"exp"`
	Machineid string `json:"machineid"`
	Time      string `json:"time"`
}

func GenLic(licInfo *LicInfo, key string) (string, error) {
	if len(key) != 16 {
		return "", errors.New("key length err")
	}
	licInfo.Time = time.Now().Format("2006-01-02 15:04:05")
	if _, err := time.Parse("2006-01-02 15:04:05", licInfo.Exp); err != nil {
		return "", err
	}

	//1. 明文license信息字段
	licInfoJsonArry, err := json.Marshal(licInfo)
	if err != nil {
		return "", err
	}
	licInfoJsonStr := base64.StdEncoding.EncodeToString(licInfoJsonArry)

	//2. 密文
	licKey := utils.GetRandomString(16)
	licInfoScrtStr := string(licInfoJsonArry);
	if licInfoScrtStr, err = crypto.NewSM4Utils().GetEncStr(&licInfoScrtStr, &licKey); err != nil {
		return "", err
	}

	//3. 密文key加密
	r := &crypto.RSASecurity{}
	pubkey, prikey, err := crypto.GenRsaKey(1024)
	if err != nil {
		return "", err
	}
	if err := r.SetPublicKey(pubkey); err != nil {
		return "", errors.New(`set public key :` + err.Error())
	}
	if err := r.SetPrivateKey(prikey); err != nil {
		return "", errors.New(`set private key :` + err.Error())
	}
	if sc, err := r.PriKeyENCTYPT([]byte(licKey)); err != nil {
		return "", err
	} else {
		licKey = base64.StdEncoding.EncodeToString(sc)
	}

	//4. 公钥
	pubkeyStr := crypto.GetStrByPem(pubkey)
	if pubkeyStr, err = serialize.Serialize(pubkeyStr); err != nil {
		return "", err
	}
	if pubkeyStr, err = crypto.NewSM4Utils().GetEncStr(&pubkeyStr, &key); err != nil {
		return "", err
	}

	return licInfoJsonStr + "." + licInfoScrtStr + "." + pubkeyStr + "." + licKey, nil

}

func LicCheck(lic, key string) (string, error) {
	var err error
	licArry := strings.Split(lic, ".")
	if len(licArry) != 4 {
		return "", errors.New("解析错误")
	}
	licInfoStr := licArry[0]
	licInfoScrtStr := licArry[1]
	pubkeyStr := licArry[2]
	licKeyStr := licArry[3]

	//1. 解密公钥pubkey
	if pubkeyStr, err = crypto.NewSM4Utils().GetDecStr(&pubkeyStr, &key); err != nil {
		return "", errors.New("解析错误")
	}
	if pubkeyStr, err = serialize.Deserialize(pubkeyStr); err != nil {
		return "", errors.New("解析错误")
	}

	// 2. 用公钥解密密文key
	r := &crypto.RSASecurity{}
	r.SetPublicKey(crypto.GetPemByStr(pubkeyStr))
	if  k, err := base64.StdEncoding.DecodeString(licKeyStr);err != nil {
		return "", errors.New("解析错误")
	}else {
		licKeyStr = string(k)
	}
	if k, err := r.PubKeyDECRYPT([]byte(licKeyStr)); err != nil {
		return "", errors.New("解析错误")
	}else {
		licKeyStr = string(k)
	}

	//3. 用lickey解密密文
	if licInfoScrtStr, err = crypto.NewSM4Utils().GetDecStr(&licInfoScrtStr, &licKeyStr); err != nil {
		return "", errors.New("解析错误")
	}

	//4. 明文信息
	licInfoArry, err := base64.StdEncoding.DecodeString(licInfoStr)
	if err != nil {
		return "", errors.New("解析错误")
	}

	//5. 明文与解密后密文比较
	if string(licInfoArry) != licInfoScrtStr {
		return "", errors.New("解析错误")
	}

	return licInfoScrtStr, nil
}
