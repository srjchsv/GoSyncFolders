package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
)

// CopyFilesIoutil copies files
func CopyFilesIoutil(src, dst string) error {
	return filepath.Walk(src, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// copy to this path
		outpath := filepath.Join(dst, strings.TrimPrefix(path, src))

		if info.IsDir() {
			err := os.MkdirAll(outpath, info.Mode())
			if err != nil {
				return err
			}
			log.Infof("copied folder:%v", info.Name())
			return nil // means recursive
		}
		// handle irregular files
		if !info.Mode().IsRegular() {
			switch info.Mode().Type() & os.ModeType {
			case os.ModeSymlink:
				link, err := os.Readlink(path)
				if err != nil {
					return err
				}
				return os.Symlink(link, outpath)
			}
			return nil
		}
		// open input
		in, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		// copy content
		err = ioutil.WriteFile(outpath, in, info.Mode())
		if err != nil {
			return err
		}
		size, err := os.Stat(outpath)
		if err != nil {
			return err
		}

		log.Infof("Copied file of size %d bytes: %v", size.Size(), info.Name())
		return err
	})
}

// CopyFilesIoCopy copies files
func CopyFilesIoCopy(src, dst string) error {
	return filepath.Walk(src, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// copy to this path
		outpath := filepath.Join(dst, strings.TrimPrefix(path, src))

		if info.IsDir() {
			err := os.MkdirAll(outpath, info.Mode())
			if err != nil {
				return err
			}
			return nil // means recursive
		}
		// handle irregular files
		if !info.Mode().IsRegular() {
			switch info.Mode().Type() & os.ModeType {
			case os.ModeSymlink:
				link, err := os.Readlink(path)
				if err != nil {
					return err
				}
				return os.Symlink(link, outpath)
			}
			return nil
		}
		// open input
		in, err := os.Open(path)
		if err != nil {
			return err
		}
		defer in.Close()
		// create output
		fh, err := os.Create(outpath)
		if err != nil {
			return err
		}
		defer fh.Close()

		// make it the same
		err = fh.Chmod(info.Mode())
		if err != nil {
			return err
		}
		// copy content
		_, err = io.Copy(fh, in)
		if err != nil {
			return err
		}

		// Get the size of the copied file and log it
		size, err := fh.Stat()
		if err != nil {
			return err
		}
		log.Infof("Copied file size is %d bytes\n", size.Size())
		return err
	})
}

func Hash(path string) (value string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, f); err != nil {
		log.Fatal(err)
	}
	value = hex.EncodeToString(hasher.Sum(nil))
	return
}
