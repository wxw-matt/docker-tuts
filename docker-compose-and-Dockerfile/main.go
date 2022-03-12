package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/large-scale-deployment/coordinator/models"

	"github.com/dexterdarwich/ws-compare/handler"
	"github.com/dexterdarwich/ws-compare/service"
	"github.com/labstack/echo/v4"
)

func collectStatus() *models.ServiceStatusData {
	nodeName := os.Getenv("MY_NODE_NAME")
	podIP := os.Getenv("MY_POD_IP")
	hostIP := os.Getenv("MY_HOST_IP")
	podName := os.Getenv("MY_POD_NAME")
	podNamespace := os.Getenv("MY_POD_NAMESPACE")
	appGroup := os.Getenv("MY_APP_GROUP")
	appVersion := os.Getenv("MY_APP_VERSION")
	appName := os.Getenv("MY_APP_NAME")

	ssData := &models.ServiceStatusData{}
	ssData.Name = appName
	ssData.Version = appVersion
	ssData.Group = appGroup
	ssData.NodeName = nodeName
	ssData.HostIP = hostIP
	ssData.PodIP = podIP
	ssData.PodName = podName
	ssData.PodNamespace = podNamespace
	return ssData
}

func sendStatus() {
	url := "http://192.168.1.106:1323/registry/statuses"
	fmt.Println("URL:>", url)
	status := collectStatus()
	jsonData, err := json.Marshal(status)

	fmt.Println("JSON:>", jsonData, err)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
	defer resp.Body.Close()
}

func main() {
	sendStatus()
	e := echo.New()
	helloWorld := service.NewHelloWorldService()
	greeting := service.NewGreetingService()
	fibonacci := service.NewFibonacciService()
	handler := handler.NewHandler(helloWorld, greeting, fibonacci)
	handler.Register(e)
	e.Logger.Fatal(e.Start(":8080"))
}
