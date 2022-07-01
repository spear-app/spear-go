package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	errs "github.com/spear-app/spear-go/pkg/err"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"time"
)

type textAndDiarization struct {
	text        string
	diarization string
}

var ConversationStarTime time.Time

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
	var textAndSpeakerResponse textAndDiarization
	textAndSpeakerResponse.text, err = GetText(filePath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errs.NewResponse("couldn't get text from speech", http.StatusInternalServerError))
		return
	}
	// so far have text, now we need to get speaker diarization
	audioPlayTime := time.Now()
	err = PlayAudio(filePath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errs.NewResponse("couldn't play audio", http.StatusInternalServerError))
		return
	}
	duration, err := SubtractTime(ConversationStarTime, audioPlayTime)
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

func PlayAudio(filePath string) error {
	prg := "/usr/bin/play"
	_, err := exec.Command(prg, filePath).Output()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
func StartConversation(w http.ResponseWriter, r *http.Request) {

	ConversationStarTime = time.Now()

}
func SubtractTime(time1 time.Time, time2 time.Time) (int, error) {
	hour1, min1, second1 := time1.Clock()
	hour2, min2, second2 := time2.Clock()
	if hour2-hour1 != 0 {
		return 0, errors.New("max conversation time is 15 minutes")
	}
	duration := (min2*60 + second2) - (min1*60 + second1)
	if duration <= 0 {
		return 0, errors.New("invalid time duration")
	}
	return duration, nil
}
