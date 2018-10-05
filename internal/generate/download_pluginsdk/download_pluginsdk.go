package main

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You need a URL to download the TS3 Plugin SDK ZIP file from.")
		return
	}
	if len(os.Args) < 3 {
		fmt.Println("You need a target path to extra the TS3 Plugin SDK ZIP file to.")
		return
	}

	url := os.Args[1]
	target := os.Args[2]

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error starting HTTP request to donwload ZIP:", err)
		os.Exit(1)
		return
	}
	defer resp.Body.Close()

	f, err := ioutil.TempFile("", "ts3plugin-")
	if err != nil {
		fmt.Println("Error creating temporary file to download ZIP to:", err)
		os.Exit(1)
		return
	}
	defer func() {
		f.Close()
		os.Remove(f.Name())
	}()

	fmt.Println("Downloading:", url, "=>", f.Name())
	written, err := io.Copy(f, resp.Body)
	if err != nil {
		fmt.Println("Error saving ZIP file:", err)
		os.Exit(1)
		return
	}

	// f.Close()
	// f, err = os.Open(f.Name())
	// if err != nil {
	// 	fmt.Println("Error opening temporary ZIP file:", err)
	// 	os.Exit(1)
	// 	return
	// }

	_, err = f.Seek(0, os.SEEK_SET)
	if err != nil {
		fmt.Println("Error seeking to beginning of ZIP after download:", err)
		os.Exit(1)
		return
	}

	fmt.Println("Unzipping:", f.Name(), "=>", target)
	_, err = Unzip(f, written, target)
	if err != nil {
		fmt.Println("Error unpacking ZIP file:", err)
		os.Exit(1)
		return
	}
}

func Unzip(src io.ReaderAt, size int64, dest string) ([]string, error) {
	// mostly copy-paste from https://stackoverflow.com/questions/20357223/easy-way-to-unzip-file-with-golang

	var filenames []string

	r, err := zip.NewReader(src, size)
	if err != nil {
		return filenames, err
	}

	dest, err = filepath.Abs(dest)
	if err != nil {
		return filenames, err
	}

	for _, f := range r.File {

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}
		defer rc.Close()

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)

		// Check for ZipSlip. More Info: http://bit.ly/2MsjAWE
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {

			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)

		} else {

			// Make File
			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				return filenames, err
			}

			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return filenames, err
			}

			_, err = io.Copy(outFile, rc)

			// Close the file without defer to close before next iteration of loop
			outFile.Close()

			if err != nil {
				return filenames, err
			}

		}
	}
	return filenames, nil
}
