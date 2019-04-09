// see: https://medium.com/@j.d.livni/create-a-load-bar-in-go-f158837ff4c4
package main

import (
	"fmt"
	"strings"
	"time"
)

func showProgressBar(iteration int, total int, prefix string, suffix string, length int, fill string) {
	percent := float64(iteration) / float64(total)
	filledLength := int(length * iteration / total)
	end := ">"

	if iteration == total {
		end = "="
	}

	bar := strings.Repeat(fill, filledLength) + end + strings.Repeat("-", (length-filledLength))
	fmt.Printf("\r%s [%s] %f%% %s", prefix, bar, percent, suffix)
	if iteration == total {
		fmt.Println()
	}
}

func main() {
	for i := 0; i < 30; i++ {
		time.Sleep(500 * time.Millisecond)
		showProgressBar(i+1, 30, "Progress", "Complete", 25, "=")
	}
}
