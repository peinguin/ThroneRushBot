package utils

import (
	"os"
    "path/filepath"
    "html/template"
    "log"
    "bytes"
)

func Template(name string) *template.Template{
	var buffer bytes.Buffer
    dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
        log.Fatal(err)
    }
    buffer.WriteString(dir)
    buffer.WriteString("/")
    buffer.WriteString("static/")
    buffer.WriteString(name)
    buffer.WriteString(".html")

    t, err := template.ParseFiles(buffer.String())
    if err != nil {
        log.Fatal("There was an error:", err)
    }
    return t
}