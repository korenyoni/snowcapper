package download

import (
	"fmt"
	"github.com/yonkornilov/snowcapper/config"
	"io"
	"net/http"
	"os"
)

func Run(p config.Package, target string) error {
	fmt.Printf("Downloading %s from %s ...\n", p.Name, p.Source)
	out, err := os.Create(target)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(p.Source)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("Successfully downloaded %s to %s\n", p.Name, target)
	return nil
}
