package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var vlcCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Pack file using varible-length code",
	Run:   pack,
}

const packedExtention = "vlc"

var ErrEmptyPath = errors.New("path to files is nit specified")

func pack(_ *cobra.Command, args []string) {
	if len(args) == 0 || args[0] == "" {
		handleErr(ErrEmptyPath)
	}
	filePath := args[0]

	r, err := os.Open(filePath)
	if err != nil {
		handleErr(err)
	}

	data, err := io.ReadAll(r)
	if err != nil {
		handleErr(err)
	}

	packed := ""
	fmt.Println(string(data))

	err = os.WriteFile(packedFileName(filePath), []byte(packed), 0644)
	if err != nil {
		handleErr(err)
	}

}

func packedFileName(path string) string {
	fileName := filepath.Base(path)

	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "." + packedExtention
}

func init() {
	packCmd.AddCommand(vlcCmd)
}
