// root.go
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/word-count/pkg"
)


func init() {
	rootCmd.Flags().StringP("file", "f", "newFile.txt", "File path")
	rootCmd.Flags().IntP("routines", "r", 1, "Number of routines")

	
}

var rootCmd = &cobra.Command{
	Use:   "word-count",
	Short: "Count statistics in a file",
	Run: func(cmd *cobra.Command, args []string) {
		filePath, _ := cmd.Flags().GetString("file")
		routines, _ := cmd.Flags().GetInt("routines")
		processFile(filePath, routines)
	},
}



func processFile(filePath string, routines int) {
	start := time.Now()
	results := make(chan pkg.CountsResult, routines)

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	fileSize := stat.Size()
	chunkSize := fileSize / int64(routines)

	reader := bufio.NewReader(file)

	for i := 0; i < routines; i++ {
		chunk := make([]byte, chunkSize)
		_, err := reader.Read(chunk)
		if err != nil {
			log.Fatal(err)
		}

		go pkg.Counts(chunk, results)
	}

	totalCounts := pkg.CountsResult{}

	for i := 0; i < routines; i++ {
		result := <-results
		totalCounts.LineCount += result.LineCount
		totalCounts.WordsCount += result.WordsCount
		totalCounts.VowelsCount += result.VowelsCount
		totalCounts.PunctuationCount += result.PunctuationCount
	}
	fmt.Println("no of ",routines)
	fmt.Println("Number of lines:", totalCounts.LineCount)
	fmt.Println("Number of words:", totalCounts.WordsCount)
	fmt.Println("Number of vowels:", totalCounts.VowelsCount)
	fmt.Println("Number of punctuation:", totalCounts.PunctuationCount)
	fmt.Println("Run Time:", time.Since(start))
}

func Execute() {
	rootCmd.Execute()
}
