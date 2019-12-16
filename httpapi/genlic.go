package httpapi

import (
	"sync"
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
	"log"
	"github.com/MeloQi/license/license"
)

type GenLicHttpApi struct {
	httpaddr string
}

var genLicHttpApiInst *GenLicHttpApi = nil
var genLicHttpApiOnce sync.Once

func GetGenLicHttpApiInst(httpaddr string) *GenLicHttpApi {
	genLicHttpApiOnce.Do(func() {
		if len(httpaddr) == 0 {
			httpaddr = ":8010"
		}
		genLicHttpApiInst = &GenLicHttpApi{httpaddr: httpaddr}
	})
	return genLicHttpApiInst
}

func (s *GenLicHttpApi) Start() {
	go func() {
		api := rest.NewApi()
		api.Use(rest.DefaultDevStack...)
		router, err := rest.MakeRouter(
			rest.Post("/lic/getlic/:key", s.Post),
		)
		if err != nil {
			panic(err)
		}
		api.SetApp(router)
		log.Fatal(http.ListenAndServe(s.httpaddr, api.MakeHandler()))
	}()
}

func (s *GenLicHttpApi) Post(w rest.ResponseWriter, r *rest.Request) {
	req := &license.LicInfo{}
	if err := r.DecodeJsonPayload(req); err != nil {
		log.Print("json 解析错误：", err)
		w.Header().Set("Content-Type", "text/plain")
		w.(http.ResponseWriter).WriteHeader(400)
		w.(http.ResponseWriter).Write([]byte("json 解析错误"))
		return
	}
	key := r.PathParam("key")
	licStr, err := license.GenLic(req, key)
	if err != nil {
		log.Print("生成license错误：", err)
		w.Header().Set("Content-Type", "text/plain")
		w.(http.ResponseWriter).WriteHeader(400)
		w.(http.ResponseWriter).Write([]byte("生成license错误"))
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	w.(http.ResponseWriter).Write([]byte(licStr))
}
