package onewolf

import (
	"io"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/taylormonacelli/coalfoot"
	"golang.org/x/tools/txtar"
)

func Main() int {
	slog.Debug("onewolf", "test", true)
	x := coalfoot.NewTxtarTemplate()
	x.FetchFromRemoteIfOld()
	archive, err := txtar.ParseFile(x.LocalPathUnrendered)
	if err != nil {
		slog.Error("error", "error", err.Error())
		return 1
	}

	slog.Debug("archive", "length", len(archive.Files))

	var file txtar.File

	if len(archive.Files) == 0 {
		slog.Debug("filename", "filename", file.Name)
		return 1
	}

	file = archive.Files[0]

	filePath := filepath.Join("/tmp", file.Name)
	slog.Debug("creating file", "path", filePath)

	f, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = io.WriteString(f, string(file.Data))
	if err != nil {
		panic(err)
	}

	println("File has been written successfully.")

	return 0
}
