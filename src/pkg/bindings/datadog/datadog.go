package datadog

import (
	"bytes"
	"html/template"
	"io"
)

type DataDog struct {
	apiKey string
}

const datadogTemplate = `
api_key: {{.ApiKey}}
metadata_collectors:
  - name: 'resources'
    interval: 60
log_file: AGENT_LOG_FILE
apm_config:
  enabled: true
  log_file: TRACE_LOG_FILE
`

func NewDataDog(apiKey string) *DataDog {
	return &DataDog{
		apiKey: apiKey,
	}
}

func (s *DataDog) GenerateContent() (io.Reader, error) {
	var b *bytes.Buffer = new(bytes.Buffer)
	tmpl, err := template.New("template").Parse(datadogTemplate)
	if err != nil {
		return nil, err
	}

	c := struct {
		ApiKey string
	}{
		ApiKey: s.apiKey,
	}
	if err = tmpl.Execute(b, c); err != nil {
		return nil, err
	}
	return b, nil
}

func (s *DataDog) Name() string {
	return "datadog.yaml"
}
