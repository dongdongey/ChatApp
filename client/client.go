package client

import (
	"encoding/json"
	"net"
)

type Client struct {
	Name       string
	connection net.Conn
}
type Message struct {
	Name    string
	Content string
}

func NewClient(name string) *Client {
	return &Client{Name: name, connection: nil}
}

func (c *Client) Connect(targetIP string) error {
	conn, err := net.Dial("tcp", targetIP)
	if err != nil {
		return err
	}
	c.connection = conn
	return nil
}

func (c *Client) Send(message string) error {
	data := Data{}
	data.SetName([]byte(c.Name))
	data.SetPayload([]byte(message))
	dataBytes := data.ToBytes()

	_, err := c.connection.Write(dataBytes)
	return err
}

func (c *Client) Destroy() {
	err := c.connection.Close()
	if err != nil {
		c.connection = nil
		return
	}
	c.connection = nil
}

func (c *Client) Receive() (string, error) {
	buf := make([]byte, 1024)
	_, err := c.connection.Read(buf)
	if err != nil {
		return "", err
	}
	data := FromBytes(buf)
	message := Message{Name: string(data.name), Content: string(data.payload)}
	jsonMes, err := json.Marshal(message)
	if err != nil {
		return "", err
	}
	return string(jsonMes), nil
}
