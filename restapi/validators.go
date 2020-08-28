package restapi

import (
  "fmt"
  "net/url"
)

func validateRemoteSchemaUrl(v interface{}, k string) (ws []string, errors []error) {
  value := v.(string)
  _, err := url.ParseRequestURI(value)
  if err != nil {
    errors = append(errors, fmt.Errorf("%q is not a valid URI", k))
  }
  return
}
