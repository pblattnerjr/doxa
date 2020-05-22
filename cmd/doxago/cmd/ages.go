package cmd

import (
	"fmt"
	agesAres "github.com/liturgiko/doxa/pkg/ages/ares"
	agesAtem "github.com/liturgiko/doxa/pkg/ages/atem2liml"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
	"time"
)

// Note: the ages command will go away once AGES is using doxa instead of ALWB
var agesCmd = &cobra.Command{
	Use:   "ages",
	Short: "ages provides commands to prepare legacy AGES ares and atem files for use in Doxa",
	Long: `ages provides commands to prepare legacy AGES ares and atem files for use in Doxa.
The ages command will be removed once AGES, Initiatives is using doxa instead of ALWB.
`,
	Run: func(cmd *cobra.Command, args []string) {
		start := time.Now()

		// setRecord popPath the logger, which will be passed to the functions that do the processing
		LogFile, err := os.OpenFile(LogFilename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
		if err != nil {
			panic(err)
		}
		defer LogFile.Close()

		Logger.SetOutput(LogFile)
		Logger.SetFlags(log.Ldate + log.Ltime + log.Lshortfile)
		a2l, _ := cmd.Flags().GetBool("atem2lml")
		fixares, _ := cmd.Flags().GetBool("cleanares")

		if a2l {
			msg := fmt.Sprintf("converting atem to lml...\n")
			fmt.Println(msg)
			Logger.Println(msg)
			if err = agesAtem.Process(Paths.SysPath, Paths.AtemPath, Paths.TemplatesPath); err != nil {
				Logger.Println(err.Error())
			}
		} else if fixares {
			in := filepath.Join(Paths.ReposPath,"ares")
			out := filepath.Join(Paths.HomePath, "fixed/ares")
			fmt.Printf("Fixing problems in ares files and writing new files to %s\n", out)
			if err = agesAres.CleanAresFiles(in,out,false,&Logger); err != nil {
				Logger.Println(err.Error())
			}
		} else {
			fmt.Printf("You must provide a flag, e.g. to convert templates: doxago ages --atem2lml; to fix ares: doxago ages --cleanares\n")
		}
		Elapsed(start)
	},
}

func init() {
	rootCmd.AddCommand(agesCmd)
	agesCmd.Flags().Bool("atem2lml" , false, "convert atem templates to lml files")
	agesCmd.Flags().Bool("cleanares", false, "analyze ares files and fix problems")
}

