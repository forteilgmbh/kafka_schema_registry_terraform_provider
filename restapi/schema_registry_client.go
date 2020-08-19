package restapi

import (
  "encoding/json"
  "bytes"
  "io/ioutil"
  "net/http"
  "fmt"
)

type credentials struct {
	apiKey    string
	apiSecret string
}

type schemaRegistryClient struct {
	httpClient  *http.Client
	createUri   func(string) string
	deleteUri   func(string) string
	credentials *credentials
}

func NewSchemaRegistryClient(uri string, credentials *credentials) (*schemaRegistryClient, error) {
	client := schemaRegistryClient{
		createUri:   func(subject string) string { return uri + "/subjects/" + subject + "/versions" },
		deleteUri:   func(subject string) string { return uri + "/subjects/" + subject },
		httpClient:  &http.Client{},
		credentials: credentials,
  }

  return &client, nil
}

func (client *schemaRegistryClient) createSubject(subject string, schema string) error {
	jsonData := map[string]string{"schema": schema}
  jsonValue, err := json.Marshal(jsonData)

  if err != nil {
    return err
  }

	req, err := http.NewRequest("POST", client.createUri(subject), bytes.NewBuffer(jsonValue))
	if err != nil {
		return err
	}
	if client.credentials != nil {
		req.SetBasicAuth(client.credentials.apiKey, client.credentials.apiSecret)
	}
	req.Header.Set("Content-Type", "application/vnd.schemaregistry.v1+json")

	response, err := client.httpClient.Do(req)

  if err != nil {
    return err
  }

  data, err := ioutil.ReadAll(response.Body)

  if err != nil {
    return err
  }

  if response.StatusCode != http.StatusOK {
    err = fmt.Errorf("response code is %d: %s", response.StatusCode, data)
    return err
  }

  return nil
}

func (client *schemaRegistryClient) deleteSubject(subject string) error {
	request, err := http.NewRequest("DELETE", client.deleteUri(subject), nil)

  if err != nil {
    return err
  }

	if client.credentials != nil {
		request.SetBasicAuth(client.credentials.apiKey, client.credentials.apiSecret)
  }

	response, err := client.httpClient.Do(request)

  if err != nil {
    return err
  }

  data, err := ioutil.ReadAll(response.Body)

  if err != nil {
    return err
  }

  if response.StatusCode != http.StatusOK {
    err = fmt.Errorf("response code is %d: %s", response.StatusCode, data)
    return err
  }

  defer response.Body.Close()

  return err
}
