package main

import (
	"fmt"
)

var (
	VERSION      string
	commitHash   = ""
	commitID     = ""
	commitUnix   = ""
	buildVersion = "2015.6.2-6-gfd7e2d1-dev"
	buildTime    = "2015-06-16-0431 UTC"
	buildCount   = ""
	buildUnix    = ""
	branchName   = ""
)

func main() {
	version := fmt.Sprintf("%s\nBuild time: %s\n", VERSION, buildTime)
	Execute(fmt.Sprintf(version))
}
