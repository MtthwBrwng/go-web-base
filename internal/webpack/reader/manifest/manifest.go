package manifest

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"strings"
)

type AssetList map[string][]string
type assetResponse map[string]string

// Read assets-manifest-plugin format manifest
func Read(path string) (AssetList, error) {
	//log.Println("read:", path+"/manifest.json")
	data, err := ioutil.ReadFile(path + "/manifest.json")
	if err != nil {
		return nil, errors.Wrap(err, "go-webpack: Error when loading manifest from file")
	}

	return unmarshalManifest(data)
}

func unmarshalManifest(data []byte) (AssetList, error) {
	response := make(assetResponse)
	err := json.Unmarshal(data, &response)
	if err != nil {
		return nil, errors.Wrap(err, "go-webpack: Error unmarshaling manifest file")
	}

	assets := make(AssetList, len(response))
	for key, value := range response {
		//log.Println("found asset", key, value)
		if !strings.HasSuffix(value, ".map") {
			assets[key] = []string{value}
		}
	}
	return assets, nil
}
