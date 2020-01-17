package main

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

/*
Order - a record of an order.
*/
type Order struct {
	ID        string `xml:"id" json:"id"`
	Data      string `xml:"data" json:"data"`
	CreatedAt string `xml:"createdAt" json:"createdAt"`
	UpdatedAt string `xml:"updatedAt" json:"updatedAt"`
}

/*
Request - a listing of Orders
*/
type Request struct {
	Orders []Order `xml:"order" json:"orderList"`
}

/*
Process - a worker to transform an Order
*/
func Process(index int, orders []Order, waitGroup *sync.WaitGroup) {
	log.Printf("Processing order[%d]\t%s\n", index, orders[index].ID)
	defer waitGroup.Done()
	orders[index].Data = strings.ToUpper(orders[index].Data)
}

func processRequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST supported", http.StatusMethodNotAllowed)
		return
	}

	request := Request{}
	var data, err = ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = xml.Unmarshal(data, &request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var waitGroup sync.WaitGroup
	for index := range request.Orders {
		waitGroup.Add(1)
		go Process(index, request.Orders, &waitGroup)
	}
	waitGroup.Wait()

	jsonString, err := json.MarshalIndent(request, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonString)
}

func main() {
	http.HandleFunc("/process", processRequestHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
