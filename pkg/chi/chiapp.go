package chi

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"
	"time"
)

func FoldersInit(out, name string) {
	folders := []string{
		fmt.Sprintf("%s/%s/cmd", out, name),
		fmt.Sprintf("%s/%s/cmd/api", out, name),
		fmt.Sprintf("%s/%s/pkg", out, name),
		fmt.Sprintf("%s/%s/pkg/storage", out, name),
		fmt.Sprintf("%s/%s/pkg/broker", out, name),
		fmt.Sprintf("%s/%s/internal/domain", out, name),
		fmt.Sprintf("%s/%s/internal/ports", out, name),
		fmt.Sprintf("%s/%s/internal/handlers", out, name),
		fmt.Sprintf("%s/%s/internal/services", out, name),
	}

	for _, folder := range folders {
		if err := os.MkdirAll(folder, os.ModePerm); err != nil {
			log.Fatal(err)
		}

		// Boilerplate files
		archives := strings.Split(folder, "/")
		archive := archives[len(archives)-1:][0]
		switch archive {
		case "cmd":
			file, err := os.Create(fmt.Sprintf("%s/main.go", folder))
			if err != nil {
				fmt.Errorf("%v", err)
				os.Exit(1)
			}
			defer file.Close()

			template.Must(
				template.New("").Parse(mainTemplate)).
				Execute(file, struct {
					Timestamp  time.Time
					ModuleName string
				}{
					time.Now(),
					name,
				})
		case "api":
			file, err := os.Create(fmt.Sprintf("%s/routes.go", folder))
			if err != nil {
				fmt.Errorf("%v", err)
				os.Exit(1)
			}
			defer file.Close()

			template.Must(
				template.New("routes").Parse(routesTemplate)).
				Execute(file, struct {
					Timestamp  time.Time
					ModuleName string
				}{
					time.Now(),
					name,
				})
			// start
			fileStart, err := os.Create(fmt.Sprintf("%s/start.go", folder))
			if err != nil {
				fmt.Errorf("%v", err)
				os.Exit(1)
			}
			defer fileStart.Close()

			template.Must(
				template.New("start").Parse(startTemplate)).
				Execute(fileStart, struct {
					Timestamp  time.Time
					ModuleName string
				}{
					time.Now(),
					name,
				})
		case "domain":
			file, err := os.Create(fmt.Sprintf("%s/basic.go", folder))
			if err != nil {
				fmt.Errorf("%v", err)
				os.Exit(1)
			}
			defer file.Close()

			template.Must(
				template.New("").Parse(domainTemplate)).
				Execute(file, struct {
					Timestamp time.Time
				}{
					time.Now(),
				})
		case "ports":
			file, err := os.Create(fmt.Sprintf("%s/basic.go", folder))
			if err != nil {
				fmt.Errorf("%v", err)
				os.Exit(1)
			}
			defer file.Close()

			template.Must(
				template.New("").Parse(portsTemplate)).
				Execute(file, struct {
					Timestamp time.Time
				}{
					time.Now(),
				})
		case "handlers":
			file, err := os.Create(fmt.Sprintf("%s/basic.go", folder))
			if err != nil {
				fmt.Errorf("%v", err)
				os.Exit(1)
			}
			defer file.Close()

			template.Must(
				template.New("").Parse(handlerTemplate)).
				Execute(file, struct {
					Timestamp  time.Time
					ModuleName string
				}{
					time.Now(),
					name,
				})
		case "services":
			file, err := os.Create(fmt.Sprintf("%s/basic.go", folder))
			if err != nil {
				fmt.Errorf("%v", err)
				os.Exit(1)
			}
			defer file.Close()

			template.Must(
				template.New("").Parse(serviceTemplate)).
				Execute(file, struct {
					Timestamp  time.Time
					ModuleName string
				}{
					time.Now(),
					name,
				})
		}
	}
}
func FilesInit(path, name string, files ...string) {
	for _, fileInit := range files {

		switch fileInit {
		case "go.mod":
			file, err := os.Create(fmt.Sprintf("%s/%s/%s", path, name, fileInit))
			if err != nil {
				fmt.Errorf("%v", err)
				os.Exit(1)
			}
			defer file.Close()

			template.Must(
				template.New("").Parse(gomodTemplate)).
				Execute(file, struct {
					Timestamp  time.Time
					ModuleName string
				}{
					time.Now(),
					name,
				})
		case ".gitignore":
			file, err := os.Create(fmt.Sprintf("%s/%s/%s", path, name, fileInit))
			if err != nil {
				fmt.Errorf("%v", err)
				os.Exit(1)
			}
			defer file.Close()

			template.Must(
				template.New("").Parse(gitignoreTemplate)).
				Execute(file, struct {
					Timestamp time.Time
				}{
					time.Now(),
				},
				)
		case ".env":
			file, err := os.Create(fmt.Sprintf("%s/%s/%s", path, name, fileInit))
			if err != nil {
				fmt.Errorf("%v", err)
				os.Exit(1)
			}
			defer file.Close()

		}
	}
}
