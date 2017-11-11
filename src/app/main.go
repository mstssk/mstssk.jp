package app

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"strconv"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
)

func init() {
	http.HandleFunc("/api/admin/health", handleAPIHealth)
	http.HandleFunc("/api/sleep", handleAPISleepTest)
	http.HandleFunc("/api/admin/testout", handleOutboundReqTest)
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is dummy.")
}

func handleAPISleepTest(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	log.Debugf(c, "%#v", c)

	formSec, _ := strconv.ParseInt(r.FormValue("sec"), 0, 0)
	sec := time.Duration(formSec) * time.Second
	time.Sleep(sec)
	fmt.Fprintf(w, "sleep in %s", sec)
}

func handleOutboundReqTest(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	log.Debugf(c, "%#v", c)

	sec := r.FormValue("sec")
	if sec == "" {
		sec = "0"
	}

	client := makeClient(c)
	for i := 0; i < 2; i++ {
		resp, err := httpGet(c, client, "https://mstssk.jp/api/sleep?sec="+sec)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err.Error())
			return
		}
		fmt.Fprintf(w, "%d get: %s\n", i, resp)
	}
}

func makeClient(c context.Context) *http.Client {
	c, _ = context.WithTimeout(c, 10*time.Second)
	return urlfetch.Client(c)
}

func httpGet(c context.Context, client *http.Client, url string) ([]byte, error) {
	// c, _ = context.WithTimeout(c, 10*time.Second)
	// c, cancel := context.WithTimeout(c, 10*time.Second)
	// defer cancel()
	// client := urlfetch.Client(c)
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body, nil
}

func handleAPIHealth(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	log.Debugf(c, "%#v", c)

	w.Header().Set("Content-Type", "application/json")
	res := map[string]string{
		"version":                  runtime.Version(),
		"app_id":                   appengine.AppID(c),
		"module_name":              appengine.ModuleName(c),
		"version_id":               appengine.VersionID(c),
		"datacenter":               appengine.Datacenter(c),
		"default_version_hostname": appengine.DefaultVersionHostname(c),
		"instance_id":              appengine.InstanceID(),
		"server_software":          appengine.ServerSoftware(),
	}
	json, _ := json.Marshal(res)
	fmt.Fprint(w, string(json))
}
