package remote_schema_provider

import (
  "archive/zip"
  "fmt"
  "io"
  "io/ioutil"
  "net/http"
  "net/url"
  "os"
  "path"
  "path/filepath"
  "strings"
)

func cleanOldDir(dir string) error {
  return os.RemoveAll(dir)
}

func getZipName(targetUrl string) (string, error) {
  parsedUrl, err := url.Parse(targetUrl)
  if err != nil {
    return "", err
  }
  return path.Base(parsedUrl.Path), nil
}

func getDirName(zipName string) string {
  return strings.TrimSuffix(zipName, filepath.Ext(zipName))
}

func download(url string, dest string) error {
  // Get the data
  resp, err := http.Get(url)
  if err != nil {
    return err
  }
  defer resp.Body.Close()

  // Create the file
  out, err := os.Create(dest)
  if err != nil {
    return err
  }
  defer out.Close()

  // Write the body to file
  _, err = io.Copy(out, resp.Body)
  return err
}

func unzip(src string, dest string) (paths map[string]string, err error) {
  r, err := zip.OpenReader(src)
  if err != nil {
    return paths, err
  }
  defer r.Close()

  paths = make(map[string]string)
  for _, f := range r.File {

    // Store filename/path for returning and using later on
    fpath := filepath.Join(dest, f.Name)

    // Check for ZipSlip vulnerability: https://snyk.io/research/zip-slip-vulnerability
    if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
      return paths, fmt.Errorf("%s: illegal file path", fpath)
    }

    if f.FileInfo().IsDir() {
      // Make Folder
      os.MkdirAll(fpath, os.ModePerm)
      continue
    }

    // Make File
    if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
      return paths, err
    }

    outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
    if err != nil {
      return paths, err
    }

    rc, err := f.Open()
    if err != nil {
      return paths, err
    }

    _, err = io.Copy(outFile, rc)

    // Close the file without defer to close before next iteration of loop
    outFile.Close()
    rc.Close()

    if err != nil {
      return paths, err
    }

    paths[f.Name] = fpath
  }
  return paths, nil
}

func readFile(file string) (string, error){
  content, err := ioutil.ReadFile(file)
  return string(content[:]), err
}
