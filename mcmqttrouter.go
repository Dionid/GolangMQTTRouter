package mcmqttrouter

import (
	"github.com/eclipse/paho.mqtt.golang"
	"os"
	"time"
	"log"
	"fmt"
)

func CreateConnOpts(brokerName string, clientId string, debug bool, showError bool, keepAlive time.Duration, pingTimeout time.Duration) *mqtt.ClientOptions {
	if debug {
		mqtt.DEBUG = log.New(os.Stdout, "", 0)
	}
	if showError {
		mqtt.ERROR = log.New(os.Stdout, "", 0)
	}
	opts := mqtt.NewClientOptions().AddBroker(brokerName).SetClientID(clientId)
	opts.SetKeepAlive(keepAlive * time.Second)
	opts.SetDefaultPublishHandler(standardHandler)
	opts.SetPingTimeout(pingTimeout * time.Second)
	opts.AutoReconnect = true
	opts.OnConnectionLost = func (c mqtt.Client, err error) {
		fmt.Println("!!!! MQTT CONNECTION LOST BECAUSE: " + err.Error())
	}

	return opts
}

func CreateClient(opts *mqtt.ClientOptions) *mqtt.Client {
	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return &c
}