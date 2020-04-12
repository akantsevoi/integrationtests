package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/akantsevoi/integrationtests/cmd/database"
	"github.com/julienschmidt/httprouter"
)

// Server _
type Server struct {
	Router  *httprouter.Router
	storage database.Storage
}

// NewServer _
func NewServer(db database.Options) (Server, error) {
	var s Server

	router := httprouter.New()

	router.GET("/read", s.read)
	router.POST("/write", s.write)

	s.Router = router
	storage, err := database.NewStorage(db)
	if err != nil {
		return Server{}, err
	}
	s.storage = storage

	return s, nil
}

func (s *Server) write(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var m struct{ Message string }
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		responseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, err := s.storage.AddMessage(m.Message); err != nil {
		responseError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responeJSON(w, map[string]string{"success": "ok"})
}

func (s *Server) read(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	messages, err := s.storage.GetMessages()
	if err != nil {
		responseError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responeJSON(w, messages)
}

func responeJSON(w http.ResponseWriter, item interface{}) {
	json, err := json.Marshal(item)
	if err != nil {
		responseError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(json)
}

func responseError(w http.ResponseWriter, message string, code int) {
	http.Error(w, fmt.Sprintf(`{"error": "%v"}`, message), code)
}
