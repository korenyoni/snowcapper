package extract

import (
	"errors"
	"fmt"
	"github.com/mholt/archiver"
	"github.com/yonkornilov/snowcapper/config"
	"io"
	"os"
	"strings"
)

func Run(b config.Binary, downloadPath string) (binaryPath string, err error) {
	binaryPath = b.GetBinaryPath()

	fmt.Printf("Extracting %s from %s\n", b.Format, downloadPath)
	err, extractedPath := extract(b.Format, downloadPath)
	if err != nil {
		return "", err
	}

	fmt.Printf("Copying %s to %s\n", getExtractedBinaryPath(b, extractedPath), binaryPath)
	err = copyToTarget(getExtractedBinaryPath(b, extractedPath), binaryPath)
	if err != nil {
		return "", err
	}

	fmt.Printf("Removing %s\n", extractedPath)
	err = os.RemoveAll(extractedPath)
	if err != nil {
		return "", err
	}

	fmt.Printf("Removing %s\n", downloadPath)
	err = os.RemoveAll(b.Src)
	if err != nil {
		return "", err
	}

	return binaryPath, nil
}

func extract(archiveType string, src string) (error, string) {
	extractedPath := getExtractedPath(archiveType, src)
	var err error
	switch archiveType {
	case "zip":
		err = archiver.Zip.Open(src, extractedPath)
	case "tar":
		err = archiver.Tar.Open(src, extractedPath)
	case "tar.gz":
		err = archiver.TarGz.Open(src, extractedPath)
	case "tar.bz2":
		err = archiver.TarBz2.Open(src, extractedPath)
	case "tar.xz":
		err = archiver.TarXZ.Open(src, extractedPath)
	case "tar.lz4":
		err = archiver.TarLz4.Open(src, extractedPath)
	case "tar.sz":
		err = archiver.TarSz.Open(src, extractedPath)
	case "rar":
		err = archiver.Rar.Open(src, extractedPath)
	default:
		err = errors.New(fmt.Sprintf("Error: 'Type' must be one of: %s", archiver.SupportedFormats))
	}

	if err != nil {
		return err, ""
	}

	return nil, extractedPath
}

func copyToTarget(src string, target string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(target)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	err = out.Close()
	if err != nil {
		return err
	}

	return nil
}

func getExtractedPath(archiveType string, src string) string {
	extractedPath := strings.Replace(src, "."+archiveType, "", -1)
	return extractedPath
}

func getExtractedBinaryPath(b config.Binary, extractedPath string) string {
	return extractedPath + "/" + b.Name
}
