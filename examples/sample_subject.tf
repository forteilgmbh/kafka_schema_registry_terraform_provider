provider "schemaregistry" {
  uri = "http://localhost:8081"
}

resource "schemaregistry_subject" "schema_sample_from_string" {
  subject = "com.test.myapp.test-from-string"
  schema  = "{\"type\": \"string\"}"
}

resource "schemaregistry_subject" "schema_sample_from_file" {
  subject = "com.test.myapp.test-from-file"
  schema  = "${file("schema_sample.avro.json")}"
}

data "schemaregistry_remote_schema" "test_schema" {
  url = "https://some-remote-url.com/schemas-1.0.0.zip"
  path = "test/schema.avsc"
}

resource "schemaregistry_subject" "schema_sample_from_remote" {
  subject = "com.test.myapp.test-from-remote"
  schema = data.schemaregistry_remote_schema.test_schema.value
}
