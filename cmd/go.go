/*
Copyright © 2022 srjchsv@gmail.com
*/

package cmd

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/srjchsv/gosyncfolders/internal/scan"
	"os"
	"sync"
	"time"
)

var (
	consoleLoop       = 30
	console1LineBreak = 10
	console2LineBreak = 12
	console3LineBreak = 17
	mu                sync.RWMutex
)

// goCmd represents the go command
var goCmd = &cobra.Command{
	Use:   "go",
	Short: "Would start the folders sync",
	Long: `
	This command will start sync source and destination folders that you have to input.

	For example:
	wwww go src/ dst/ 

	src/ as source folder in current directory. 
	dst/ as the destination folder in current directory.

	And the sync will start.

`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			fmt.Printf("You have to add two arguments with source/ and destination/ folders after the wwww go")
			return
		case 1:
			fmt.Printf("You have to add one more argument with destination/ folder")
			return
		case 2:
			break
		default:
			fmt.Println("You can only add two arguments, for example: source/ destination/ ")
			return

		}
		LogFile, err := os.OpenFile("log.txt", os.O_TRUNC|os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
			err = LogFile.Close()
			if err != nil {
				panic(err)
			}
		}()

		log.SetOutput(LogFile)

		if err != nil {
			log.Panic(err)
		}
		log.SetFormatter(&log.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			PrettyPrint:     true,
		})

		log.Info("Starting the program")
		fmt.Println("==============================")
		fmt.Println("=========wwww========")
		fmt.Println("==============================")
		src := args[0]
		dst := args[1]

		_, err2 := os.Stat(src)
		if err2 != nil {
			err2 := os.Mkdir(src, 0750)
			if err2 != nil {
				log.Errorf("err2or creating source path:%v", err2)
			}
		}

		fmt.Printf("Source folder path:\n%v\n", src)
		fmt.Printf("Destination folder path:\n%v\n", dst)

		for i := 0; i < consoleLoop; i++ {
			fmt.Print("*")
			if i == console1LineBreak {
				fmt.Println()
			}
			if i == console2LineBreak {
				fmt.Print("Sync Started")
			}
			fmt.Print("*")
			if i == console3LineBreak {
				fmt.Println()
			}

			fmt.Print("*")
			time.Sleep(time.Millisecond * 30)
		}

		ctx, cancel := context.WithCancel(context.TODO())

		var errCh chan error
		go scan.Source(ctx, src, dst, &mu, errCh)
		go scan.Destination(ctx, src, dst, &mu, errCh)
		<-errCh
		cancel()
		log.Info("Exiting program...")
	},
}

func init() {
	rootCmd.AddCommand(goCmd)
}
