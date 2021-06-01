// Package remote_schema_provider abstracts away access to schema files in remote locations
package remote_schema_provider

import (
  "fmt"
  "net/http"
  "sync"
)

// RemoteSchemaProvider provides access to remote schema files
type RemoteSchemaProvider interface {
  // GetZippedSchema returns raw contents of a schema file located under `path` inside a zip file from `url` remote
  GetZippedSchema(url string, path string, auth *Auth) (string, error)
}

// New creates a new instance of RemoteSchemaProvider
func New() RemoteSchemaProvider {
  return &remoteSchemaSources{
    make(map[string]*remoteSchemaSource),
    sync.RWMutex{},
  }
}

type remoteSchemaSources struct {
  sources map[string]*remoteSchemaSource
  lock    sync.RWMutex
}

func (r *remoteSchemaSources) readSource(url string) (*remoteSchemaSource, bool) {
  r.lock.RLock()
  defer r.lock.RUnlock()
  s, ok := r.sources[url]
  return s, ok
}

func (r *remoteSchemaSources) writeSource(url string, source *remoteSchemaSource) {
  r.lock.Lock()
  defer r.lock.Unlock()
  r.sources[url] = source
}

func (r *remoteSchemaSources) GetZippedSchema(url string, path string, auth *Auth) (string, error) {
  _, ok := r.readSource(url)
  if !ok {
    r.writeSource(url, &remoteSchemaSource{})
  }
  s, _ := r.readSource(url)
  return s.getSchema(url, path, auth)
}

type remoteSchemaSource struct {
  paths map[string]string
  once  sync.Once
}

func (r *remoteSchemaSource) getSchema(url string, path string, auth *Auth) (schema string, err error) {
  zipName, err := getZipName(url)
  if err != nil {
    return "", err
  }
  dirName := getDirName(zipName)

  r.once.Do(func() {
    var client *http.Client
    client, err = auth.getHttpClientWithAuthentication()
    if err == nil {
      err = download(url, zipName, client)
      if err == nil {
        r.paths, err = unzip(zipName, dirName)
      }
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
