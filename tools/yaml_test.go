package tools

import "testing"

func TestYamlHandler(t *testing.T) {
	var bs = []byte(`
kind: Namespace
metadata:
   name: test
---

kind: bbb
name: aaa`)
	_ = YamlHandler(bs)
}
