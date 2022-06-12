package scan

import (
	"context"
	"github.com/srjchsv/gosyncfolders/pkg/utils"
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"sync"

	log "github.com/sirupsen/logrus"
)

// Source folder scan to sync with destination folder
func Source(ctx context.Context, src, dst string, mu *sync.RWMutex, ch chan error) {
	mu.RLock()
	defer mu.RUnlock()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	tick := time.NewTicker(time.Second)
	defer tick.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-tick.C:
			srcScan := filepath.Walk(src, func(path string, info fs.FileInfo, err error) error {
				if path != src {
					path = path[len(src):]
					dstPath := filepath.Join(dst, path)
					srcPath := filepath.Join(src, path)
					_, err := os.Stat(dstPath)
					if err != nil {
						err := utils.CopyFilesIoutil(srcPath, dstPath)
						if err != nil {
							ch <- err
							log.Errorf("error copying files: %v", err)
						}
					} else {
						if utils.Hash(srcPath) != utils.Hash(dstPath) {
							log.Infof("File %v in source edited and will be copied.", srcPath)
							err := utils.CopyFilesIoutil(srcPath, dstPath)
							if err != nil {
								ch <- err
								log.Errorf("error copying files: %v", err)
							}
						}
					}
				}
				return err
			})
			if srcScan != nil {
				log.Errorf("error scanning the source folder: %v", srcScan)
			}
		}
	}
}

// Destination folder scan for unwanted files and deletes them
func Destination(ctx context.Context, src, dst string, mu *sync.RWMutex, ch chan error) {
	mu.RLock()
	defer mu.RUnlock()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	tick := time.NewTicker(time.Second)
	defer tick.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-tick.C:
			dstScan := filepath.Walk(dst, func(path string, info fs.FileInfo, err error) error {
				if path != dst {
					path = path[len(dst):]
					dstPath := filepath.Join(dst, path)
					srcPath := filepath.Join(src, path)
					_, err := os.Stat(srcPath)
					if err != nil {
						file, err := os.Stat(dstPath)
						if err != nil {
							log.Errorf("error getting the file stats: %v", err)
						}
						err = os.Remove(dstPath)
						if err != nil {
							ch <- err
							log.Errorf("error removing the files: %v", err)
						}
						log.Infof("removed %v bytes of unwanted files: %v :", file.Size(), dstPath)
					}
				}
				return err
			})
			if dstScan != nil {
				log.Errorf("error scanning the destination folder: %v", dstScan)
			}
		}
	}
}
