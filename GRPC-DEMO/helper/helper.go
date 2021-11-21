package helper

import (
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
)

// GetServerCreds 服务端证书配置
func GetServerCreds() credentials.TransportCredentials {
	cert, _ := tls.LoadX509KeyPair("server.pem", "server.key")

	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("ca.pem")

	certPool.AppendCertsFromPEM(ca)

	return credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})
}

// GetClientCreds 客户端证书配置
func GetClientCreds() credentials.TransportCredentials {
	cert, _ := tls.LoadX509KeyPair("client.pem", "client.key")
	certPool := x509.NewCertPool()

	ca, _ := ioutil.ReadFile("ca.pem")

	certPool.AppendCertsFromPEM(ca)

	return credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "localhost",
		RootCAs:      certPool,
	})
}
