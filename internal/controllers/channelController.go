package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/fuckthinkpad/internal/services"
	"github.com/fuckthinkpad/internal/utils"
	"github.com/fuckthinkpad/internal/ws"
)

// func ChannelMasterController(w http.ResponseWriter, r *http.Request) {

// 	if r.Method == "GET" {
// 		createChannel(w, r)
// 	}

// 	if r.Method == "GET" {
// 		getChannels(w, r)
// 	}

// }

func CreateChannel(w http.ResponseWriter, r *http.Request) {
	var (
		reqBody struct {
			ChannelName string `json:"channelName,omitempty"`
			Password    string `json:"password,omitempty"`
			OwnerSlug   string `json:"ownerSlug,omitempty"`
		}
	)

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		log.Warn("JSON Parsing error", "err", err)
		return
	}
	managerName := utils.PetNameGen()
	if err := services.CreateChannelService(reqBody, managerName); err != nil {
		log.Warn("Error Inserting in Database", "err", err)
		return
	}

	//create WS Connection
	m := ws.NewManager(managerName)
	ws.MasterManager.SetManager(managerName, m)

	//check if manager exists
	m.ServeWS(w, r)
}

func GetChannels(w http.ResponseWriter, r *http.Request) {

	// var (
	// 	reqBody struct {
	// 		ChannelName string `json:"channelName,omitempty"`
	// 		Password    string `json:"password,omitempty"`
	// 	}
	// )

	// fmt.Println(r.URL.Query())

	// if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
	// 	log.Warn("JSON Parsing error", "err", err)
	// 	return
	// }

	// ch, err := services.GetChannelService(reqBody)
	// if err != nil {
	// 	log.Warn("Error Querying Database", "err", err)
	// 	return
	// }

	// if ch.Password != reqBody.Password {
	// 	json.NewEncoder(w).Encode(map[string]interface{}{
	// 		"message": "Room Authentication failed.Disallowed",
	// 	})
	// 	return
	// }

	// join to the main server
	m := ws.MasterManager.GetManager("root")
	m.ServeWS(w, r)
}
