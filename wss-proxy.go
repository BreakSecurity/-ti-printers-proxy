package main

import (
	"flag"
	"log"
  "net"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	// ReadBufferSize:  1024,
	// WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // NOT a good practice

var (
	toURL        = flag.String("to", "http://127.0.0.1:80", "the address and port for which to proxy requests to")
	fromURL      = flag.String("from", "127.0.0.1:4430", "the tcp address and port this proxy should listen for requests on")
	certFile     = flag.String("cert", "", "path to a tls certificate file")
	keyFile      = flag.String("key", "", "path to a private key file")
)

const (
	HTTPSPrefix     = "https://"
	HTTPPrefix      = "http://"
)

func handler(w http.ResponseWriter, r *http.Request) {

    println("Connection")

    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
      log.Println(err)
      return
    }

    _, p, err := ws.ReadMessage()
    if err != nil {
      log.Println(err)
      return
    }

    println("Printing",*toURL)

    tcpAddr, err := net.ResolveTCPAddr("tcp", *toURL)
    if err != nil {
      log.Println("ResolveTCPAddr failed:", err.Error())
      return
    }

    connTCP, err := net.DialTCP("tcp", nil, tcpAddr)
    if err != nil {
      log.Println("Dial failed:", err.Error())
      return
    }

    _, err = connTCP.Write([]byte(p))
    if err != nil {
        log.Println("Write to server failed:", err.Error())
        return
    }

    connTCP.Close()

		println("Print done!",*toURL)
}

func main() {
  flag.Parse()
  validCertFile := *certFile != ""
	validKeyFile := *keyFile != ""

	if (!validCertFile || !validKeyFile){
		log.Fatal("No existing cert or key specified")
  }

  log.Printf("Proxying calls from https://%s (SSL/TLS) to %s", *fromURL, *toURL)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServeTLS(*fromURL, *certFile, *keyFile, nil))
}
