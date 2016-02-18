package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (a *Api) scaleContainer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	vars := mux.Vars(r)
	containerId := vars["id"]
	n := r.URL.Query()["n"]

	if len(n) == 0 {
		http.Error(w, "你必须输入一个数字 (param: n)", http.StatusBadRequest)
		return
	}

	numInstances, err := strconv.Atoi(n[0])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if numInstances <= 0 {
		http.Error(w, "你必须输入一个有效的值", http.StatusBadRequest)
		return
	}

	result := a.manager.ScaleContainer(containerId, numInstances)
	// If we received any errors, continue to write result to the writer, but return a 500
	if len(result.Errors) > 0 {
		w.WriteHeader(http.StatusInternalServerError)
	}
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
