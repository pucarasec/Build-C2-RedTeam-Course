package admin

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"../models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type AdminHandler struct {
	db *gorm.DB
}

func (h *AdminHandler) rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello admin!")
}

func (h *AdminHandler) agentsHandler(w http.ResponseWriter, r *http.Request) {
	var agents []models.Agent
	tx := h.db.Find(&agents)
	if tx.Error != nil {
		errorString := fmt.Sprintf("Error getting agents: %s", tx.Error)
		http.Error(w, errorString, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(agents)
}

func (h *AdminHandler) agentHandler(w http.ResponseWriter, r *http.Request) {
	var agent models.Agent
	vars := mux.Vars(r)
	tx := h.db.First(&agent, "id = ?", vars["id"])
	if tx.Error == nil {
		json.NewEncoder(w).Encode(agent)
	} else if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		http.NotFound(w, r)
	} else {
		errorString := fmt.Sprintf("Error getting agents: %s", tx.Error)
		http.Error(w, errorString, http.StatusInternalServerError)
	}
}

func (h *AdminHandler) agentCommandsHandler(w http.ResponseWriter, r *http.Request) {
	var commands []models.Command
	vars := mux.Vars(r)
	tx := h.db.Find(&commands, "agent_id = ?", vars["agentId"])
	if tx.Error != nil {
		errorString := fmt.Sprintf("Error getting agent commands: %s", tx.Error)
		http.Error(w, errorString, http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(commands)
}

func (h *AdminHandler) postAgentCommandHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var args []string
	agentId := vars["agentId"]
	err := json.NewDecoder(r.Body).Decode(&args)
	if err != nil {
		errorString := fmt.Sprintf("Error decoding json: %s", err)
		http.Error(w, errorString, http.StatusInternalServerError)
		return
	}

	argsBytes, _ := json.Marshal(args)
	command := models.Command{
		AgentId: agentId,
		Args:    argsBytes,
	}
	tx := h.db.Create(&command)
	if tx.Error != nil {
		errorString := fmt.Sprintf("Error creating command: %s", tx.Error)
		http.Error(w, errorString, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&command)
}

func (h *AdminHandler) agentCommandResultsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var commandResults []models.CommandResult
	agentId := vars["agentId"]
	commandId, _ := strconv.ParseUint(vars["commandId"], 10, 64)

	tx := h.db.Find(&commandResults, &models.CommandResult{
		AgentId:   agentId,
		CommandId: uint(commandId),
	})
	if tx.Error != nil {
		errorString := fmt.Sprintf("Error getting agent command results: %s", tx.Error)
		http.Error(w, errorString, http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&commandResults)
}

func NewAdminHandler(prefix string, db *gorm.DB) *mux.Router {
	r := mux.NewRouter().PathPrefix(prefix).Subrouter()
	h := AdminHandler{db: db}
	r.HandleFunc("/", h.rootHandler).Methods("GET")
	r.HandleFunc("/agents/", h.agentsHandler).Methods("GET")
	r.HandleFunc("/agents/{id}", h.agentHandler).Methods("GET")
	r.HandleFunc("/agents/{agentId}/commands", h.agentCommandsHandler).
		Methods("GET")
	r.HandleFunc("/agents/{agentId}/commands", h.postAgentCommandHandler).
		Methods("POST")
	r.HandleFunc("/agents/{agentId}/commands/{commandId}/results", h.agentCommandResultsHandler).
		Methods("GET")

	return r
}
