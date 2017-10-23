package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/cpanato/ditos_gauchos/ditos"
)

const (
	responseType = "in_channel"
	username     = "Gaucho Macho"
	responseIcon = "http://www.eev.com.br/sipat/imagens/imgFotos65446.jpg"
)

type Handler struct {
	ditos *ditos.Ditos
}

func New(ditos *ditos.Ditos) *Handler {
	return &Handler{
		ditos: ditos,
	}
}

type message struct {
	ResponseType string `json:"response_type"`
	Text         string `json:"text"`
	Username     string `json:"username"`
	IconUrl      string `json:"icon_url"`
}

func (h *Handler) HandleBah(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	bah := &message{
		ResponseType: responseType,
		Text:         h.ditos.Random(),
		Username:     username,
		IconUrl:      responseIcon,
	}

	if err := json.NewEncoder(res).Encode(bah); err != nil {
		errMsg := fmt.Sprintf("failed to encode response: %v", err)
		log.Printf("%s\n", err)

		http.Error(res, fmt.Sprintf(`{"error": %s}`, errMsg), http.StatusInternalServerError)
		return
	}
}
