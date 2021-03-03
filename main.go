package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/namsral/flag"
	"github.com/tidwall/pretty"
)

type Alert struct {
	Data struct {
		Alerts []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
			GID  string `json:"group_id"`
		} `json:"alerts"`
	} `json:"data"`
}

var (
	vmalertHost   = flag.String("host", "localhost", "Host where VMAlert responds")
	vmalertSchema = flag.String("schema", "http", "Use http|https")
	vmalertPort   = flag.Int("port", 8880, "VMAlert port")
	vmalertAction = flag.String("action", "groups", "VMAlert action to take {groups|alerts|metrics|reload|status <groupName> <alertName>}")
	prettyPrint   = flag.Bool("pretty", false, "Pretty print {false|true}")
)

func init() {
	flag.Parse()
}

func ErrCheck(e error) {
	if e != nil {
		log.Println(e)
	}
}

func GetJSONData(apiBase string, apiEndpoint string) []byte {
	response, err := http.Get(apiBase + apiEndpoint)
	ErrCheck(err)
	data, _ := ioutil.ReadAll(response.Body)

	if *prettyPrint {
		return pretty.Pretty(data)
	} else {
		return data
	}
}

func NameToID(itemName, itemType, apiBase string) string {
	var alertResponse Alert
	var itemID string

	response, err := http.Get(apiBase + "/api/v1/alerts")
	ErrCheck(err)
	data, _ := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(data, &alertResponse)
	ErrCheck(err)

	switch operation := itemType; operation {
	case "alert":
		for i := 0; i < len(alertResponse.Data.Alerts); i++ {
			if itemName == alertResponse.Data.Alerts[i].Name {
				itemID = alertResponse.Data.Alerts[i].ID
			} else {
				itemID = ""
			}
		}
	case "group":
		for i := 0; i < len(alertResponse.Data.Alerts); i++ {
			if itemName == alertResponse.Data.Alerts[i].Name {
				itemID = alertResponse.Data.Alerts[i].GID
			} else {
				itemID = ""
			}
		}
	}
	return itemID
}

func main() {
	host := *vmalertHost
	action := *vmalertAction
	schema := *vmalertSchema

	vmalertBase := schema + "://" + host + ":" + strconv.Itoa(*vmalertPort)

	switch takeAction := action; takeAction {
	case "groups":
		endpoint := "/api/v1/groups"
		fmt.Println(string(GetJSONData(vmalertBase, endpoint)))
	case "alerts":
		endpoint := "/api/v1/alerts"
		fmt.Println(string(GetJSONData(vmalertBase, endpoint)))
	case "metrics":
		endpoint := "/metrics"
		fmt.Println(string(GetJSONData(vmalertBase, endpoint)))
	case "reload":
		endpoint := "/-/reload"
		fmt.Println(string(GetJSONData(vmalertBase, endpoint)))
	case "status":
		alertName := os.Args[len(os.Args)-1]
		endpoint := "/api/v1/" + NameToID(alertName, "group", vmalertBase) +
			"/" + NameToID(alertName, "alert", vmalertBase) + "/status"
		fmt.Println(string(GetJSONData(vmalertBase, endpoint)))
	}
}
