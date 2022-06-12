package scan

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"sync"
	"testing"
	"time"
)

var (
	toCopy      int
	srcPath     = "../../test/source"
	dstPath     = "../../test/destination"
	mu          sync.RWMutex
	timeout     = time.Millisecond * 500
	ctx, cancel = context.WithCancel(context.TODO())
	errCh       chan error
)

func init() {
	logrus.SetOutput(ioutil.Discard)
	os.MkdirAll(srcPath, 0750)
	os.MkdirAll(dstPath, 0750)
	if _, err := os.Stat(srcPath); err == nil {
		for i := 0; i < 3; i++ {
			os.Create(fmt.Sprintf("../../test/source/%v.txt", i))
		}
	}
	time.Sleep(time.Second + 2)
	files, _ := ioutil.ReadDir(srcPath)
	toCopy = len(files)
}

func TestSource(t *testing.T) {
	req := require.New(t)
	testCase := func(src, dst string, mu *sync.RWMutex, timeout time.Duration, want int) func(t *testing.T) {
		return func(t *testing.T) {
			go Source(ctx, src, dst, mu, errCh)
			time.Sleep(time.Second * 2)
			files, _ := ioutil.ReadDir(dstPath)
			res := len(files)
			req.Equal(want, res)
		}
	}
	t.Run("sync folders", testCase(srcPath, dstPath, &mu, timeout, toCopy))
}

func TestDestination(t *testing.T) {
	req := require.New(t)
	testCase := func(want int) func(t *testing.T) {
		return func(t *testing.T) {
			os.RemoveAll(srcPath)
			os.MkdirAll(srcPath, 0750)
			files, _ := ioutil.ReadDir(srcPath)
			res := len(files)
			go Destination(ctx, srcPath, dstPath, &mu, errCh)
			time.Sleep(time.Second * 2)
			req.Equal(want, res)
			os.RemoveAll("../../test/")
		}
	}
	t.Run("sync folders", testCase(0))
}
