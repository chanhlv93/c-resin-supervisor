package cliclient

import (
	"log"
	"time"
	"encoding/json"
	"strconv"
	"github.com/go-resty/resty"
	"github.com/resin-io/resin-supervisor/gosuper/supermodels"
)

type Client struct {
	BaseApiEndpoint string
	ApiKey 		string
}

/*func NewClient(apiEndpoint, apiKey string) (client *Client) {
	client.BaseApiEndpoint = apiEndpoint
	client.ApiKey = apiKey
	return
}*/

type DeviveRegister struct {
	Id 		int		`json:"Id,omitempty"`
	Appid 		int 		`json:"appid"`
	Name 		string 		`json:"name"`
	Uuid 		string 		`json:"uuid"`
	Devicetype 	string 		`json:"devicetype"`
}

type DeviceState struct {
	AppId 		int 	`json:"appId"`
	DeviceId 	int 	`json:"deviceId"`
	State 		string  `json:"state"`
}

func (client *Client) CheckConnectivity() (check bool, err error) {
	resp, err := resty.R().Get(client.BaseApiEndpoint + "/v1/ping")
	if check, err = strconv.ParseBool(resp.String()); err == nil {
		return
	}
	return
}

func (client *Client) GetApps(uuid, registryEndpoint, deviceId string) (apps []supermodels.App, err error) {
	return 
}

func (client *Client) Getapplication() (apps []supermodels.App, err error) {
	resp, err := resty.R().
		//SetQueryString("apikey=" + client.ApiKey).
		SetHeader("Accept", "application/json").
		Get(client.BaseApiEndpoint + "/v1/app")
	if err != nil {
		log.Println(err)
	}
	log.Println(resp.Body())
	if err := json.Unmarshal(resp.Body(), &apps); err != nil {
		log.Println(err)
	}

	return
}

func (client *Client) RegisterDevice(devRegister DeviveRegister) (registeredAt int, deviceId int, err error) {
	resp, err := resty.R().
		SetQueryString("apikey=" + client.ApiKey).
		SetHeader("Content-Type", "application/json").
		/*SetHeader("Accept", "application/json").*/
		SetBody(devRegister).

		Post(client.BaseApiEndpoint + "/v1/device")

	if err != nil {
		log.Println(err)
	}
	//log.Println(resp.Body())

	var deviceRegistered DeviveRegister
	//registeredAtFloat64 := float64(int32(time.Now().Unix()))
	registeredAt = int(time.Now().Unix())
	if err := json.Unmarshal(resp.Body(), &deviceRegistered); err != nil {
		log.Println(err)
	}

	deviceId = deviceRegistered.Id

	return
}

func (client * Client) UpdateState(appid, deviceid int, status string) (err error) {
	devState := DeviceState{AppId: appid, DeviceId: deviceid, State: status}

	_, err = resty.R().
		SetQueryString("apikey=" + client.ApiKey).
		SetHeader("Content-Type", "application/json").
		SetBody(devState).
		Post(client.BaseApiEndpoint + "/v1/device/updatestate")
	return err
}
