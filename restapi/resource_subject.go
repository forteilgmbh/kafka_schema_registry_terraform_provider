package restapi

import (
  "github.com/hashicorp/terraform/helper/schema"
  "log"
)

func resourceSubject() *schema.Resource {
  return &schema.Resource{
    Create: resourceSubjectCreate,
    Read:   resourceSubjectRead,
    Update: resourceSubjectUpdate,
    Delete: resourceSubjectDelete,

    Schema: map[string]*schema.Schema{
      "subject": {
        Type:     schema.TypeString,
        Required: true,
      },
      "schema": {
        Type:     schema.TypeString,
        Required: true,
        DiffSuppressFunc: suppressEquivalentJsonDiffs,
      },
      "schema_type": {
        Type: schema.TypeString,
        Optional: true,
        Default: "AVRO",
      },
    },
  }
}

func resourceSubjectCreate(d *schema.ResourceData, m interface{}) error {
  client := m.(*schemaRegistryClient)
  subject, schema, schemaType := d.Get("subject").(string), d.Get("schema").(string), d.Get("schema_type").(string)

  log.Printf("Create subject '%s'.", client)

  err := client.createSubject(subject, schema, schemaType)

  if err != nil {
    return err
  }

  d.SetId(subject)
  return resourceSubjectRead(d, m)
}

func resourceSubjectRead(d *schema.ResourceData, m interface{}) error {
  client := m.(*schemaRegistryClient)
  subject := d.Get("subject").(string)
  s, err := client.readSubjectSchema(subject)
  if err != nil {
    return err
  }
  return d.Set("schema", s)
}

func resourceSubjectUpdate(d *schema.ResourceData, m interface{}) error {
  client := m.(*schemaRegistryClient)
  subject, schema, schemaType := d.Get("subject").(string), d.Get("schema").(string), d.Get("schema_type").(string)

  log.Printf("Update subject '%s'.", client)

  err := client.createSubject(subject, schema, schemaType)

  if err != nil {
    return err
  }

  return resourceSubjectRead(d, m)
}

func resourceSubjectDelete(d *schema.ResourceData, m interface{}) error {
  client := m.(*schemaRegistryClient)
  subject := d.Get("subject").(string)

  log.Printf("Delete subject '%s'.", client)

  err := client.deleteSubject(subject)

  if err != nil {
    return err
  }

  d.SetId("")

  return nil
}
