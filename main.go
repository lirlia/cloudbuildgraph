package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/patrickhoefler/cloudbuildgraph/internal/cloudbuild2dot"
)

func main() {
	log.SetFlags(0)

	var (
		filename = flag.String("config", "cloudbuild.yaml", "cloudbuild config name")
	)
	flag.Parse()

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

	cmd := exec.Command("dot", "-Tpdf", fmt.Sprintf("-o%s.pdf", *filename), dotFile.Name())
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Sprintln("Successfully created %t.pdf", *filename)
}
