package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/charmbracelet/log"
	"github.com/fuckthinkpad/internal/services"
	"github.com/fuckthinkpad/internal/utils"
	"github.com/fuckthinkpad/internal/ws"
)

func ChannelMasterController(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		createChannel(w, r)
	}

	if r.Method == "GET" {
		joinChannel(w, r)
	}
}

func createChannel(w http.ResponseWriter, r *http.Request) {
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

	//create a Manager in the server's RAM
	ws.MasterManager.SetManager(managerName, ws.NewManager(managerName))

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "room created",
		"room-url": "GET " + r.URL.Host + "/channel?channelName=" + reqBody.ChannelName + "&password=" + reqBody.Password,
	})
}

func joinChannel(w http.ResponseWriter, r *http.Request) {

	var (
		reqQuery struct {
			ChannelName string
			Password    string
		}
	)

	reqQuery.ChannelName = r.URL.Query().Get("channelName")
	reqQuery.Password = r.URL.Query().Get("password")

	if reqQuery.Password == "" || reqQuery.ChannelName == "" {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Channel name and password is required!",
		})
		return
	}

	ch, err := services.GetChannelService(reqQuery.ChannelName)
	if err != nil {
		log.Warn("Error Querying Database", "err", err)
		return
	}

	fmt.Println("ack - ", ch)

	if time.Now().After(ch.TTL) {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "channel death due to ttl",
		})
		return
	}

	if ch.Password != reqQuery.Password {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Room Authentication failed.Disallowed",
		})
		return
	}

	//valid request,join to the server RAM pool of managers
	m := ws.MasterManager.GetManager(ch.ManagerName)
	m.ServeWS(w, r,ch.TTL)
}
