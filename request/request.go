package request

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/ferretcode-freelancing/clade-cli/connect"
)

func MakeRequest(path string, request interface{}, insecure bool) (string, error) {
	if connect.CurrentConnection.ServerAddress == "" || connect.CurrentConnection.AuthSecret == "" {
		return "", errors.New("the connection is not configured yet")
	}

	data, err := json.Marshal(request)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(
		"POST",
		connect.CurrentConnection.ServerAddress+path,
		bytes.NewReader(data),
	)
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", "Basic "+connect.CurrentConnection.AuthSecret)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(bytes), nil

	return "", err
}
