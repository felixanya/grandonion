package archive_pkg

import (
    "archive/tar"
    "bytes"
    "fmt"
    "io"
    "log"
    "os"
    "testing"
)

func TestArchiveWriter_Operate(t *testing.T) {
    // Create a buffer to write our archive to.
    buf := new(bytes.Buffer)
    // Create a new tar archive.
    tw := tar.NewWriter(buf)
    // Add some files to the archive.
    var files = []struct {
        Name, Body string
    }{
        {"readme.txt", "This archive contains some text files."},
        {"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
        {"todo.txt", "Get animal handling licence."},
    }
    for _, file := range files {
        hdr := &tar.Header{
            Name: file.Name,
            Size: int64(len(file.Body)),
        }
        if err := tw.WriteHeader(hdr); err != nil {
            log.Fatalln(err)
        }
        if _, err := tw.Write([]byte(file.Body)); err != nil {
            log.Fatalln(err)
        }
    }
    // Make sure to check the error on Close.
    if err := tw.Close(); err != nil {
        log.Fatalln(err)
    }
    // Open the tar archive for reading.
    r := bytes.NewReader(buf.Bytes())
    tr := tar.NewReader(r)
    // Iterate through the files in the archive.
    for {
        hdr, err := tr.Next()
        if err == io.EOF {
            // end of tar archive
            break
        }
        if err != nil {
            log.Fatalln(err)
        }
        fmt.Printf("Contents of %s:\n", hdr.Name)
        if _, err := io.Copy(os.Stdout, tr); err != nil {
            log.Fatalln(err)
        }
        fmt.Println()
    }
}
