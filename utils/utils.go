// Package utils contains methods for reading tests from files.
package utils

import "strings"
import "io/ioutil"
import "path/filepath"
import "encoding/json"
import "gopkg.in/yaml.v1"
import "github.com/golang/glog"
import "github.com/crackcomm/action-test/testing"

func ReadTests(dirname string) (tests testing.Tests, err error) {
	ext := filepath.Ext(dirname)
	if ext == "" {
		glog.V(3).Infof("Tests flag is a directory: %v", dirname)
		// If `-tests` flag contains `*` it's already a pattern
		if !strings.Contains(dirname, "*") {
			dirname = filepath.Join(dirname, "*")
		}

		// Look for tests
		glog.V(3).Infof("Looking for tests in %s", dirname)
		var files []string
		files, err = filepath.Glob(dirname)
		if err != nil {
			return
		}

		// Read tests from files in this directory
		glog.V(3).Infof("Reading files %v", files)
		tests, err = ReadFiles(files)
		return
	}

	tests, err = ReadFiles([]string{dirname})
	return
}

func ReadFiles(files []string) (tests testing.Tests, err error) {
	tests = testing.Tests{}
	for _, fname := range files {
		more := testing.Tests{}
		ext := filepath.Ext(fname)
		switch ext {
		case ".json":
			// Read json file
			var body []byte
			glog.V(3).Infof("Reading json test %s", fname)
			body, err = ioutil.ReadFile(fname)
			if err != nil {
				return
			}

			// Unmarshal json file
			err = json.Unmarshal(body, &more)
			if err != nil {
				return
			}
		case ".yaml":
			// Read yaml file
			var body []byte
			glog.V(3).Infof("Reading yaml test %s", fname)
			body, err = ioutil.ReadFile(fname)
			if err != nil {
				return
			}

			// Unmarshal yaml file
			err = yaml.Unmarshal(body, &more)
			if err != nil {
				return
			}
		default:
			glog.Warningf("Ignoring file %s (ext=%#v)", fname, ext)
		}
		tests = append(tests, more...)
	}
	return
}
