package layer0

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
)

type HttpClient struct {
	targetUrl string
}

func NewHttpClient(targetUrl string) *HttpClient {
	return &HttpClient{
		targetUrl: targetUrl,
	}
}

func (client *HttpClient) SendMsg(outgoingMsg []byte) ([]byte, error) {
	encodedOutgoingMsg := base64.StdEncoding.EncodeToString(outgoingMsg)
	resp, err := http.PostForm(client.targetUrl, url.Values{"m": {string(encodedOutgoingMsg)}})

	fmt.Printf("Message: %s\n", encodedOutgoingMsg)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("got HTTP status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	msg, err := getMsgFromBody(body)
	if err != nil {
		return nil, err
	}

	return msg, nil
}

func getMsgFromBody(body []byte) ([]byte, error) {
	fmt.Printf("%s\n", body)
	re := regexp.MustCompile(`(<!--)([A-Za-z0-9/+=]*|=[^=]|={3,})(-->)`)
	match := re.Find(body)
	encoded := string(match[4 : len(match)-3])
	data, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}

	return data, nil
}
