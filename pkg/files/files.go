package files

import (
	"os"
    "io"
	"strings"
    "net/http"
)

func SanitizeFilename(file string) string {
	
	file = strings.Join(strings.Fields(file),"")
	file = strings.ReplaceAll(file, "//", "-")
	file = strings.ReplaceAll(file, "/", "-")
	file = strings.ReplaceAll(file, ">", "-")
	file = strings.ReplaceAll(file, ":", "-")
	file = strings.ReplaceAll(file, ".", "_")
	file = strings.ReplaceAll(file, "@", "-AT-")
	file = strings.ReplaceAll(file, "https", "-")
	file = strings.ReplaceAll(file, "http", "")
	file = strings.ReplaceAll(file, "--", "")
	file = strings.TrimPrefix(file, "-")
	file = strings.TrimSuffix(file, "-")

	return file
}

func FileMissingErr(err error) bool {
  if err == nil { return false }
	if strings.Contains(err.Error(), "no such file") { return true }
	if strings.Contains(err.Error(), "cannot find the file") { return true }
	return false
}


func DownloadFile(url string, dir string, clearFolder ...bool) (string, error) {

	// get filename
	filename := NameFromPath(url)
	filepath := dir + filename

	// validate dir
	if !strings.HasSuffix(dir, "/") {
		dir = dir + "/"
	}
	
	if len(clearFolder) > 0 {
		os.Remove(dir)
	} 

	// make directory, ignore error
	MakeDirectory(dir)

	// don't worry about errors
	response, err := http.Get(url)
	if err != nil {
		return filepath, err
	}
	defer response.Body.Close()

	//open a file for writing
	file, err := os.Create(filepath)
	if err != nil {
		return filepath, err
	}
	defer file.Close()

	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return filepath, err
	}

	return filepath, nil
}

func Delete(filepath string) error {
	err := os.Remove(filepath) 
	return err
}

func NameFromPath(filepath string) string {
	spl := strings.Split(filepath, "/")
	if len(spl) == 0 {return filepath}
	return spl[len(spl)-1]
}

func Write(filepath string, content string) error {
 	err := os.WriteFile(filepath, []byte(content), 0644)
	return err
}
