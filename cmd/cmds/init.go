package cmds

import (
	"flag"
	"fmt"

	"github.com/kenriortega/gotool/pkg/standard"
)

var (
	dstProjectPath = flag.String("out", ".", "Project Location")
	nameProject    = flag.String("name", "basics", "Project Name")
)

func Start() {
	flag.Parse()

	fmt.Println("Creating project structure...")
	standard.FoldersStructure(*dstProjectPath, *nameProject)
	standard.FilesInit(*dstProjectPath, *nameProject, "go.mod", ".gitignore")
	fmt.Printf("Go to path %s/%s\n", *dstProjectPath, *nameProject)

}
