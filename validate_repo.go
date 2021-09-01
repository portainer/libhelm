package libhelm

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"
)

const invalidChartRepo = "%q is not a valid chart repository or cannot be reached"

func ValidateHelmRepositoryURL(repoUrl string) error {
	repoUrl = strings.TrimSuffix(repoUrl, "/")

	if repoUrl == "" {
		return errors.New("URL is required")
	}

	url, err := url.ParseRequestURI(repoUrl)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid chart URL format: %s", repoUrl))
	}

	if !strings.EqualFold(url.Scheme, "http") && !strings.EqualFold(url.Scheme, "https") {
		return errors.New(fmt.Sprintf("invalid chart URL format: %s", repoUrl))
	}

	var client = &http.Client{
		Timeout: time.Second * 10,
	}
	res, err := client.Head(repoUrl + "/index.yaml")
	if err != nil {
		return errors.Wrapf(err, invalidChartRepo, repoUrl)
	}

	// Some servers return odd responses.  We need to check more than just content length to determine failure
	if res.ContentLength < 0 || res.StatusCode > 400 {
		return errors.Errorf(invalidChartRepo, repoUrl)
	}

	return nil
}
