package crypto

import (
	"testing"
	"errors"
)

func TestAll(t *testing.T) {
	// 公钥加密私钥解密
	if err := applyPubEPriD(); err != nil {
		t.Error(err)
		t.Fail()
	} else {
		t.Log("公钥加密私钥解密 OK")
	}
	// 公钥解密私钥加密
	if err := applyPriEPubD(); err != nil {
		t.Error(err)
		t.Fail()
	} else {
		t.Log("公钥解密私钥加密 OK")
	}
}

// 初始化设置公钥和私钥
func setKey(r *RSASecurity) error {
	pubkey, prikey, err := GenRsaKey(1024)
	if err != nil {
		return err
	}
	if err := r.SetPublicKey(pubkey); err != nil {
		return errors.New(`set public key :` + err.Error())
	}
	if err := r.SetPrivateKey(prikey); err != nil {
		return errors.New(`set private key :` + err.Error())
	}
	return nil
}

// 公钥加密私钥解密
func applyPubEPriD() error {
	r := &RSASecurity{}
	if err := setKey(r); err != nil {
		return err
	}
	pubenctypt, err := r.PubKeyENCTYPT([]byte(`hello world`))
	if err != nil {
		return err
	}

	pridecrypt, err := r.PriKeyDECRYPT(pubenctypt)
	if err != nil {
		return err
	}
	if string(pridecrypt) != `hello world` {
		return errors.New(`解密失败`)
	}
	return nil
}

// 公钥解密私钥加密
func applyPriEPubD() error {
	r := &RSASecurity{}
	if err := setKey(r); err != nil {
		return err
	}
	prienctypt, err := r.PriKeyENCTYPT([]byte(`hello world`))
	if err != nil {
		return err
	}

	pubdecrypt, err := r.PubKeyDECRYPT(prienctypt)
	if err != nil {
		return err
	}
	if string(pubdecrypt) != `hello world` {
		return errors.New(`解密失败`)
	}
	return nil
}

func TestRsaGen(t *testing.T) {
	pubkey, prikey, err := GenRsaKey(1024)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	t.Log("pubkeyPem:", pubkey)
	t.Log("pubkeyStr:", GetStrByPem(pubkey))
	t.Log("pubkeyStrToPem:", GetPemByStr(GetStrByPem(pubkey)))
	t.Log("prikeyPem:", prikey)
	t.Log("prikeyStr:", GetStrByPem(prikey))
	t.Log("prikeyStrToPem:", GetPemByStr(GetStrByPem(prikey)))
}
