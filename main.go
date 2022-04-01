package main

import (
	"errors"
	"log"

	sidra "github.com/ZoeySimone/sidra/internal"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

type ErrorLog struct {
	plugin *protogen.Plugin
}

func NewErrorLog(p *protogen.Plugin) *ErrorLog {
	elog := &ErrorLog{}
	elog.plugin = p

	return elog
}

func (e *ErrorLog) Write(p []byte) (n int, err error) {
	e.plugin.Error(errors.New(string(p)))
	return len(p), nil
}

func main() {
	var provider sidra.Provider

	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		log.SetOutput(NewErrorLog(gen))
		var filename string
		var importPath protogen.GoImportPath
		for _, f := range gen.Files {
			if f.Generate {
				filename = f.GeneratedFilenamePrefix
				importPath = f.GoImportPath
				sidra.GenerateFile(gen, f, &provider)
			}
		}

		// generate the terraform provider file
		providerFile := gen.NewGeneratedFile(filename+"_provider.go", importPath)
		sidra.WriteTemplate("templates/provider.tmpl", providerFile, provider)

		// parse opt flag if provided
		if gen.Request.Parameter != nil {
			opt := *gen.Request.Parameter
			switch opt {
			case "no-grpc":
				// generate service interfaces if not using gRPC
				serviceFile := gen.NewGeneratedFile(filename+"_services.go", importPath)
				sidra.WriteTemplate("templates/service.tmpl", serviceFile, provider)
			}
		}

		// add support for optional fields
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

		return nil
	})
}
