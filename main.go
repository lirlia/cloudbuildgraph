package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/lirlia/cloudbuildgraph/internal/cloudbuild2dot"
)

func main() {
	log.SetFlags(0)

	var (
		outputTypeList = []string{"jpeg", "pdf", "png", "svg"}
		filename       = flag.String("config", "cloudbuild.yaml", "cloudbuild config name")
		outputType     = flag.String("type", "pdf", "output type ("+strings.Join(outputTypeList, "/")+")")
	)
	flag.Parse()

	if !contains(outputTypeList, *outputType) {
		log.Fatal(fmt.Errorf("invalud output type"))
	}

	cloudBuildConfig, err := cloudbuild2dot.LoadCloudBuildConfig(*filename)
	if err != nil {
		log.Fatal(err)
	}

	dotFile, err := ioutil.TempFile("", "cloudbuild.*.dot")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(dotFile.Name())

	_, err = dotFile.Write([]byte(cloudbuild2dot.BuildDotFile(cloudBuildConfig)))
	if err != nil {
		log.Fatal(err)
	}

	err = dotFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("dot", fmt.Sprintf("-T%s", *outputType), fmt.Sprintf("-o%s.%s", *filename, *outputType), dotFile.Name())
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Successfully created %s.%s\n", *filename, *outputType)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
