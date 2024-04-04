package autoderm_go_client

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

type Client struct {
	apiKey string
}

func NewClient(apiKey string) *Client {
	return &Client{apiKey: apiKey}
}

func (c *Client) Query(skinImageName string, skinImageReader io.Reader) ([]byte, error) {
	const (
		apiURL = "https://autoderm.firstderm.com/v1/query?model=autoderm_v2_0&language=en"
	)

	var (
		requestBody = &bytes.Buffer{}
		writer      = multipart.NewWriter(requestBody)
	)

	// Add image to request
	{
		imageWriter, err := writer.CreateFormFile("file", skinImageName)
		if err != nil {
			return nil, err
		}

		_, err = io.Copy(imageWriter, skinImageReader)
		if err != nil {
			return nil, err
		}

		err = writer.Close()
		if err != nil {
			return nil, err
		}
	}

	request, err := http.NewRequest("POST", apiURL, requestBody)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	request.Header.Set("Api-Key", c.apiKey)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d", response.StatusCode)
	}

	return io.ReadAll(response.Body)
}
