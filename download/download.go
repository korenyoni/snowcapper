package download

import (
	"fmt"
	"github.com/yonkornilov/snowcapper/config"
	"io"
	"net/http"
	"os"
)

func Run(b config.Binary, target string) error {
	fmt.Printf("Downloading %s from %s ...\n", b.Name, b.Src)
	out, err := os.Create(target)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(b.Src)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("Successfully downloaded %s to %s\n", b.Name, target)
	return nil
}
