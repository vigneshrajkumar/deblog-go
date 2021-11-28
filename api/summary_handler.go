package api

import (
	"deblog-go/chain"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func getSummaryChain(module string) (func() chain.Chain, error) {
	switch module {
	case "post":
		return chain.GetSummaryPostChain, nil
	}
	return nil, errors.New("unable to find fetch chain")
}

func (s *Server) summaryHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println(r.Method, r.RequestURI)

	module := r.URL.Query().Get("module")
	id := r.URL.Query().Get("id")
	log.Println("module: ", module)
	if module == "" || id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := map[string]interface{}{
		"db": s.DB,
		"id": id,
	}
	derivedChain, err := getSummaryChain(module)
	if err != nil {
		log.Println("unable to find summary chain")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	chn := derivedChain()
	if err := chn.Exec(ctx); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	dataList := ctx["summary"]

	jsonResp, err := json.Marshal(dataList)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}
