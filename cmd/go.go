/*
Copyright © 2022 srjchsv@gmail.com
*/

package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/srjchsv/gosyncfolders/internal/scan"
	"github.com/srjchsv/gosyncfolders/pkg/utils"
)

var (
	mu sync.RWMutex
)

// goCmd represents the go command
var goCmd = &cobra.Command{
	Use:   "go",
	Short: "Would start the folders sync",
	Long: `
	This command will start sync source and destination folders that you have to input.

	For example:
	gosyncfolders go src/ dst/ 

	src/ as source folder in current directory. 
	dst/ as the destination folder in current directory.

	And the sync will start.

`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			fmt.Printf("You have to add two arguments with source/ and destination/ folders after the gosyncfolders go")
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
		fmt.Println("=========gosyncfolders========")
		fmt.Println("==============================")
		src := args[0]
		dst := args[1]

		for {
			if _, err := os.Stat(src); os.IsNotExist(err) {
				fmt.Println("Source folder path not found. Enter the right path:")
				fmt.Scanln(&src)
				continue
			}

			if _, err := os.Stat(dst); os.IsNotExist(err) {
				fmt.Println("Destination folder path not found. Enter the right path:")
				fmt.Scanln(&dst)
				continue
			}
			break
		}

		fmt.Printf("Source folder path:\n%v\n", src)
		fmt.Printf("Destination folder path:\n%v\n", dst)

		utils.PrettyConsole()

		go scan.Source(src, dst, &mu)
		go scan.Destination(src, dst, &mu)

		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-c
		log.Info("Shutting down program...")

	},
}

func init() {
	rootCmd.AddCommand(goCmd)
}
