// Package classification of Logger API
//
// Documentation for OpenStack VM Logger API
//
//  Schemes: http
//  BasePath: /
//  Version: 1.0.0
//
//  Consumes:
//  - application/json
//
//  Produces:
//  - application/json
// swagger:meta
package handlers

import (
	"log"
	"net/http"

	"github.com/JungBin-Eom/OpenStack-Logger/data"
)

// A list of logs returns in the response
// swagger:response logsResponse
type logsResponseWrapper struct {
	// All logs in the system
	// in: body
	Body []data.Log
}

type MyLogs struct {
	l *log.Logger
}

func NewLogs(l *log.Logger) *MyLogs {
	return &MyLogs{l}
}

func (m *MyLogs) IndexHandler(rw http.ResponseWriter, r *http.Request) {
	http.Redirect(rw, r, "/index.html", http.StatusTemporaryRedirect)
}

// swagger:route GET / logs listLogs
// Retruns a list of Logs
// responses:
//  200: logsResponse
func (m *MyLogs) GetLogs(rw http.ResponseWriter, r *http.Request) {
	m.l.Println("Handle GET Logs")

	logs := data.GetLogs()
	err := logs.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (m *MyLogs) LogSync(rw http.ResponseWriter, r *http.Request) {
	m.l.Println("Handle Log Synchronizing")

	// curl -XGET '15.164.210.67:9200/neutron-2021.06.29/_search?q=log_level:ERROR&pretty&filter_path=hits.hits._source.logmessage'

	req, err := http.NewRequest("GET", "15.164.210.67:9200/neutron-2021.06.29/_search", nil)
	if err != nil {
		http.Error(rw, "Unable to get logs", http.StatusInternalServerError)
	}

	res, err := http.DefaultClient.Do(req)
	m.l.Println(res)
	defer res.Body.Close()
}
