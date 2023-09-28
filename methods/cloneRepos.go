package methods

import (
	"CC1/types"
	"fmt"

	"gopkg.in/src-d/go-git.v4"
)

func CloneRepos(username string, data []types.ResponsData) error {
	localFolderPath := fmt.Sprintf("github/%s", username)

	for _, repo := range data {
		repoName := repo.Name
		cloneURL := repo.CloneURL

		repoFolderPath := fmt.Sprintf("%s/%s", localFolderPath, repoName)

		_, err := git.PlainClone(repoFolderPath, false, &git.CloneOptions{
			URL: cloneURL,
		})

		if err != nil {
			return err
		}
	}

	return nil
}
