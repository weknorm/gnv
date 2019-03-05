package remote

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/yinqiwen/gotoolkit/ots"
	"github.com/yinqiwen/gsnova/common/channel"
	httpChannel "github.com/yinqiwen/gsnova/common/channel/http"
	"github.com/yinqiwen/gsnova/common/channel/websocket"
	"github.com/yinqiwen/gsnova/common/logger"
)

// hello world, the web server
func indexCallback(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, strings.Replace(html, "${Version}", channel.Version, -1))
}

func statCallback(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintf(w, "Version:    %s\n", channel.Version)
	ots.Handle("stat", w)
}

func stackdumpCallback(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	ots.Handle("stackdump", w)
}

func startHTTPProxyServer(listenAddr string, cert, key string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexCallback)
	mux.HandleFunc("/stat", statCallback)
	mux.HandleFunc("/stackdump", stackdumpCallback)
	mux.HandleFunc("/ws", websocket.WebsocketInvoke)
	mux.HandleFunc("/http/pull", httpChannel.HTTPInvoke)
	mux.HandleFunc("/http/push", httpChannel.HTTPInvoke)
	mux.HandleFunc("/http/test", httpChannel.HttpTest)

	logger.Info("Listen on HTTP address:%s", listenAddr)
	var err error
	if len(cert) == 0 {
		err = http.ListenAndServe(listenAddr, mux)
	} else {
		err = http.ListenAndServeTLS(listenAddr, cert, key, mux)
	}

	if nil != err {
		logger.Error("Listen HTTP server error:%v", err)
	}
}

const html = `
Under construction...<br/>
Coming back soon...<br/>
Waiting for update...<br/>
`
