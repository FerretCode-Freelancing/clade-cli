package connect

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

var CurrentConnection Connection

type Connection struct {
	AuthSecret    string
	ServerAddress string
}

func Preload() error {
	connection, err := ReadConnectFile()
	if err != nil {
		return err
	}

	CurrentConnection = connection

	return nil
}

func Connect(args []string, insecure bool) error {
	if len(args) != 2 {
		connectHelp()
	}

	serverAddress := args[0]
	secret := args[1]

	_, err := url.ParseRequestURI(serverAddress)
	if err != nil {
		return err
	}

	if !insecure && !strings.HasPrefix(serverAddress, "https://") {
		return errors.New("the insecure flag must be set to use http")
	}

	connection := Connection{
		ServerAddress: serverAddress,
		AuthSecret:    secret,
	}

	err = writeConnectFile(connection)
	if err != nil {
		return err
	}

	fmt.Println("the connection has been set up correctly")

	return nil
}

func getConnectFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	appDir := filepath.FromSlash(homeDir + "/.clade")

	if _, err := os.Stat(appDir); err != nil {
		return "", err
	}

	return filepath.FromSlash(appDir + "/connect.json"), nil
}

func ReadConnectFile() (Connection, error) {
	connectFilePath, err := getConnectFilePath()
	if err != nil {
		return Connection{}, err
	}

	bytes, err := os.ReadFile(connectFilePath)
	if err != nil {
		return Connection{}, err
	}

	connection := Connection{}

	if err := json.Unmarshal(bytes, &connection); err != nil {
		return Connection{}, err
	}

	return connection, nil
}

func writeConnectFile(connection Connection) error {
	connectFilePath, err := getConnectFilePath()
	if err != nil {
		return err
	}

	bytes, err := json.Marshal(connection)
	if err != nil {
		return err
	}

	err = os.WriteFile(connectFilePath, bytes, os.ModeAppend)
	if err != nil {
		return err
	}

	return nil
}

func connectHelp() {
	log.Fatalf("\nconnect\n\t1. server address: the hostname or IP address of the clade server\n\t2. auth secret: the secret used to authenticate requests made to the clade server\n")
}
