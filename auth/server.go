package auth

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/emicklei/go-restful"
	"github.com/fanux/fist/tools"
	"github.com/wonderivan/logger"
)

var (
	//AuthHTTPSPort is cmd port param
	AuthHTTPSPort uint16
	//AuthHTTPPort is cmd port param
	AuthHTTPPort uint16
	//AuthCert is cmd cert file
	AuthCert string
	//AuthKey is cmd key file
	AuthKey string
	//PrivateKey is gen keypair privateKey
	PrivateKey string
	//PublicKey is gen keypair publicKey
	PublicKey string

	//authHTTPSPortString is string of AuthHTTPSPort
	authHTTPSPortString string
	//authHTTPPortString is string of AuthHTTPPort
	authHTTPPortString string
)

//Serve start a auth server
func Serve() {
	var err error
	//var err error
	Pub, Priv, err = loadKeyPair()
	if err != nil {
		logger.Error(err)
		os.Exit(-1)
	}
	if _, err = os.Stat(AuthCert); err != nil {
		logger.Error(err)
		os.Exit(-1)
	}
	if _, err = os.Stat(AuthKey); err != nil {
		logger.Error(err)
		os.Exit(-1)
	}
	go httpServer()
	httpsServer()
}

func httpsServer() {
	wsContainer := restful.NewContainer()
	wsContainer.Router(restful.CurlyRouter{})
	auth := new(restful.WebService)
	//registry k8s auth and fist auth
	K8sRegister(auth)
	wsContainer.Add(auth)
	//process port for command
	authHTTPSPortString = ":" + strconv.FormatUint(uint64(AuthHTTPSPort), 10)
	logger.Info("start listening on localhost", authHTTPSPortString)
	server := &http.Server{Addr: authHTTPSPortString, Handler: wsContainer}
	//process cert/key for command
	logger.Info("certFile is :", AuthCert, ";keyFile is:", AuthKey)

	log.Fatal(server.ListenAndServeTLS(AuthCert, AuthKey))
}

func httpServer() {
	wsContainer := restful.NewContainer()
	wsContainer.Router(restful.CurlyRouter{})
	auth := new(restful.WebService)
	//registry k8s auth and fist auth
	TokenRegister(auth)
	wsContainer.Add(auth)
	//cors
	tools.Cors(wsContainer)
	//process port for command
	authHTTPPortString = ":" + strconv.FormatUint(uint64(AuthHTTPPort), 10)
	logger.Info("start listening on localhost", authHTTPPortString)
	server := &http.Server{Addr: authHTTPPortString, Handler: wsContainer}

	log.Fatal(server.ListenAndServe())
}
