package tools

import (
	"bytes"
	"github.com/wonderivan/logger"
	"io"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/yaml"
)

//YamlHandler is
func YamlHandler(rawBytes []byte) []string {
	reader := bytes.NewReader(rawBytes)
	ext := runtime.RawExtension{}
	d := yaml.NewYAMLOrJSONDecoder(reader, 4096)
	var returnJsonArr []string
	for {
		if err := d.Decode(&ext); err != nil {
			if err == io.EOF {
				return returnJsonArr
			}
		}
		logger.Info(string(ext.Raw))
		returnJsonArr = append(returnJsonArr, string(ext.Raw))
	}
}
