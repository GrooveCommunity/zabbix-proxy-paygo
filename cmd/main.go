package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/GrooveCommunity/glib-noc-event-structs/entity"
	"github.com/GrooveCommunity/zabbix-proxy-paygo/internal"
	"github.com/gorilla/mux"
)

var (
	projectID, topicDispatcher, topicMetrics string
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/healthy", handleValidateHealthy).Methods("GET")
	router.HandleFunc("/webhook", handleWebhook).Methods("POST")

	projectID = os.Getenv("PROJECT_ID")
	topicDispatcher = os.Getenv("TOPIC_ID_DISPATCHER_ZABBIX")

	if projectID == "" || topicDispatcher == "" || topicMetrics == "" {
		log.Fatal("Nem todas as variáveis de ambiente requeridas foram fornecidas. ")
	}

	log.Fatal(http.ListenAndServe(":"+os.Getenv("APP_PORT"), router))
}

func handleValidateHealthy(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(entity.Healthy{Status: "Success!"})
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	var zabbixRequest entity.ZabbixRequest

	body, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(body, &zabbixRequest)

	log.Println(string(body))
	log.Println("=========================================\n\n\n")

	internal.ForwardIssue(zabbixRequest, body, projectID, topicDispatcher)
}
