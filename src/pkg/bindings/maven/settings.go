package maven

import (
	"bytes"
	"html/template"
	"io"
)

type Settings struct {
	username string
	password string
	repoId   string
}

const settingsTemplate = `
<settings xmlns="http://maven.apache.org/SETTINGS/1.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
  xsi:schemaLocation="http://maven.apache.org/SETTINGS/1.0.0 https://maven.apache.org/xsd/settings-1.0.0.xsd">
  <servers>
    <server>
      <id>{{.RepoId}}</id>
      <username>{{.Username}}</username>
      <password>{{.Password}}</password>
      <configuration></configuration>
    </server>
  </servers>
</settings>
`

func NewSettings(username, password, repo string) *Settings {
	return &Settings{
		username: username,
		password: password,
		repoId:   repo,
	}
}

func (s *Settings) GenerateContent() (io.Reader, error) {
	var b *bytes.Buffer = new(bytes.Buffer)
	tmpl, err := template.New("template").Parse(settingsTemplate)
	if err != nil {
		return nil, err
	}

	c := struct {
		Username string
		Password string
		RepoId   string
	}{
		Username: s.username,
		Password: s.password,
		RepoId:   s.repoId,
	}
	if err = tmpl.Execute(b, c); err != nil {
		return nil, err
	}
	return b, nil
}

func (s *Settings) Name() string {
	return "settings.xml"
}
