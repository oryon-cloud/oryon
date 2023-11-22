// Copyright 2023 The Oryon Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pkm

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/oryon-cloud/oryon/app"
)

type remoteFetcher struct {
	*localFetcher
	downloadUrl string
}

// download and extract tarball
func (f *remoteFetcher) downloadAndExtractTarball() error {
	// check url valid
	urlFile, err := url.Parse(f.downloadUrl)
	if err != nil {
		return err
	}
	// check url scheme
	if urlFile.Scheme != "http" && urlFile.Scheme != "https" {
		return errors.New("invalid url scheme")
	}
	// check filename is tar.gz or tgz
	if !strings.HasSuffix(urlFile.Path, ".tar.gz") && !strings.HasSuffix(urlFile.Path, ".tgz") {
		return errors.New("invalid file type")
	}

	// download tarball
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	resp, err := client.Get(f.downloadUrl)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	// decompress tarball
	gr, err := gzip.NewReader(resp.Body)
	if err != nil {
		return err
	}
	defer gr.Close()

	tmpPath, err := os.MkdirTemp(os.TempDir(), "oryon")
	if err != nil {
		return err
	}

	// extract tarball
	tarReader := tar.NewReader(gr)
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("ExtractTarGz: Next() failed: %s", err.Error())
		}

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.Mkdir(filepath.Join(tmpPath, header.Name), os.FileMode(header.Mode)); err != nil {
				log.Fatalf("ExtractTarGz: Mkdir() failed: %s", err.Error())
			}
		case tar.TypeReg:
			outFilePath := filepath.Join(tmpPath, header.Name)
			// create directory if not exist
			if _, err := os.Stat(filepath.Dir(outFilePath)); os.IsNotExist(err) {
				if err := os.MkdirAll(filepath.Dir(outFilePath), os.ModePerm); err != nil {
					return err
				}
			}
			outFile, err := os.Create(outFilePath)
			if err != nil {
				log.Fatalf("ExtractTarGz: Create() failed: %s", err.Error())
			}
			if _, err := io.Copy(outFile, tarReader); err != nil {
				log.Fatalf("ExtractTarGz: Copy() failed: %s", err.Error())
			}
			outFile.Close()

		}

	}

	// get subdirectory
	// eg :/var/folders/kq/1w7hvp5n2rg1mwc_2vbngrph0000gn/T/oryon2730360437/test-0.0.0
	subdirs, err := os.ReadDir(tmpPath)
	if err != nil {
		return err
	}
	if len(subdirs) != 1 {
		return errors.New("invalid tarball")
	}

	f.localPath = filepath.Join(tmpPath, subdirs[0].Name())
	return nil
}

// implement Fetcher interface
func (f *remoteFetcher) Fetch() (*app.ModuleManifest, error) {
	if err := f.downloadAndExtractTarball(); err != nil {
		return nil, err
	}

	mm, err := f.localFetcher.Fetch()
	if err != nil {
		return nil, err
	}
	return mm, nil
}

// newRemoteFetcher creates a remote fetcher
func newRemoteFetcher(url string) (Fetcher, error) {
	return &remoteFetcher{
		downloadUrl:  url,
		localFetcher: &localFetcher{},
	}, nil
}
