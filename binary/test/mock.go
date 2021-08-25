package test

import (
	"github.com/portainer/libhelm"
	"github.com/portainer/libhelm/options"
	"github.com/portainer/libhelm/release"
)

const (
	MockDataIndex  = "mock-index"
	MockDataChart  = "mock-chart"
	MockDataReadme = "mock-readme"
	MockDataValues = "mock-values"
)

// helmMockPackageManager is a test package for helm related http handler testing
// Note: this package currently uses a slice in a way that is not thread safe.
// Do not use this package for concurrent tests.
type helmMockPackageManager struct{}

// NewMockHelmBinaryPackageManager initializes a new HelmPackageManager service (a mock instance)
func NewMockHelmBinaryPackageManager(binaryPath string) libhelm.HelmPackageManager {
	return &helmMockPackageManager{}
}

var mockCharts = []release.ReleaseElement{}

func newMockReleaseElement(installOpts options.InstallOptions) *release.ReleaseElement {
	return &release.ReleaseElement{
		Name:       installOpts.Name,
		Namespace:  installOpts.Namespace,
		Updated:    "date/time",
		Status:     "deployed",
		Chart:      installOpts.Chart,
		AppVersion: "1.2.3",
	}
}

func newMockRelease(re *release.ReleaseElement) *release.Release {
	return &release.Release{
		Name:      re.Name,
		Namespace: re.Namespace,
	}
}

// Install a helm chart (not thread safe)
func (hpm *helmMockPackageManager) Install(installOpts options.InstallOptions) (*release.Release, error) {

	releaseElement := newMockReleaseElement(installOpts)

	// Enforce only one chart with the same name per namespace
	for i, rel := range mockCharts {
		if rel.Name == installOpts.Name && rel.Namespace == installOpts.Namespace {
			mockCharts[i] = *releaseElement
			return newMockRelease(releaseElement), nil
		}
	}

	mockCharts = append(mockCharts, *releaseElement)
	return newMockRelease(releaseElement), nil
}

// Show values/readme/chart etc
func (hpm *helmMockPackageManager) Show(showOpts options.ShowOptions) ([]byte, error) {
	switch showOpts.OutputFormat {
	case options.ShowChart:
		return []byte(MockDataChart), nil
	case options.ShowReadme:
		return []byte(MockDataReadme), nil
	case options.ShowValues:
		return []byte(MockDataValues), nil
	}
	return nil, nil
}

// Uninstall a helm chart (not thread safe)
func (hpm *helmMockPackageManager) Uninstall(uninstallOpts options.UninstallOptions) error {
	for i, rel := range mockCharts {
		if rel.Name == uninstallOpts.Name && rel.Namespace == uninstallOpts.Namespace {
			mockCharts = append(mockCharts[:i], mockCharts[i+1:]...)
		}
	}
	return nil
}

// List a helm chart (not thread safe)
func (hpm *helmMockPackageManager) List(listOpts options.ListOptions) ([]release.ReleaseElement, error) {
	return mockCharts, nil
}
