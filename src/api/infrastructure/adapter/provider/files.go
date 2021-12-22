package provider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"
)

type FilesProvider interface {
	GetFileAsString(string) (string, error)
	GetFileAsJSON(string) (*map[string]interface{}, error)
}

type filesProviderImplementation struct {
	resourcePath string
}

func NewFilesProvider() *filesProviderImplementation {
	return &filesProviderImplementation{
		resourcePath: getLocalResourcesPath(),
	}
}

func (impl *filesProviderImplementation) GetFileAsJSON(fileName string) (*map[string]interface{}, error) {
	jsonMap := make(map[string]interface{})
	jsonString, err := impl.GetFileAsString(fileName)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(jsonString), &jsonMap); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &jsonMap, nil
}

func (impl *filesProviderImplementation) GetFileAsString(fileName string) (string, error) {
	content, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", impl.resourcePath, fileName))

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return string(content), nil
}

func getLocalResourcesPath() string {
	_, filename, _, ok := runtime.Caller(1)

	if !ok {
		log.Fatal("error loading configs", nil)
		return ""
	}

	return filepath.Join(filepath.Dir(filename), "../../../resources")
}
