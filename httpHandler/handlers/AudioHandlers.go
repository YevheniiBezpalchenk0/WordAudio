package handlers

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
)

const uploadPath = "./Audio"

func GetAudioHandler(w http.ResponseWriter, r *http.Request) {
	fileName := "./Audio/output.mp3"
	updateFile, err := os.Open(fileName)

	defer func(updateFile *os.File) {
		err := updateFile.Close()
		if err != nil {
			http.Error(w, "Update file not found.", 404)
			return
		}
	}(updateFile)

	if err != nil {
		http.Error(w, "Update file not found.", 404)
		return
	}

	fileHeader := make([]byte, 512)
	_, err = updateFile.Read(fileHeader)
	if err != nil {
		http.Error(w, "Update file not found.", 404)
		return
	}
	fileType := http.DetectContentType(fileHeader)

	fileInfo, _ := updateFile.Stat()
	fileSize := fileInfo.Size()

	w.Header().Set("Expires", "0")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Content-Control", "private, no-transform, no-store, must-revalidate")
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", fileType)
	w.Header().Set("Content-Length", strconv.FormatInt(fileSize, 10))

	_, err = updateFile.Seek(0, 0)
	if err != nil {
		http.Error(w, "Update file not found.", 404)
		return
	}
	_, err = io.Copy(w, updateFile)
	if err != nil {
		http.Error(w, "Update file not found.", 404)
		return
	}
	return
}

func SendAudioHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(100)
	if err != nil {
		return
	}
	mForm := r.MultipartForm

	for k := range mForm.File {
		// k is the key of file part
		file, fileHeader, err := r.FormFile(k)
		if err != nil {
			fmt.Println("FormFile error:", err)
			return
		}
		defer func(file multipart.File) {
			err := file.Close()
			if err != nil {
				http.Error(w, "Update file not found.", 404)
				return
			}
		}(file)
		fmt.Printf("the uploaded file: name[%s], size[%d], header[%#v]\n",
			fileHeader.Filename, fileHeader.Size, fileHeader.Header)

		// store uploaded file into local path
		localFileName := uploadPath + "/" + fileHeader.Filename
		out, err := os.Create(localFileName)
		if err != nil {
			fmt.Printf("failed to open the file %s for writing", localFileName)
			return
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			fmt.Printf("copy file err:%s\n", err)
			return
		}
		fmt.Printf("file %s uploaded ok\n", fileHeader.Filename)
	}

}
