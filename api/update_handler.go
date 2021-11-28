package api

import (
	"deblog-go/chain"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

func getUpdateChain(module string) (func() chain.Chain, error) {
	switch module {
	case "post":
		return chain.GetUpdatePostChain, nil
	}
	return nil, errors.New("unable to find update chain")
}

func (s *Server) updateHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "PUT" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println(r.Method, r.RequestURI)

	var body map[string]interface{}

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("unable to parse body")
		w.WriteHeader(http.StatusBadRequest)
	}

	if err := json.Unmarshal(bodyBytes, &body); err != nil {
		log.Println("unable to parse body")
		w.WriteHeader(http.StatusBadRequest)
	}

	log.Println(body)

	module := body["module"].(string)
	log.Println("module: ", module)
	if module == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := map[string]interface{}{
		"db":   s.DB,
		"data": body["data"].(map[string]interface{}),
	}
	derivedChain, err := getUpdateChain(module)
	if err != nil {
		log.Println("unable to find update chain")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	chn := derivedChain()
	if err := chn.Exec(ctx); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	data := ctx["data"]

	jsonResp, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}
