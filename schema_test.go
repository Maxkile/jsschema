package schema_test

import (
	schema "github.com/Maxkile/jsschema"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadSchema(t *testing.T) {
	files := []string{"schema.json", "qiita.json"}
	for _, f := range files {
		file := filepath.Join("test", f)
		_, err := readSchema(file)
		if !assert.NoError(t, err, "readSchema(%s) should succeed", file) {
			return
		}
	}
}

func readSchema(f string) (*schema.Schema, error) {
	in, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	return schema.Read(in)
}

func TestExtras(t *testing.T) {
	const src = `{
  "extra1": "foo",
  "extra2": ["bar", "baz"]
}`
	s, err := schema.Read(strings.NewReader(src))
	if !assert.NoError(t, err, "schema.Read should succeed") {
		return
	}

	for _, ek := range []string{"extra1", "extra2"} {
		_, ok := s.Extras[ek]
		if !assert.True(t, ok, "Extra item '%s' should exist", ek) {
			return
		}
	}
}
