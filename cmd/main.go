package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/GrooveCommunity/zabbix-proxy-paygo/entity"
	"github.com/gorilla/mux"
)

var (
	projectID, topicDispatcher, topicMetrics string
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/healthy", handleValidateHealthy).Methods("GET")
	router.HandleFunc("/webhook", handleWebhook).Methods("POST")

	/*projectID = os.Getenv("PROJECT_ID")
	topicDispatcher = os.Getenv("TOPIC_ID_DISPATCHER")
	topicMetrics = os.Getenv("TOPIC_ID_METRICS")

	if projectID == "" || topicDispatcher == "" || topicMetrics == "" {
		log.Fatal("Nem todas as variáveis de ambiente requeridas foram fornecidas. ")
	}*/

	log.Fatal(http.ListenAndServe(":"+os.Getenv("APP_PORT"), router))
}

func handleValidateHealthy(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(entity.Healthy{Status: "Success!"})
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)

	log.Println(string(body))
}
