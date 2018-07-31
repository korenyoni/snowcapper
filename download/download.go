package download

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"hash"
	"errors"
	"encoding/hex"

	"crypto/sha512"
	"crypto/sha256"
	"crypto/sha1"
	"crypto/md5"

	"github.com/yonkornilov/snowcapper/config"
	"github.com/yonkornilov/snowcapper/context"
)

func Run(c *context.Context, downloadable interface{}) (downloadPath string, err error) {
	switch t := downloadable.(type) {
	case config.Binary:
		return downloadBinary(c, downloadable.(config.Binary))
	case config.Extend:
		return downloadExtend(c, downloadable.(config.Extend))
	default:
		return "", errors.New(fmt.Sprintf("not a valid Downloadable type: %s", t))
	}
}

func downloadBinary(c *context.Context, b config.Binary) (downloadPath string, err error) {
	target := b.GetDownloadPath()
	if c.IsDryRun {
		fmt.Printf("DRY-RUN: Downloading %s from %s ...\n", b.Name, b.Src)
		fmt.Printf("DRY-RUN: Successfully downloaded %s to %s\n", b.Name, target)
		return target, nil
	}
	fmt.Printf("Downloading %s from %s ...\n", b.Name, b.Src)
	out, err := os.Create(target)
	if err != nil {
		return "", err
	}
	defer out.Close()

	resp, err := http.Get(b.Src)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	hashExists, err := checkHashIfExists(respBodyBytes, b.SrcHash)
	if err != nil {
		return "", err
	}

	_, err = out.Write(respBodyBytes) 
	if err != nil {
		return "", err
	}

	if hashExists {
		fmt.Printf("Successfully downloaded %s to %s with hashsum %s\n", b.Name, target, b.SrcHash)
	} else {
		fmt.Printf("Successfully downloaded %s to %s\n", b.Name, target)
	}
	return target, nil
}

func downloadExtend(c *context.Context, e config.Extend) (downloadPath string, err error) {
	return "", nil
}

func checkHashIfExists(body []byte, hash string) (exists bool, err error) {
	hashTypeCode := config.GetHashType(hash)
	if hashTypeCode == 0 {
		return false, nil
	}
	hasher, err := getHasher(hashTypeCode) 
	if err != nil {
		return false, err
	}
	hasher.Write(body)
	fileHashSumHex := hex.EncodeToString(hasher.Sum(nil))
	if fileHashSumHex != hash {
		return false, errors.New("file does not match hashsum")
	}
	return true, nil
}

func getHasher(hashTypeCode int) (hash.Hash, error) {
	switch hashTypeCode {
	case config.Sha512:
		return sha512.New(), nil
	case config.Sha384:
		return sha512.New384(), nil
	case config.Sha256:
		return sha256.New(), nil
	case config.Sha224:
		return sha256.New224(), nil
	case config.Sha1:
		return sha1.New(), nil
	case config.Md5:
		return md5.New(), nil
	default:
		return nil, errors.New(fmt.Sprintf("Invalid hash type code: %d", hashTypeCode))
	}
}
