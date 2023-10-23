package generate

import (
	"fmt"

	"github.com/yeom-c/golang-fiber-api/tool/cmd/generate/template"
	"github.com/yeom-c/golang-fiber-api/tool/enum"

	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

type genFile struct {
	name    string
	genType enum.GenType
}

func Run(cmd *cobra.Command, argsName []string) {
	genFiles := []genFile{}
	// module (controller, service, repository, model)
	if len(argsName) > 0 {
		modules := strings.Split(argsName[0], ",")
		for _, name := range modules {
			genFiles = append(genFiles, genFile{
				name:    name,
				genType: enum.GenTypeModuleController,
			})
			genFiles = append(genFiles, genFile{
				name:    name,
				genType: enum.GenTypeModuleControllerTest,
			})
			genFiles = append(genFiles, genFile{
				name:    name,
				genType: enum.GenTypeModuleService,
			})
			genFiles = append(genFiles, genFile{
				name:    name,
				genType: enum.GenTypeModuleServiceTest,
			})
			genFiles = append(genFiles, genFile{
				name:    name,
				genType: enum.GenTypeRepository,
			})
			genFiles = append(genFiles, genFile{
				name:    name,
				genType: enum.GenTypeRepositoryTest,
			})
			genFiles = append(genFiles, genFile{
				name:    name,
				genType: enum.GenTypeModel,
			})
		}
	}

	// controller
	flagController := cmd.Flag("controller").Value.String()
	if flagController != "" {
		controllers := strings.Split(flagController, ",")
		for _, name := range controllers {
			genFiles = append(genFiles, genFile{
				name:    name,
				genType: enum.GenTypeController,
			})
			genFiles = append(genFiles, genFile{
				name:    name,
				genType: enum.GenTypeControllerTest,
			})
		}
	}

	// service
	flagService := cmd.Flag("service").Value.String()
	if flagService != "" {
		services := strings.Split(flagService, ",")
		for _, name := range services {
			genFiles = append(genFiles, genFile{
				name:    name,
				genType: enum.GenTypeService,
			})
			genFiles = append(genFiles, genFile{
				name:    name,
				genType: enum.GenTypeServiceTest,
			})
		}
	}

	// repository
	flagRepository := cmd.Flag("repository").Value.String()
	if flagRepository != "" {
		repositories := strings.Split(flagRepository, ",")
		for _, name := range repositories {
			genFiles = append(genFiles, genFile{
				name:    name,
				genType: enum.GenTypeRepository,
			})
			genFiles = append(genFiles, genFile{
				name:    name,
				genType: enum.GenTypeRepositoryTest,
			})
		}
	}

	// model
	flagModel := cmd.Flag("model").Value.String()
	if flagModel != "" {
		models := strings.Split(flagModel, ",")
		for _, name := range models {
			genFiles = append(genFiles, genFile{
				name:    name,
				genType: enum.GenTypeModel,
			})
		}
	}

	// make generate file
	for _, genFile := range genFiles {
		genFile.generate()
	}
}

func makeFile(filename, path string, fileString string) error {
	// create the directory if it does not exist
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}

	// check file exists
	fullPath := filepath.Join(path, filename)
	_, err = os.Stat(fullPath)
	if os.IsNotExist(err) {
		// file does not exist
		// Create the file
		file, err := os.Create(fullPath)
		if err != nil {
			return err
		}
		defer file.Close()

		// write the contents to the file
		_, err = file.WriteString(fileString)
		if err != nil {
			panic(err)
		}
	} else {
		return fmt.Errorf("already exists %s", fullPath)
	}

	return nil
}

func (gf *genFile) generate() {
	name := gf.name
	controllerPath := fmt.Sprintf("controller/%s", name)
	servicePath := fmt.Sprintf("service/%s", name)
	repositoryPath := fmt.Sprintf("repository/%s", name)
	modelPath := fmt.Sprintf("model/entity/%s", name)

	var path, fileName string
	switch gf.genType {
	case enum.GenTypeController,
		enum.GenTypeModuleController:
		path = controllerPath
		fileName = fmt.Sprintf("%s_controller.go", name)
	case enum.GenTypeControllerTest,
		enum.GenTypeModuleControllerTest:
		path = controllerPath
		fileName = fmt.Sprintf("%s_controller_test.go", name)
	case enum.GenTypeService,
		enum.GenTypeModuleService:
		path = servicePath
		fileName = fmt.Sprintf("%s_service.go", name)
	case enum.GenTypeServiceTest,
		enum.GenTypeModuleServiceTest:
		path = servicePath
		fileName = fmt.Sprintf("%s_service_test.go", name)
	case enum.GenTypeRepository:
		path = repositoryPath
		fileName = fmt.Sprintf("%s_repository.go", name)
	case enum.GenTypeRepositoryTest:
		path = repositoryPath
		fileName = fmt.Sprintf("%s_repository_test.go", name)
	case enum.GenTypeModel:
		path = modelPath
		fileName = fmt.Sprintf("%s_entity.go", name)
	}

	// make go file
	genTemplate := template.GetTemplate(gf.genType, name)
	err := makeFile(fileName, path, genTemplate)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("created %s/%s\n", path, fileName)
	}
}
