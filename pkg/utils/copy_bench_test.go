package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(ioutil.Discard)
}

func BenchmarkCopyFilesIoCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := fmt.Sprintf("../../test/source/%v", i)
		dst := fmt.Sprintf("../../test/destination/%v", i)
		os.Create(src)
		for i := 0; i < b.N; i++ {
			CopyFilesIoCopy(src, dst)
		}
	}
}

func BenchmarkCopyFilesIoutil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := fmt.Sprintf("../../test/source/%v", i)
		dst := fmt.Sprintf("../../test/destination/%v", i)
		os.Create(src)
		for i := 0; i < b.N; i++ {
			CopyFilesIoutil(src, dst)
		}
	}
}
