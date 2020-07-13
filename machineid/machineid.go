package machineid

import "github.com/MeloQi/machineid"

func GetMachineid(appID string) (string, error) {
	return machineid.ProtectedID(appID)
}
