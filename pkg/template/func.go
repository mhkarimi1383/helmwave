package template

import (
	"bytes"
	"github.com/Masterminds/sprig/v3"
	"github.com/helmwave/helmwave/pkg/helper"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"text/template"
)

func Tpl2yml(tpl string, yml string, data interface{}) error {
	log.WithFields(log.Fields{
		"from": tpl,
		"to":   yml,
	}).Info("📄 Render file")

	if data == nil {
		data = map[string]interface{}{}
	}

	src, err := ioutil.ReadFile(tpl)
	if err != nil {
		return err
	}

	// Template
	t := template.Must(template.New("tpl").Funcs(FuncMap()).Parse(string(src)))
	var buf bytes.Buffer
	err = t.Execute(&buf, data)
	if err != nil {
		return err
	}

	log.Debug(yml, " contents\n", buf.String())

	f, err := helper.CreateFile(yml)
	if err != nil {
		return err
	}

	_, err = f.WriteString(buf.String())
	if err != nil {
		return err
	}

	return f.Close()
}

func FuncMap() template.FuncMap {
	aliased := template.FuncMap{}

	aliases := map[string]string{
		"get":    "sprigGet",
		"hasKey": "sprigHasKey",
	}

	funcMap := sprig.TxtFuncMap()

	for orig, alias := range aliases {
		aliased[alias] = funcMap[orig]
	}

	funcMap["toYaml"] = ToYaml
	funcMap["fromYaml"] = FromYaml
	funcMap["exec"] = Exec
	funcMap["setValueAtPath"] = SetValueAtPath
	funcMap["requiredEnv"] = RequiredEnv
	funcMap["required"] = Required
	funcMap["readFile"] = ReadFile
	funcMap["get"] = Get
	funcMap["hasKey"] = HasKey

	for name, f := range aliased {
		funcMap[name] = f
	}

	return funcMap
}
