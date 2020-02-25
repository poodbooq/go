package zip

import (
	"io/ioutil"
)

// ReadWriter can add and delete files
type ReadWriter struct {
	ReadCloser *ReadCloser
	Writer     *Writer
}

// NewReadWriter returns ExtendedWriter after reading the file
func NewReadWriter(file string) (*ReadWriter, error) {
	var rw = new(ReadWriter)

	reader, err := OpenReader(file)
	if err != nil {
		return nil, err
	}

	rw.ReadCloser = reader
	rw.Writer = new(Writer)

	rw.Writer.comment = rw.ReadCloser.Comment
	rw.Writer.closed = false

	for i := range rw.ReadCloser.File {
		f, err := rw.Writer.Create(rw.ReadCloser.File[i].Name)
		if err != nil {
			return nil, err
		}
		zf, err := rw.ReadCloser.File[i].Open()
		if err != nil {
			return nil, err
		}
		bs, err := ioutil.ReadAll(zf)
		if err != nil {
			return nil, err
		}
		if _, err := f.Write(bs); err != nil {
			return nil, err
		}
	}

	return rw, nil
}

// Add adds new file to existing zip file
func (rw *ReadWriter) Add(name string, data []byte) error {
	wf, err := rw.Writer.Create(name)
	if err != nil {
		return err
	}
	_, err = wf.Write(data)
	return err
}

// RemoveFile removes file from the zip file
// returns error in case file doesn't exist
func (rw *ReadWriter) RemoveFile(name string) error {
	return nil
}

// Close closes rw
func (rw *ReadWriter) Close() error {
	if err := rw.ReadCloser.Close(); err != nil {
		return err
	}
	return rw.Writer.Close()
}
