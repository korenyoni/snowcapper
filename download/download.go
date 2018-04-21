package download

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Run(name string, path string, url string) error {
	fmt.Printf("Downloading %s from %s ...", name, url)
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("Successfully downloaded %s to %s", name, path)
	return nil
}
