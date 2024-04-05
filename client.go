package autoderm_go_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
)

type Client struct {
	apiKey string
}

func NewClient(apiKey string) *Client {
	return &Client{apiKey: apiKey}
}

func (c *Client) Query(skinImageName string, skinImageReader io.Reader, saveImage bool) (*QueryResponse, error) {
	var (
		apiURL = "https://autoderm.ai/v1/query" +
			"?model=autoderm_v2_2" +
			"&language=en" +
			"&save_image=" + strconv.FormatBool(saveImage)
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

	result := &QueryResponse{}
	err = json.NewDecoder(response.Body).Decode(result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
