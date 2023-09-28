package methods

import (
	"CC1/types"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-resty/resty/v2"
)

func GetRepos(numRepos int, ch chan<- []types.ResponsData, errCh chan<- error) {
	defer close(ch)

	githubUsername, githubUsernameSet := os.LookupEnv("GITHUB_USERNAME")
	if !githubUsernameSet {
		errCh <- fmt.Errorf("GITHUB_USERNAME not set")
		return
	}

	githubToken, githubTokenSet := os.LookupEnv("GITHUB_TOKEN")
	if !githubTokenSet {
		errCh <- fmt.Errorf("GITHUB_TOKEN not set")
		return
	}

	apiURL := fmt.Sprintf("https://api.github.com/users/%s/repos?sort=updated&per_page=%d", githubUsername, numRepos)

	if !githubTokenSet {
		errCh <- fmt.Errorf("GITHUB_TOKEN not set")
		return
	}

	var response []types.ResponsData
	restyClient := resty.New()
	resp, err := restyClient.R().
		EnableTrace().
		SetAuthToken(githubToken).
		Get(apiURL)

	if err != nil {
		errCh <- err
		return
	}

	if resp.StatusCode() != http.StatusOK {
		errCh <- fmt.Errorf("Error > %s", resp.Status())
		return
	}

	if err := json.Unmarshal(resp.Body(), &response); err != nil {
		errCh <- err
		return
	}

	ch <- response
}
