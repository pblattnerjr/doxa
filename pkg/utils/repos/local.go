package repos

import (
	"bufio"
	"fmt"
	"github.com/liturgiko/doxa/pkg/ages/ares"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// maps lines from ares files in Github repositories
// to LineParts
func Ares2LpFromLocalDir(rootDir string,
	fileSuffix string,
	printProgress bool,
	logger *log.Logger) <-chan *ares.LineParts {
	out := make(chan *ares.LineParts)
	go func() {
		err := filepath.Walk(rootDir,
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if strings.HasSuffix(info.Name(), ".git") {
					fmt.Println(fmt.Sprintf("Processing: %s", path[:len(path)-5]))
				}
				if !info.IsDir() && strings.HasSuffix(path, fileSuffix) {
					file, err := os.Open(path)
					if err != nil {
						log.Fatal(err)
					}
					defer file.Close()

					scanner := bufio.NewScanner(file)
					var fileNameParts ares.FilenameParts
					fileNameParts, err = ares.ParseAresFileName(file.Name())

					if printProgress {
						fmt.Printf("\r%-80s", "")
						fmt.Print(fmt.Sprintf("\rProcessing %s_%s_%s_%s.ares", fileNameParts.Topic, fileNameParts.Language, fileNameParts.Country, fileNameParts.Realm))
					}
					var lineCnt int
					inCommentBlock := false

					for scanner.Scan() {

						lineCnt = lineCnt + 1
						line := strings.TrimSpace(scanner.Text())
						line = strings.TrimSpace(line)
						if strings.HasPrefix(line, "/*") {
							inCommentBlock = true
							continue
						}
						if strings.HasPrefix(line, "*/") || strings.HasSuffix(line, "*/") {
							inCommentBlock = false
							continue
						}
						if !inCommentBlock {
							lineParts, err := ares.ParseLine(fileNameParts, line)
							if err != nil {
								logger.Printf("%s: %s: line %d", err.Error(), file.Name(), lineCnt)
								continue
							}
							lineParts.LineNbr = lineCnt
							if lineParts.IsAresId || lineParts.IsBlank || lineParts.IsCommentedOut {
								// ignore
							} else {
								out <- &lineParts
							}
						}
					}
				}
				return nil
			})
		if err != nil {
			log.Println(err)
		}
		close(out)
	}()
	return out
}
