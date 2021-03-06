# Terraform provider for Confluent Kafka Schema Registry

This terraform provider allows you to interact with Confluent Kafka Schema Registry.

There are a few requirements about how the API must work for this provider to be able to do its thing:
* The API is expected to support the following HTTP methods:
    * POST: create or update a schema
    * DELETE: remove a schema

Have a look at the [examples directory](examples) for some use cases

&nbsp;

## Provider configuration
- `uri` (string, required): URI of the Schema Registry REST API endpoint. This serves as the base of all requests. Example: `http://localhost:8001`.
- `api_key` (string, optional): API key to access Confluent Cloud Schema Registry
- `api_secret` (string, optional): API secret to access Confluent Cloud Schema Registry

&nbsp;

## `schemaregistry_subject` resource configuration
- `subject` (string, required): The name of the subject to be created.
- `schema` (string, required): The schema of the subject that has to be created.
- `schema_type` (string, optional): The schema type of the subject that has to be created. Default: "AVRO"

## `schemaregistry_remote_schema` data source configuration
- `url`: (string, required): URL to a remote zip file
- `path`: (string, required): Path to schema file inside zip file
- `value`: (string, computed): Contents of the schema file
- `auth`: (block, optional): Authentication to remote schema source. Contents:
    - `aws`: (block, optional): AWS-based authentication. Uses default authentication provider chain. Contents:
        - `region`: (string, required): Region of AWS bucket

Each distinct zip file will be downloaded only once. 

## Installation

There are two standard methods of installing this provider detailed [in Terraform's documentation](https://www.terraform.io/docs/configuration/providers.html#third-party-plugins). You can place the file in the directory of your .tf file in `terraform.d/plugins/{OS}_{ARCH}/` or place it in your home directory at `~/.terraform.d/plugins/{OS}_{ARCH}/`

Once downloaded, be sure to make the plugin executable by running `chmod +x terraform-provider-schemaregistry`.

## Extras

In the [examples directory](examples) there is a `docker-compose.yml` file.
This is useful for development or for testing purposes.

## Todo

- [x] write docker compose file
- [x] write samples
- [ ] write tests
