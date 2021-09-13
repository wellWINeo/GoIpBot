package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"

	"github.com/wellWINeo/GoIpBot"
)

func (h *WebHandler) errorResponse(w http.ResponseWriter, code int, err error) {
	GoIpBot.Log("api.go").Error(err)
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	resp, _ := json.Marshal(map[string]string{"msg": err.Error()})
	w.Write(resp)
}

func (h *WebHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	GoIpBot.Log("api.go").Info("get all users")
	users, err := h.services.Admin.GetUsers()
	if err != nil {
		h.errorResponse(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp, _ := json.Marshal(users)
	w.Write(resp)
}

func (h *WebHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	GoIpBot.Log("api.go").Info("get user by id")
	query, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		h.errorResponse(w, http.StatusBadRequest, err)
		return
	} else if query["id"] == nil {
		h.errorResponse(w, http.StatusBadRequest, errors.New("nil id variable"))
		return
	}

	id, err := strconv.Atoi(query["id"][0])
	if err != nil {
		h.errorResponse(w, http.StatusBadRequest, err)
		return
	}

	user, err := h.services.GetUser(id)
	if err != nil {
		h.errorResponse(w, http.StatusInternalServerError, err)
		return
	}

	resp, err := json.Marshal(user)
	if err != nil {
		h.errorResponse(w, http.StatusInternalServerError, err)
		return
	}

	w.Write(resp)
}

func (h *WebHandler) GetHistoryByTg(w http.ResponseWriter, r *http.Request) {
	GoIpBot.Log("api.go").Info("get user's history by telegram tag")
	query, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		h.errorResponse(w, http.StatusBadRequest, err)
		return
	} else if query["tag"] == nil {
		h.errorResponse(w, http.StatusBadRequest, errors.New("nil id variable"))
		return
	}

	history, err := h.services.GetUserHistory(query["tag"][0])
	if err != nil {
		h.errorResponse(w, http.StatusInternalServerError, err)
		return
	}

	resp, err := json.Marshal(history)
	if err != nil {
		h.errorResponse(w, http.StatusInternalServerError, err)
		return
	}

	w.Write(resp)
}

func (h *WebHandler) RemoveHistory(w http.ResponseWriter, r *http.Request) {
	GoIpBot.Log("api.go").Info("remove history record by ID")
	query, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		h.errorResponse(w, http.StatusBadRequest, err)
		return
	} else if query["id"] == nil {
		h.errorResponse(w, http.StatusBadRequest, errors.New("nil id variable"))
		return
	}

	id, err := strconv.Atoi(query["id"][0])
	if err != nil {
		h.errorResponse(w, http.StatusBadRequest, err)
		return
	}

	err = h.services.RemoveHistory(id)
	if err != nil {
		h.errorResponse(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
