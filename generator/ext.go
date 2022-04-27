package generator

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/golang/protobuf/proto"
)

var toolName = filepath.Base(os.Args[0])

func init() {
	temp := strings.Split(toolName, "-")
	toolName = temp[len(temp)-1]
}

func (g *Generator) toolName() string {
	return "protoc-gen-" + toolName
}
func (g *Generator) pluginName() string {
	return plugins[0].Name()
}
func Main(ps ...Plugin) {
	for _, p := range ps {
		RegisterPlugin(p)
	}
	// Begin by allocating a generator. The request and response structures are stored there
	// so we can do error handling easily - the response structure contains the field to
	// report failure.
	var (
		data []byte
		err  error
	)
	// just for debug PB_INPUT=xx protoc ....
	if in, ok := os.LookupEnv("PB_INPUT"); ok {
		f, err := os.Open(in)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		data, err = ioutil.ReadAll(f)
	} else if out, ok := os.LookupEnv("PB_OUTPUT"); ok { // just for debug PB_OUTPUT=xx protoc ....
		data, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		err = ioutil.WriteFile(out, data, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	} else {
		data, err = ioutil.ReadAll(os.Stdin)
	}

	g := New()

	if err != nil {
		g.Error(err, "reading input")
	}

	if err := proto.Unmarshal(data, g.Request); err != nil {
		g.Error(err, "parsing input proto")
	}

	if len(g.Request.FileToGenerate) == 0 {
		g.Fail("no files to generate")
	}

	//g.CommandLineParameters(AddPluginToParams(g.Request.GetParameter()))
	g.CommandLineParameters(g.Request.GetParameter())

	// Create a wrapped version of the Descriptors and EnumDescriptors that
	// point to the file that defines them.
	g.WrapTypes()

	g.SetPackageNames()
	g.BuildTypeNameMap()

	g.GenerateAllFiles()

	// Send back the results.
	data, err = proto.Marshal(g.Response)
	if err != nil {
		g.Error(err, "failed to marshal output proto")
	}
	_, err = os.Stdout.Write(data)
	if err != nil {
		g.Error(err, "failed to write output proto")
	}
}
