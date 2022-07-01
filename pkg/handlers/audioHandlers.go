package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	errs "github.com/spear-app/spear-go/pkg/err"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
)

type textAndDiarization struct {
	text        string
	diarization string
}

func Wav(w http.ResponseWriter, r *http.Request) {
	//dat, err := ioutil.ReadFile("public_path()" + "/forTest/record.wav")
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.FormDataContentType()
	file, h, err := r.FormFile("audio")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errs.NewResponse(err.Error(), http.StatusBadRequest))
		return
	}
	filePath := "/home/rahma/conversation_audio/" + h.Filename
	tmpfile, err := os.Create(filePath)
	defer tmpfile.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errs.NewResponse(err.Error(), http.StatusInternalServerError))
		return
	}
	_, err = io.Copy(tmpfile, file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errs.NewResponse(err.Error(), http.StatusInternalServerError))
		return
	}
	w.WriteHeader(200)
	return
}

func GetText(filePath string) (string, error) {
	prg := "/usr/bin/python3"
	arg1 := "/home/rahma/spear-python/speech_to_text.py"
	cmd, err := exec.Command(prg, arg1, filePath).Output()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(cmd), nil
}
