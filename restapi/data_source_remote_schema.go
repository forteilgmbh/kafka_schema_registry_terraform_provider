package restapi

import (
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
  return client.remoteSchemaProvider.GetZippedSchema(url, path)
}
