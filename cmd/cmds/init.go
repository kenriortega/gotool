package cmds

import (
	"flag"
	"fmt"

	"github.com/kenriortega/gotool/pkg/chi"
	"github.com/kenriortega/gotool/pkg/standard"
)

var (
	typeProject    = flag.String("type", "basic", "Project Type")
	dstProjectPath = flag.String("out", ".", "Project Location")
	nameProject    = flag.String("name", "basics", "Project Name")
)

func Start() {
	flag.Parse()

	fmt.Println("Creating project structure...")

	switch *typeProject {
	case "basic":
		standard.FoldersInit(*dstProjectPath, *nameProject)
		standard.FilesInit(*dstProjectPath, *nameProject, "go.mod", ".gitignore")
		fmt.Printf("Go to path cd %s/%s\n", *dstProjectPath, *nameProject)
	case "chi":
		chi.FoldersInit(*dstProjectPath, *nameProject)
		chi.FilesInit(*dstProjectPath, *nameProject, "go.mod", ".gitignore", ".env")
		fmt.Printf("Go to path cd %s/%s\n", *dstProjectPath, *nameProject)
		fmt.Println("Then run go mod tidy")
	}

}
