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

type DownloadableHolder struct {
	BinaryPointer *config.Binary
	ExtendPointer *config.Extend
	Downloadable interface{}
}

func Run(c *context.Context, downloadableHolder DownloadableHolder) (downloadPath string, err error) {
	var target string
	var name string
	var src string
	var src_hash string
	switch t := downloadableHolder.Downloadable.(type) {
	case config.Binary:
		binary := *downloadableHolder.BinaryPointer
		target = binary.GetDownloadPath()
		src = binary.Src
		src_hash = binary.SrcHash
		name = binary.Name
	case config.Extend:
		extend := *downloadableHolder.ExtendPointer
		target = extend.GetDownloadPath()
		src = extend.Src
		src_hash = extend.SrcHash
		name = src
	default:
		return "", errors.New(fmt.Sprintf("not a valid Downloadable type: %s", t))
	}
	if c.IsDryRun {
		fmt.Printf("DRY-RUN: Downloading %s from %s ...\n", name, src)
		fmt.Printf("DRY-RUN: Successfully downloaded %s to %s\n", name, target)
		return target, nil
	}
	fmt.Printf("Downloading %s from %s ...\n", name, src)
	out, err := os.Create(target)
	if err != nil {
		return "", err
	}
	defer out.Close()

	resp, err := http.Get(src)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	hashExists, err := checkHashIfExists(respBodyBytes, src_hash)
	if err != nil {
		return "", err
	}

	_, err = out.Write(respBodyBytes) 
	if err != nil {
		return "", err
	}

	if hashExists {
		fmt.Printf("Successfully downloaded %s to %s with hashsum %s\n", name, target, src_hash)
	} else {
		fmt.Printf("Successfully downloaded %s to %s\n", name, target)
	}
	return target, nil
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
