package archive_pkg

import (
    "archive/tar"
    "bytes"
)

type Archive interface {
    Operate() error
    Close() error
}

type ArchiveWriter struct {
    tw      *tar.Writer
}

type File struct {
    Name    string
    Body    string
}

func (aw *ArchiveWriter) Operate() error {
    return aw.write()
}

func (aw *ArchiveWriter) write() error {
    buf := new(bytes.Buffer)
    tw := tar.NewWriter(buf)

    var files = []File{
        {"readme.txt", "This is archive contains some text files."},
        {"gopher.txt", "Gopher names:\nAaron\nShawn\nJeffery"},
        {"todo.txt", "Get animal handling licence."},
    }

    for _, file := range files {
        hdr := &tar.Header{
            Name: file.Name,
            Size: int64(len(file.Body)),
        }
        if err := tw.WriteHeader(hdr); err != nil {
            return err
        }
        if _, err := tw.Write([]byte(file.Body)); err != nil {
            return err
        }
    }

    return nil
}

func (aw *ArchiveWriter) Close() error {
    return aw.tw.Close()
}


type ArchiveReader struct {
    tw      *tar.Reader
}

func (ar *ArchiveReader) Operate() error {
    return nil
}

func (ar *ArchiveReader) read() error {
    return nil
}

func (ar *ArchiveReader) Close() error {
    return nil
}
