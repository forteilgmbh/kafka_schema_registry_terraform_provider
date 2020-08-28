// Package remote_schema_provider abstracts away access to schema files in remote locations
package remote_schema_provider

import (
  "fmt"
  "sync"
)

// RemoteSchemaProvider provides access to remote schema files
type RemoteSchemaProvider interface {
  // GetZippedSchema returns raw contents of a schema file located under `path` inside a zip file from `url` remote
  GetZippedSchema(url string, path string) (string, error)
}

// New creates a new instance of RemoteSchemaProvider
func New() RemoteSchemaProvider {
  var r remoteSchemaSources
  r = make(map[string]*remoteSchemaSource)
  return r
}

type remoteSchemaSources map[string]*remoteSchemaSource

func (r remoteSchemaSources) GetZippedSchema(url string, path string) (string, error) {
  _, ok := r[url]
  if !ok {
    r[url] = &remoteSchemaSource{}
  }
 return r[url].getSchema(url, path)
}

type remoteSchemaSource struct {
  paths map[string]string
  once  sync.Once
}

func (r *remoteSchemaSource) getSchema(url string, path string) (schema string, err error) {
  zipName, err := getZipName(url)
  if err != nil {
    return "", err
  }
  dirName := getDirName(zipName)

  r.once.Do(func() {
    err = download(url, zipName)
    if err == nil {
      r.paths, err = unzip(zipName, dirName)
    }
  })
  if err != nil {
    return "", err
  }
  file, ok := r.paths[path]
  if !ok {
    return "", fmt.Errorf("%s not found in %s", path, url)
  }
  return readFile(file)
}