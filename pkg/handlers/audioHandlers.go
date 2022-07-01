package handlers

import (
	"bytes"
	"encoding/json"
	errs "github.com/spear-app/spear-go/pkg/err"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

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
	tmpfile, err := os.Create("./" + h.Filename)
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
