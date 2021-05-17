package http_server

import (
	"log"
	"net/http"
	"strings"

	"gitee.com/kelvins-io/common/env"
	"github.com/yongcycchen/mall-users/proto"
)

// SwaggerHandler
func SwaggerHandler(w http.ResponseWriter, r *http.Request) {
	if env.IsProdMode() {
		return
	}
	if !strings.HasSuffix(r.URL.Path, "swagger.json") {
		log.Printf("Not Found: %s", r.URL.Path)
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,api_key,Authorization")
	w.Header().Set("content-type", "application/json")

	path := strings.TrimPrefix(r.URL.Path, "/swagger/")
	bytes, err := proto.Asset(path)
	if err == nil {
		w.Write(bytes)
	}
}