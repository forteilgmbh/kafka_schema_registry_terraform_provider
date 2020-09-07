package remote_schema_provider

import (
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/aws/signer/v4"
  "github.com/sha1sum/aws_signing_client"
  "net/http"
)

type Auth struct {
  Aws map[string]interface{}
}

func (a Auth) getHttpClientWithAuthentication() (*http.Client, error) {
  if len(a.Aws) > 0 {
    return a.getHttpClientWithAwsAuthentication()
  }
  return http.DefaultClient, nil
}

func (a Auth) getHttpClientWithAwsAuthentication() (*http.Client, error) {
  config := aws.Config{
    Region: aws.String(a.Aws["region"].(string)),
  }
  sess := session.Must(session.NewSession(&config))
  signer := v4.NewSigner(sess.Config.Credentials)
  return aws_signing_client.New(signer, nil, "s3", a.Aws["region"].(string))
}