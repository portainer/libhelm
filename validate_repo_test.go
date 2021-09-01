package libhelm

import (
	"testing"

	"github.com/portainer/libhelm/libhelmtest"
	"github.com/stretchr/testify/assert"
)

type validationTestCase struct {
	name    string
	url     string
	invalid bool
}

func validateUrl(test validationTestCase, t *testing.T) {
	is := assert.New(t)
	t.Run(test.name, func(t *testing.T) {
		t.Parallel()
		err := ValidateHelmRepositoryURL(test.url)
		if test.invalid {
			is.Errorf(err, "error expected: %s", test.url)
		} else {
			is.NoError(err, "no error expected: %s", test.url)
		}
	})
}

func Test_ValidateHelmRepositoryURL(t *testing.T) {
	libhelmtest.EnsureIntegrationTest(t)

	tests := []validationTestCase{
		{"blank", "", true},
		{"slashes", "//", true},
		{"slash", "/", true},
		{"invalid scheme", "garbage://a.b.c", true},
		{"invalid domain", "https://invaliddomain/", true},
		{"not helm repo", "http://google.com", true},
		{"not valid repo with trailing slash", "http://google.com/", true},
		{"not valid repo with trailing slashes", "http://google.com////", true},
		{"bitnami helm repo", "https://charts.bitnami.com/bitnami/", false},
		{"gitlap helm repo", "https://charts.gitlab.io/", false},
		{"portainer helm repo", "https://portainer.github.io/k8s/", false},
		{"bitnami helm repo", "https://helm.elastic.co/", false},
	}

	for _, test := range tests {
		validateUrl(test, t)
	}
}
