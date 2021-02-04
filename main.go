package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/namsral/flag"
	"github.com/tidwall/pretty"
)

var (
	vmalertHost   = flag.String("host", "localhost", "Host where VMAlert responds")
	vmalertSchema = flag.String("schema", "http", "Use http|https")
	vmalertPort   = flag.Int("port", 8880, "VMAlert port")
	vmalertAction = flag.String("action", "groups", "VMAlert action to take {groups|alerts|metrics|reload|status <groupName> <alertId>}")
	prettyPrint   = flag.Bool("pretty", false, "Pretty print {false|true}")
)

func init() {
	flag.Parse()
}

func getJsonData(apiBase string, apiEndpoint string) []byte {
	response, err := http.Get(apiBase + apiEndpoint)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)

	if *prettyPrint {
		return pretty.Pretty(data)
	} else {
		return data
	}
}

func main() {
	host := *vmalertHost
	action := *vmalertAction
	schema := *vmalertSchema

	vmalertBase := schema + "://" + host + ":" + strconv.Itoa(*vmalertPort)

	switch takeAction := action; takeAction {
	case "groups":
		endpoint := "/api/v1/groups"
		fmt.Println(string(getJsonData(vmalertBase, endpoint)))
	case "alerts":
		endpoint := "/api/v1/alerts"
		fmt.Println(string(getJsonData(vmalertBase, endpoint)))
	case "metrics":
		endpoint := "/metrics"
		fmt.Println(string(getJsonData(vmalertBase, endpoint)))
	case "reload":
		endpoint := "/-/reload"
		fmt.Println(string(getJsonData(vmalertBase, endpoint)))
	case "status":
		groupName := os.Args[len(os.Args)-2]
		alertID := os.Args[len(os.Args)-1]
		endpoint := "/api/v1/" + groupName + "/" + alertID + "/status"
		fmt.Println(string(getJsonData(vmalertBase, endpoint)))
	}
}
