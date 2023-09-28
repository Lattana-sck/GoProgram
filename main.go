package main 

import (
	"CC1/types"
	"log"
	"os"
	"CC1/methods"
)

func main() {
	numRepos := 1
	githubUsername, githubUsernameSet := os.LookupEnv("GITHUB_USERNAME")
	if !githubUsernameSet {
		log.Fatal("GITHUB_USERNAME not set")
	}

	dataCh := make(chan []types.ResponsData)
	errCh := make(chan error)

	go methods.GetRepos(numRepos, dataCh, errCh)

	select {
	case data := <-dataCh:
		err := methods.WriteCSV(githubUsername, data)
		if err != nil {
			log.Fatal(err)
		}

		err = methods.CloneRepos(githubUsername, data)
		if err != nil {
			log.Fatal(err)
		}

		err = methods.ZipGitHubFolder(githubUsername)
		if err != nil {
			log.Fatal(err)
		}

	case err := <-errCh:
		log.Fatal(err)
	}
}
