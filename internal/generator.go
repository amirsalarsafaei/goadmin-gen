package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"goadmin-gen/internal/config"
	"goadmin-gen/internal/utils"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

//go:embed sqlc.yaml.tmpl
var sqlcTemplate string

//go:embed sqlqueries.sql.tmpl
var sqlQueriesTemplate string

//go:embed table.go.tmpl
var tableTemplate string

func main() {
	data, err := os.ReadFile("./goadmin-gen.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Unmarshal the data
	var goAdminGen config.GoAdminGen
	err = yaml.Unmarshal(data, &goAdminGen)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	sqlcTemp, err := template.New("sqlcTemplate").
		Funcs(utils.TemplateFuncMap).Parse(sqlcTemplate)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	sqlcFilePath := fmt.Sprintf("%s/sqlc.yaml", goAdminGen.Out)
	createTemplateFile(sqlcTemp, goAdminGen, sqlcFilePath)

	sqlQueriesTemp, err := template.New("sqlQueriesTemplate").
		Funcs(utils.TemplateFuncMap).Parse(sqlQueriesTemplate)

	tableTemp, err := template.New("tableTemplate").
		Funcs(utils.TemplateFuncMap).Parse(tableTemplate)

	for _, table := range goAdminGen.Tables {
		createTemplateFile(sqlQueriesTemp, table,
			fmt.Sprintf("%s/%s/queries/generated.sql", goAdminGen.Out, strings.ToLower(table.Display)))
		createTemplateFile(tableTemp, table,
			fmt.Sprintf("%s/%s/admin/table.go", goAdminGen.Out, strings.ToLower(table.Display)))
	}

	cmd := exec.Command("sqlc", "generate", "-f", sqlcFilePath)

	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
	if err != nil {
		log.Fatalf("Failed to execute command: %s", err)
	}

}

func createTemplateFile(temp *template.Template, data any, path string) {

	err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		panic(err)
	}
	file, err := os.Create(path)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	buffer := bytes.Buffer{}

	err = temp.Execute(&buffer, data)
	if err != nil {
		panic(err)
	}

	//formattedCode, err := format.Source(buffer.Bytes())
	formattedCode := buffer.Bytes()

	if err != nil {
		panic(err)
	}
	_, err = file.Write(formattedCode)
	if err != nil {
		panic(err)
	}
}
