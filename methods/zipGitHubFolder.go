package methods

import (
	"fmt"
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func ZipGitHubFolder(username string) error {
	githubFolder := fmt.Sprintf("github/%s", username)

	zipFileName := fmt.Sprintf("github/%s/%s.zip", username, username)

	zipFile, err := os.Create(zipFileName)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	err = filepath.Walk(githubFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if path == githubFolder {
			return nil
		}

		relPath := strings.TrimPrefix(path, githubFolder+string(os.PathSeparator))

		zipHeader, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		zipHeader.Name = relPath

		writer, err := zipWriter.CreateHeader(zipHeader)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			_, err = io.Copy(writer, file)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	fmt.Printf("Dossier GitHub archivé dans %s\n", zipFileName)

	newZipFileName := fmt.Sprintf("github/%s.zip", username)
	err = os.Rename(zipFileName, newZipFileName)
	if err != nil {
		return err
	}

	err = removeAllContents(githubFolder)
	if err != nil {
		return err
	}

	return nil
}