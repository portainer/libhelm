package libhelm

import (
	"github.com/portainer/libhelm/options"
	"github.com/portainer/libhelm/release"
)

// HelmPackageManager represents a service that interfaces with Helm
type HelmPackageManager interface {
	Install(installOpts options.InstallOptions) (*release.Release, error)
	Show(showOpts options.ShowOptions) ([]byte, error)
}
