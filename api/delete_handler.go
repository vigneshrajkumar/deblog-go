package api

import (
	"deblog-go/chain"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func getDeleteChain(module string) (func() chain.Chain, error) {
	switch module {
	case "post":
		return chain.GetDeletePostChain, nil
	case "author":
		//return chain.GetGetAllAuthor, nil
	}
	return nil, errors.New("unable to find fetch chain")
}

func (s *Server) deleteHandler(w http.ResponseWriter, r *http.Request) {

	log.Println(r.Body)
	log.Println(r.RequestURI)

	module := r.URL.Query().Get("module")
	id := r.URL.Query().Get("id")
	log.Println("module: ", module)
	log.Println("id: ", id)
	if module == "" || id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := map[string]interface{}{
		"db": s.DB,
		"id": id,
	}
	derivedChain, err := getDeleteChain(module)
	if err != nil {
		log.Println("unable to find fetch chain")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	chn := derivedChain()
	if err := chn.Exec(ctx); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	dataList := ctx["status"]

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
