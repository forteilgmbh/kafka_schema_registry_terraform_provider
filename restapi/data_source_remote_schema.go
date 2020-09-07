package restapi

import (
  "github.com/francescop/kafka_schema_registry_terraform_provider/remote_schema_provider"
  "github.com/hashicorp/terraform/helper/hashcode"
  "github.com/hashicorp/terraform/helper/schema"
  "strconv"
)

func dataSourceRemoteSchema() *schema.Resource {
  return &schema.Resource{
    Read: dataSourceRemoteSchemaRead,

    Schema: map[string]*schema.Schema{
      "url": {
        Type:         schema.TypeString,
        Required:     true,
        ValidateFunc: validateRemoteSchemaUrl,
      },
      "path": {
        Type:     schema.TypeString,
        Required: true,
      },
      "value": {
        Type:     schema.TypeString,
        Computed: true,
      },
      "auth" : {
        Type: schema.TypeSet,
        Optional: true,
        MaxItems: 1,
        Elem: &schema.Resource{
          Schema: map[string]*schema.Schema{
            "aws": {
              Type: schema.TypeSet,
              Optional: true,
              MaxItems: 1,
              Elem: &schema.Resource{
                Schema: map[string]*schema.Schema{
                  "region": {
                    Type:      schema.TypeString,
                    Required:  true,
                  },
                },
              },
            },
          },
        },
      },
    },
  }
}

func dataSourceRemoteSchemaRead(d *schema.ResourceData, m interface{}) error {
  s, err := getRemoteSchemaContent(d, m)
  if err != nil {
    return err
  }
  d.SetId(strconv.Itoa(hashcode.String(s)))
  return d.Set("value", s)
}

func getRemoteSchemaContent(d *schema.ResourceData, m interface{}) (string, error) {
  client := m.(*schemaRegistryClient)
  url, path := d.Get("url").(string), d.Get("path").(string)
  auth := parseAuth(d)
  return client.remoteSchemaProvider.GetZippedSchema(url, path, auth)
}

func parseAuth(d *schema.ResourceData) *remote_schema_provider.Auth {
  auth := &remote_schema_provider.Auth{}

  authSet, ok := d.GetOk("auth")
  if !ok {
    return auth
  }
  authItems := authSet.(*schema.Set).List()[0].(map[string]interface{}) // MaxItems: 1 -> should be safe

  awsSet, ok := authItems["aws"]
  if ok {
    aws := awsSet.(*schema.Set).List()[0].(map[string]interface{}) // MaxItems: 1 -> should be safe
    auth.Aws = aws
  }

  return auth
}
