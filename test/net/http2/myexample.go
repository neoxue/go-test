package main

import (
	"crypto/tls"
	"flag"
	"log"
	"net"
	"net/http"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"time"

	"os"

	"golang.org/x/crypto/acme/autocert"
	"golang.org/x/net/http2"
)

var (
	prod = flag.Bool("prod", false, "Whether to configure itself to be the production http2.golang.org server.")

	httpsAddr = flag.String("https_addr", "localhost:4430", "TLS address to listen on ('host:port' or ':port'). Required.")
	httpAddr  = flag.String("http_addr", "", "Plain HTTP address to listen on ('host:port', or ':port'). Empty means no HTTP.")

	hostHTTP  = flag.String("http_host", "", "Optional host or host:port to use for http:// links to this service. By default, this is implied from -http_addr.")
	hostHTTPS = flag.String("https_host", "", "Optional host or host:port to use for http:// links to this service. By default, this is implied from -https_addr.")
)

const (
	idleTimeout   = 5 * time.Minute
	activeTimeout = 10 * time.Minute
)

// TODO: put this into the standard library and actually send
// PING frames and GOAWAY, etc: golang.org/issue/14204
func idleTimeoutHook() func(net.Conn, http.ConnState) {
	var mu sync.Mutex
	m := map[net.Conn]*time.Timer{}
	return func(c net.Conn, cs http.ConnState) {
		mu.Lock()
		defer mu.Unlock()
		if t, ok := m[c]; ok {
			delete(m, c)
			t.Stop()
		}
		var d time.Duration
		switch cs {
		case http.StateNew, http.StateIdle:
			d = idleTimeout
		case http.StateActive:
			d = activeTimeout
		default:
			return
		}
		m[c] = time.AfterFunc(d, func() {
			log.Printf("closing idle conn %v after %v", c.RemoteAddr(), d)
			go c.Close()
		})
	}
}

func registerHandlers() {
	tiles := newGopherTilesHandler()
	push := newPushHandler()

	mux2 := http.NewServeMux()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/gophertiles":
			tiles.ServeHTTP(w, r) // allow HTTP/2 + HTTP/1.x
			return
		case strings.HasPrefix(r.URL.Path, "/serverpush"):
			push.ServeHTTP(w, r) // allow HTTP/2 + HTTP/1.x
			return
		case r.TLS == nil: // do not allow HTTP/1.x for anything else
			http.Redirect(w, r, "https://"+httpsHost()+"/", http.StatusFound)
			return
		}
		if r.ProtoMajor == 1 {
			if r.URL.Path == "/reqinfo" {
				reqInfoHandler(w, r)
				return
			}
			homeOldHTTP(w, r)
			return
		}
		mux2.ServeHTTP(w, r)
	})
	mux2.HandleFunc("/", home)
	mux2.Handle("/file/gopher.png", fileServer("https://golang.org/doc/gopher/frontpage.png", 0))
	mux2.Handle("/file/go.src.tar.gz", fileServer("https://storage.googleapis.com/golang/go1.4.1.src.tar.gz", 0))
	mux2.HandleFunc("/reqinfo", reqInfoHandler)
	mux2.HandleFunc("/crc32", crcHandler)
	mux2.HandleFunc("/ECHO", echoCapitalHandler)
	mux2.HandleFunc("/clockstream", clockStreamHandler)
	mux2.Handle("/gophertiles", tiles)
	mux2.HandleFunc("/redirect", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusFound)
	})
	stripHomedir := regexp.MustCompile(`/(Users|home)/\w+`)
	mux2.HandleFunc("/goroutines", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		buf := make([]byte, 2<<20)
		w.Write(stripHomedir.ReplaceAll(buf[:runtime.Stack(buf, true)], nil))
	})
}

func main() {
	var srv http.Server
	flag.BoolVar(&http2.VerboseLogs, "verbose", false, "Verbose HTTP/2 debugging.")
	flag.Parse()

	srv.Addr = *httpsAddr
	srv.ConnState = idleTimeoutHook()

	registerHandlers()

	if *prod {
		*hostHTTP = "http2.golang.org"
		*hostHTTPS = "http2.golang.org"
		log.Fatal(serveProd())
	}
}

func serveProd() error {
	errc := make(chan error, 2)
	go func() { errc <- http.ListenAndServe(":80", nil) }()
	go func() { errc <- serveProdTLS() }()
	return <-errc
}

func serveProdTLS() error {
	const cacheDir = "/var/cache/autocert"
	if err := os.MkdirAll(cacheDir, 0700); err != nil {
		return err
	}
	m := autocert.Manager{
		Cache:      autocert.DirCache(cacheDir),
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("http2.golang.org"),
	}
	srv := &http.Server{
		TLSConfig: &tls.Config{
			GetCertificate: m.GetCertificate,
		},
	}
	http2.ConfigureServer(srv, &http2.Server{
		NewWriteScheduler: func() http2.WriteScheduler {
			return http2.NewPriorityWriteScheduler(nil)
		},
	})
	ln, err := net.Listen("tcp", ":443")
	if err != nil {
		return err
	}
	return srv.Serve(tls.NewListener(tcpKeepAliveListener{ln.(*net.TCPListener)}, srv.TLSConfig))
}

type tcpKeepAliveListener struct {
	*net.TCPListener
}
