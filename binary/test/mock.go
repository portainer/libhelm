package test

import (
	"time"

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

type helmMockPackageManager struct{}

// NewMockHelmBinaryPackageManager initializes a new HelmPackageManager service (a mock instance)
func NewMockHelmBinaryPackageManager(binaryPath string) libhelm.HelmPackageManager {
	return &helmMockPackageManager{}
}

type mockChart struct {
	Name       string    `json:"name"`
	Namespace  string    `json:"namespace"`
	Updated    time.Time `json:"updated"`
	Status     string    `json:"status"`
	Chart      string    `json:"chart"`
	AppVersion string    `json:"app_version"`
}

var mock_charts = []*mockChart{}

func newMockChart(installOpts options.InstallOptions) *mockChart {
	return &mockChart{
		Name:       installOpts.Name,
		Namespace:  installOpts.Namespace,
		Updated:    time.Now(),
		Status:     "deployed",
		Chart:      installOpts.Chart + "9.4.2",
		AppVersion: "1.21.1",
	}
}

func mockChartAsRelease(mc *mockChart) *release.Release {
	return &release.Release{
		Name:      mc.Name,
		Namespace: mc.Namespace,
	}
}

func (hpm *helmMockPackageManager) Install(installOpts options.InstallOptions) (*release.Release, error) {

	release := newMockChart(installOpts)

	// Enforce only one chart with the same name per namespace
	for i, rel := range mock_charts {
		if rel.Name == installOpts.Name && rel.Namespace == installOpts.Namespace {
			mock_charts[i] = release
			return mockChartAsRelease(release), nil
		}
	}

	mock_charts = append(mock_charts, release)
	return mockChartAsRelease(release), nil
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
