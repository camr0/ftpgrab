//go:build !windows
// +build !windows

package grabber

import (
	"os"
	"io"
)

func moveFile(oldpath, newpath string) error {
	// return os.Rename(oldpath, newpath)

	// https://stackoverflow.com/questions/1821811/how-to-read-write-from-to-a-file-using-go
	fOld, err := os.Open(oldpath)
	if err != nil {
			panic(err)
	}
	// close fOld on exit and check for its returned error
	// TODO: is this exception necessary?
	defer func() {
			if err := fOld.Close(); err != nil {
					panic(err)
			}
	}()

	fNew, err := os.Create(newpath)
	if err != nil {
			panic(err)
	}
	// close fNew on exit and check for its returned error
	// TODO: is this exception necessary?
	defer func() {
			if err := fNew.Close(); err != nil {
					panic(err)
			}
	}()

	_, err = io.Copy(fNew, fOld)
	if err != nil {
		panic(err)
	}

	err = os.Remove(oldpath)

	return err
}
