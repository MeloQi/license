package machineid

import "github.com/denisbrodbeck/machineid"

func GetMachineid() (string, error) {
	return machineid.ID()
}
