package main

import "google.golang.org/protobuf/compiler/protogen"

const (
	contextPackage = protogen.GoImportPath("context")
)

func main() {
	options := protogen.Options{}

	options.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}

			generateFile(gen, f)
		}

		return nil
	})
}

func generateFile(gen *protogen.Plugin, file *protogen.File) {
	if len(file.Services) == 0 {
		return
	}

	fileName := file.GeneratedFilenamePrefix + ".kit.go"
	g := gen.NewGeneratedFile(fileName, file.GoImportPath)

	g.P("// Code generated by protoc-gen-go-kit. DO NOT EDIT.")
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()
	generateFileContent(file, g)
}

// generateFileContent generates the gRPC service definitions, excluding the package statement.
func generateFileContent(file *protogen.File, g *protogen.GeneratedFile) {
	if len(file.Services) == 0 {
		return
	}

	for _, service := range file.Services {
		genService(g, service)
	}
}

func genService(g *protogen.GeneratedFile, service *protogen.Service) {
}
