package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"net/url"
	"os"
	"strings"

	"github.com/networld-to/go-hue/configuration"
	"github.com/networld-to/go-hue/portal"
)

func ssdpDiscover() string {
	service := "239.255.255.250:1900"
	mac_address, err := net.ResolveUDPAddr("udp4", service)
	if err != nil {
		fmt.Println("Discover Error: ", err)
	}
	send, err := net.DialUDP("udp4", nil, mac_address)
	if err != nil {
		fmt.Println("Discover Error: ", err)
	}
	defer send.Close()
	// Send SSDP Message
	ssdp_discovery_message := []byte("M-SEARCH * HTTP/1.1\r\nHOST: 239.255.255.250:1900\r\nMAN: ssdp:discover\r\nMX: 10\r\nST: \"ssdp:all\"\r\n\r\n")
	_, err = send.Write(ssdp_discovery_message)
	if err != nil {
		fmt.Println("Discover Error: ", err)
	}
	fmt.Println("Searching for Philip Hue Hub (Could take up to 30 secs)...")
	// Listen for SSDP/HTTP NOTIFY over UDP
	listen, err := net.ListenMulticastUDP("udp4", nil, mac_address)
	if err != nil {
		fmt.Println("Discover Error: ", err)
	}
	defer listen.Close()
	description_url := ""
	for {
		b := make([]byte, 256)
		_, _, err := listen.ReadFromUDP(b)
		if err != nil {
			fmt.Println("Discover Error: ", err)
		}
		payload_message := string(b)
		headers := strings.Split(payload_message, "\r\n")
		for _, header := range headers {
			datum := strings.Split(header, ": ")
			if len(datum) > 1 {
				if datum[0] == "LOCATION" {
					if strings.Contains(datum[1], "description.xml") {
						description_url = datum[1]
						break
					}
				}
			}
		}
		if strings.Contains(description_url, "description.xml") {
			break
		}
	}
	u, err := url.Parse(description_url)
	if err != nil {
		fmt.Println("Discover Error: ", err)
	}
	hostname := ""
	if strings.Contains(u.Host, ":") {
		h := strings.Split(u.Host, ":")
		hostname = h[0]
	} else {
		hostname = u.Host
	}
	fmt.Printf("Found Hub at %s\n", hostname)
	return hostname
}

type JSONConfig struct {
	Bridge string `json:"bridge"`
	Username string `json:"username"`
	Devicetype string `json:"devicetype"`
}

func main() {
	//hubHostname := ssdpDiscover()
	pp, err := portal.GetPortal()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	fmt.Println(pp)
	hubHostname := pp[0].InternalIPAddress
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please press the link button on your hub, then press [enter] to continue.")
	reader.ReadLine()
	// fmt.Println("Please enter your application name:")
	// data, _, _ := reader.ReadLine()
	// applicationName := string(data)

	applicationName := "networld-to_homecontrol"

	// fmt.Println("Please enter your device type:")
	// data1, _, _ := reader.ReadLine()
	// deviceType := string(data1)

	deviceType := "grpc-service"

	c := configuration.New(hubHostname)
	response, err := c.CreateUser(applicationName, deviceType)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	username := response[0].Success["username"].(string)
	fmt.Printf("Your username is %s\n", username)


	config := &JSONConfig{
		Bridge: pp[0].InternalIPAddress,
		Devicetype: deviceType,
		Username: username,
	}

	configStr, _ := json.Marshal(config)

	fmt.Println("Store the following JSON file to ~/.philips-hue.json")
	fmt.Println(string(configStr))
}

