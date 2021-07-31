package internal

import (
	"log"

	"encoding/json"
	"time"

	//	"github.com/GrooveCommunity/glib-cloud-storage/gcp"
	"github.com/GrooveCommunity/glib-noc-event-structs/entity"
)

func ForwardIssue(zabbixRequest entity.ZabbixRequest, body []byte, projectID, topicMetrics string) {

	zabbixRequest.Fields.RequestDate = time.Now().Format(time.RFC3339)

	payload, errPayLoad := json.Marshal(zabbixRequest)

	if errPayLoad != nil {
		log.Fatal(entity.ResponseError{
			Message:    "Erro na conversão do payload para JSON",
			StatusCode: 500,
			Error:      errPayLoad,
		})
	}

	log.Println(string(payload))

	//go gcp.PublicMessage(projectID, topicDispatcher, payload)
}

/*func ValidateEvent(zabbixRequest entity.ZabbixRequest) {
	log.Println(zabbixRequest)
}*/
