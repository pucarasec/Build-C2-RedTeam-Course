package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"../crypto/dh"
	"./comm/layer0"
	"./comm/layer1"
	"./models"
	"gorm.io/gorm"
)

type AppHandler struct {
	db *gorm.DB
}

func NewAppHandler(db *gorm.DB) *AppHandler {
	return &AppHandler{
		db: db,
	}
}

func (h *AppHandler) HandleAuthenticatedMsg(clientId string, msg []byte) ([]byte, error) {
	var agent models.Agent
	var messages []models.Message
	response := ""

	h.db.FirstOrCreate(&agent, models.Agent{Id: clientId})

	if msg[0] == '!' {
		message := models.Message{
			Agent: agent,
			Text:  string(msg[1:]),
		}
		h.db.Create(&message)
	}

	h.db.Where("created_at >= ?", agent.LastSeen).Order("created_at asc").Find(&messages)

	for _, msg := range messages {
		response += fmt.Sprintf("%s: %s\n", msg.AgentId, msg.Text)
	}

	if agent.FirstSeen.IsZero() {
		agent.FirstSeen = time.Now()
	}
	agent.LastSeen = time.Now()

	h.db.Save(&agent)

	return []byte(response), nil
}

func main() {
	key := []byte("some random key!")
	appHandler := NewAppHandler(models.GetDB())
	dhHandler := layer1.NewDHHandler(layer1.NewBasicKeyRespository(), dh.NewKeyExchange(), appHandler)
	encryptedHandler := layer0.NewEncryptedHandler(key, dhHandler)
	handler := layer0.NewHTTPHandler(encryptedHandler)
	http.Handle("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
