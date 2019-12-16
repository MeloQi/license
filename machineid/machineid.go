package machineid

import "github.com/denisbrodbeck/machineid"

func GetMachineid(appID string) (string, error) {
	return machineid.ProtectedID(appID)
}
