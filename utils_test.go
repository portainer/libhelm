package libhelm

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// To run this test
// export INTEGRATION_TEST=1
// go test -v

func Test_ValidateHelmRepositoryURL(t *testing.T) {
	ensureIntegrationTest(t)
	is := assert.New(t)

	tests := []struct {
		name    string
		url     string
		invalid bool
	}{
		{"blank", "", true},
		{"slashes", "//", true},
		{"invalid scheme", "garbage://a.b.c", true},
		{"invalid domain", "https://invaliddomain/", true},
		{"not helm repo", "http://google.com", true},
		{"not valid repo with trailing slash", "http://google.com/", true},
		{"bitnami helm repo", "https://charts.bitnami.com/bitnami", false},
		{"gitlap helm repo", "https://charts.gitlab.io/", false},
		{"portainer helm repo", "https://portainer.github.io/k8s/", false},
		{"bitnami helm repo", "https://helm.elastic.co/", false},
	}

	for _, test := range tests {
		// Copy these range vars which are pass by reference.
		test := test
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
}

func ensureIntegrationTest(t *testing.T) {
	if _, ok := os.LookupEnv("INTEGRATION_TEST"); !ok {
		t.Skip("skip an integration test")
	}
}
