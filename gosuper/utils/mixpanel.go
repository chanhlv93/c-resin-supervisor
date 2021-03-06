package utils

import (
	"log"

	"github.com/resin-io/resin-supervisor/gosuper/Godeps/_workspace/src/github.com/dukex/mixpanel"
	"github.com/resin-io/resin-supervisor/gosuper/config"
)

var client *mixpanel.Mixpanel
var username string
var distinctId string

func MixpanelInit(token string) (err error) {
	client = mixpanel.NewMixpanel(token)

	if userConfig, err := config.ReadConfig(config.DefaultConfigPath); err == nil {
		username = userConfig.Username
	}
	return
}

func MixpanelSetId(id string) {
	distinctId = id
}

func MixpanelTrack(eventName string, properties map[string]interface{}) (err error) {
	if properties == nil {
		properties = make(map[string]interface{})
	}
	properties["username"] = username
	properties["uuid"] = distinctId
	log.Printf("Event: %s %v", eventName, properties)
	_, err = client.Track(distinctId, eventName, properties)
	return
}
