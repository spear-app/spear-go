package testing

import (
	"github.com/gorilla/mux"
	"github.com/spear-app/spear-go/pkg/handlers"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAudioForwarding(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/api/audio/recorded_audio", handlers.RecordedAudio).Methods(http.MethodPost)
	w := httptest.NewRecorder()

	expected := `[{
    "id": 14,
    "user_id": 1065,
    "created_at": "0001-01-01T00:00:00Z",
    "updated_at": "0001-01-01T00:00:00Z",
    "deleted_at": null,
    "title": "knock",
    "body": "there's knock sound near you"
}]`
	assert.Equal(t, []byte(expected), w.Body.Bytes())
}
