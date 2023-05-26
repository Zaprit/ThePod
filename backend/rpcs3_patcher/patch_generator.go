package rpcs3_patcher

import (
	"bytes"
	_ "embed"
	"text/template"
)

//go:embed rpcs3_patch_template.tmpl
var templateText string

type URLBundle struct {
	HttpsURL    string
	HttpURL     string
	PresenceURL string // Used by lbp3
	LiveURL     string // Used by lbp3
}

func CreatePatch(bundle URLBundle) (string, error) {
	tmpl := template.New("game_patch")
	tmpl, err := tmpl.Parse(templateText)
	if err != nil {
		return "", err
	}
	buf := bytes.Buffer{}
	err = tmpl.Execute(&buf, bundle)
	return buf.String(), err
}
