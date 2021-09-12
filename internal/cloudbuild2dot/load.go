package cloudbuild2dot

import (
	"errors"
	"log"
	"strings"

	"github.com/spf13/afero"
)

// LoadCloudBuildConfig looks for the Cloud Build configuration file (currently
// only YAML is supported) and returns a CloudBuildConfig.
func LoadCloudBuildConfig(filename string) (cloudBuildConfig CloudBuildConfig, err error) {
	return loadConfig(afero.NewOsFs(), filename)
}

func loadConfig(AppFs afero.Fs, filename string) (CloudBuildConfig, error) {

	if strings.Contains(filename, "yaml") || strings.Contains(filename, "yml") {
		yamlContent, err := afero.ReadFile(AppFs, filename)
		if err != nil {
			log.Fatal(err)
		} else {
			return yamlToCloudBuild(yamlContent), nil
		}
	}

	if strings.Contains(filename, "json") {
		jsonContent, err := afero.ReadFile(AppFs, filename)
		if err != nil {
			log.Fatal(err)
		} else {
			return jsonToCloudBuild(jsonContent), nil
		}
	}

	return CloudBuildConfig{}, errors.New("could not find any Cloud Build config file in the current working directory")
}
