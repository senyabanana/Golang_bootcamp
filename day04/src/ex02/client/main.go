package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	keyFlag := flag.String("k", "", "Candy type abbreviation")
	countFlag := flag.Int("c", 0, "Count of candy to buy")
	moneyFlag := flag.Int("m", 0, "Amount of money given to machine")
	flag.Parse()

	// Загрузка клиентского certificate и key
	cert, err := tls.LoadX509KeyPair("../minica/client/cert.pem", "../minica/client/key.pem")
	if err != nil {
		log.Fatal("Failed to load client certificate and key:", err)
	}

	// Загрузка CA certificate
	caCert, err := os.ReadFile("../minica/minica.pem")
	if err != nil {
		log.Fatal("Failed to read CA certificate:", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Конфигурирование TLS
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	}

	// Создание HTTPS client
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	request := fmt.Sprintf(`{"money": %d, "candyType": "%s", "candyCount": %d}`, *moneyFlag, *keyFlag, *countFlag)
	resp, err := client.Post("http://localhost:3333/buy_candy", "application/json", strings.NewReader(request))
	if err != nil {
		log.Fatal("Failed to send request:", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Failed to read response:", err)
	}

	var data map[string]interface{}

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal("Failed to Unmarshal:", err)
		return
	}

	if len(data) > 1 {
		fmt.Printf("%s Your change is %v\n", data["thanks"], data["change"])
	} else {
		fmt.Println(data["error"])
	}
}
