package libhelm

import (
	"errors"

	"github.com/portainer/libhelm/binary"
)

// HelmConfig is a struct that holds the configuration for the Helm package manager
type HelmConfig struct {
	BinaryPath string `example:"/portainer/dist"`
}

var errBinaryPathNotSpecified = errors.New("binary path not specified")

// NewHelmManager returns a new instance of HelmPackageManager based on HelmConfig
func NewHelmManager(config HelmConfig) (HelmPackageManager, error) {
	if config.BinaryPath != "" {
		return binary.NewHelmBinaryPackageManager(config.BinaryPath), nil
	}
	return nil, errBinaryPathNotSpecified
}
