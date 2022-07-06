package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	errs "github.com/spear-app/spear-go/pkg/err"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"syscall"
	"time"
)

type textAndDiarization struct {
	text        string
	diarization string
}

type StartConv struct {
	Start_conversation bool `json:"start_conversation"`
}

var ConversationStarTime time.Time
var cmd *exec.Cmd

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
	println(duration)
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

func tmpStartConversation() (time.Time, error) {
	timeout := time.After(10 * time.Second)
	ticker := time.Tick(500 * time.Millisecond)

	// Keep trying until we're timed out or get a result/error
	for {
		select {
		case <-timeout:
			return ConversationStarTime, errors.New("timed out, can't start conversation")
		case <-ticker:
			ok, err := exec.Command("source", "/home/rahma/GolandProjects/spear-go/pkg/scripts/diart_run.sh").Output()
			if err != nil {
				return ConversationStarTime, errors.New("couldn't set environment for diart")
			} else if len(string(ok)) > 5 {
				ConversationStarTime = time.Now()
				return ConversationStarTime, nil
			}
		}
	}
}

func StartConversation(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	//extracting usr obj
	// TODO get user id
	var strartConv StartConv
	json.NewDecoder(r.Body).Decode(&strartConv)
	if strartConv.Start_conversation == false {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errs.NewResponse("conversation not started", http.StatusBadRequest))
		return
	}

	fmt.Println("starting conversation .........")
	cmd := exec.Command("bash", "-c", "source "+"/home/rahma/spear-go/pkg/scripts/diart_run4.sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errs.NewResponse(errs.ErrServerErr.Error(), http.StatusInternalServerError))
		return
	}
	if err := cmd.Start(); err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errs.NewResponse(errs.ErrServerErr.Error(), http.StatusInternalServerError))
		return
	}
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		if err != nil {
			// TODO kill process here
			fmt.Println(err.Error())
			break
		}
		str := string(tmp)
		if len(str) == 1024 {
			fmt.Print("str len:", len(str), "\noutput:\n", str)
			break
		}
		/*fmt.Print("str len:", len(str), "\noutput:\n", str)
		if err != nil {
			break
		}*/
	}
	/*fmt.Println("sleeping-----------")
	cmd2 := exec.Command("sleep", "20")
	if err := cmd2.Run(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("slept")*/
	/*for true {
		//fmt.Println("out:", outb.String(), "err:", errb.String())

		if len(outb.String()) > 0 {
			fmt.Println("out:", outb.String(), "err:", errb.String())
		}
	}
	fmt.Println("out:", outb.String(), "err:", errb.String())*/
	//cmd.Wait()

	/*pgid, err := syscall.Getpgid(cmd.Process.Pid)
	if err == nil {
		fmt.Println("killing the process")
		err := syscall.Kill(-pgid, 15)
		if err != nil {
			log.Fatal("failed to kill")
		} else {
			fmt.Println("process killed")
		}
	}

	cmd.Wait()*/

	/*err := cmd.Process.Kill()

	if err != nil {
		fmt.Println("failed to kill the process")
	} else {
		fmt.Println("process killed")
	}*/
}
