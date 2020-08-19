package restapi

import (
  "github.com/hashicorp/terraform/helper/schema"
  "net/url"
)

func Provider() *schema.Provider {
  return &schema.Provider{
    Schema: map[string]*schema.Schema{
      "uri": &schema.Schema{
        Type:        schema.TypeString,
        Required:    true,
        Description: "Kafka schema registry endpoint. Example: http://localhost:8000",
      },
      "api_key": &schema.Schema{
        Type:        schema.TypeString,
        Optional:    true,
        Description: "API key to access Confluent Cloud Schema Registry",
      },
      "api_secret": &schema.Schema{
        Type:        schema.TypeString,
        Optional:    true,
        Description: "API secret to access Confluent Cloud Schema Registry",
      },
    },
    ResourcesMap: map[string]*schema.Resource{
      "schemaregistry_subject": resourceSubject(),
    },
    ConfigureFunc: configureProvider,
  }
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
  endpoint := d.Get("uri").(string)
  _, err := url.ParseRequestURI(endpoint)

  if err != nil {
    return nil, err
  }

  var c *credentials
  apiKey, apiSecret := d.Get("api_key"), d.Get("api_secret")
  if apiKey != nil && apiSecret != nil {
    c = &credentials{apiKey.(string), apiSecret.(string)}
  }

  return NewSchemaRegistryClient(endpoint, c)
}
