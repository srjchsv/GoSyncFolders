package utils

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"testing"
)

var (
	srcPath  = "../../test/source/source0"
	dstPath  = "../../test/destination/source0"
	srcPath2 = "../../test/source/source1"
	dstPath2 = "../../test/destination/source1"
)

func init() {

	logrus.SetOutput(ioutil.Discard)

	os.MkdirAll("../../test/source", 0750)
	os.MkdirAll("../../test/destination", 0750)
	for i := 0; i < 2; i++ {
		os.Create(fmt.Sprintf("../../test/source/source%v", i))
		os.Create(fmt.Sprintf("../../test/destination/source%v", i))
	}
}

func TestCopyFilesIoutil(t *testing.T) {
	req := require.New(t)
	testCase := func(src, dst string, want bool) func(t *testing.T) {
		return func(t *testing.T) {
			res := CopyFilesIoutil(src, dst)
			req.NoError(res)
			_, err := os.Stat(srcPath)
			var exist bool
			if err == nil {
				exist = true
			}
			req.Equal(want, exist)
		}
	}
	t.Run("copy file", testCase(srcPath, dstPath, true))
}

func TestCopyFilesIoCopy(t *testing.T) {
	req := require.New(t)
	testCase := func(src, dst string, want bool) func(t *testing.T) {
		return func(t *testing.T) {
			res := CopyFilesIoCopy(src, dst)
			req.NoError(res)
			_, err := os.Stat(srcPath2)
			var exist bool
			if err == nil {
				exist = true
			}
			req.Equal(want, exist)
			os.RemoveAll("../../test/")
		}
	}
	t.Run("copy file", testCase(srcPath2, dstPath2, true))
}
