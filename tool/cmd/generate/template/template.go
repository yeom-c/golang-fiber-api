package template

import (
	"strings"

	"github.com/yeom-c/golang-fiber-api/tool/enum"
)

const (
	snakeNameFormat  = "[snakeName]"
	pascalNameFormat = "[pascalName]"
	camelNameFormat  = "[camelName]"

	importProjectName = "github.com/yeom-c/golang-fiber-api"
	importDatabase    = `"` + importProjectName + `/database"`
	importRepository  = `repository_` + snakeNameFormat + ` "` + importProjectName + "/repository/" + snakeNameFormat + `"`
	importService     = `service_` + snakeNameFormat + ` "` + importProjectName + "/service/" + snakeNameFormat + `"`

	// module controller
	ModuleControllerTemplate = `package ` + snakeNameFormat + `

import (
	"github.com/gofiber/fiber/v2"
	` + importDatabase + `
	` + importRepository + `
	` + importService + `
)

type ` + pascalNameFormat + `Controller struct {
	` + pascalNameFormat + `Service *service_` + snakeNameFormat + `.` + pascalNameFormat + `Service
}

func Init` + pascalNameFormat + `Controller(db *database.DB) *` + pascalNameFormat + `Controller {
	` + camelNameFormat + `Repository := repository_` + snakeNameFormat + `.New` + pascalNameFormat + `Repository(db)
	` + camelNameFormat + `Service := service_` + snakeNameFormat + `.New` + pascalNameFormat + `Service(` + camelNameFormat + `Repository)
	` + camelNameFormat + `Controller := New` + pascalNameFormat + `Controller(` + camelNameFormat + `Service)
	
	return ` + camelNameFormat + `Controller
}

func New` + pascalNameFormat + `Controller(` + camelNameFormat + `Service *service_` + snakeNameFormat + `.` + pascalNameFormat + `Service) *` + pascalNameFormat + `Controller {
	return &` + pascalNameFormat + `Controller{
		` + camelNameFormat + `Service,
	}
}

func (c *` + pascalNameFormat + `Controller) Get(ctx *fiber.Ctx) {
	c.` + pascalNameFormat + `Service.Get()
}

func (c *` + pascalNameFormat + `Controller) Post(ctx *fiber.Ctx) {
	c.` + pascalNameFormat + `Service.Post()
}

func (c *` + pascalNameFormat + `Controller) Put(ctx *fiber.Ctx) {
	c.` + pascalNameFormat + `Service.Put()
}

func (c *` + pascalNameFormat + `Controller) Delete(ctx *fiber.Ctx) {
	c.` + pascalNameFormat + `Service.Delete()
}

`

	// module controller test
	ModuleControllerTestTemplate = `package ` + snakeNameFormat + `

import (
	"testing"

	` + importDatabase + `
)

var ` + camelNameFormat + `Controller *` + pascalNameFormat + `Controller

func init() {
	db := database.GetDB()
	` + camelNameFormat + `Controller = Init` + pascalNameFormat + `Controller(db)
}

func TestGet(t *testing.T) {}

func TestPost(t *testing.T) {}

func TestPut(t *testing.T) {}

func TestDelete(t *testing.T) {}
`

	// controller
	ControllerTemplate = `package ` + snakeNameFormat + `

import (
	"github.com/gofiber/fiber/v2"
)

type ` + pascalNameFormat + `Controller struct{}
	
func Init` + pascalNameFormat + `Controller() *` + pascalNameFormat + `Controller {
	` + camelNameFormat + `Controller := New` + pascalNameFormat + `Controller()
	
	return ` + camelNameFormat + `Controller
}
	
func New` + pascalNameFormat + `Controller() *` + pascalNameFormat + `Controller {
	return &` + pascalNameFormat + `Controller{}
}

func (c *` + pascalNameFormat + `Controller) Get(ctx *fiber.Ctx) {}

func (c *` + pascalNameFormat + `Controller) Post(ctx *fiber.Ctx) {}

func (c *` + pascalNameFormat + `Controller) Put(ctx *fiber.Ctx) {}

func (c *` + pascalNameFormat + `Controller) Delete(ctx *fiber.Ctx) {}
`

	// controller test
	ControllerTestTemplate = `package ` + snakeNameFormat + `

import (
	"testing"
)

var ` + camelNameFormat + `Controller *` + pascalNameFormat + `Controller

func init() {
	` + camelNameFormat + `Controller = Init` + pascalNameFormat + `Controller()
}

func TestGet(t *testing.T) {}

func TestPost(t *testing.T) {}

func TestPut(t *testing.T) {}

func TestDelete(t *testing.T) {}
`

	// module service
	ModuleServiceTemplate = `package ` + snakeNameFormat + `

import (
	` + importRepository + `
)

type ` + pascalNameFormat + `Service struct {
	` + pascalNameFormat + `Repository *repository_` + snakeNameFormat + `.` + pascalNameFormat + `Repository
}

func New` + pascalNameFormat + `Service(` + camelNameFormat + `Repository *repository_` + snakeNameFormat + `.` + pascalNameFormat + `Repository) *` + pascalNameFormat + `Service {
	return &` + pascalNameFormat + `Service{
		` + camelNameFormat + `Repository,
	}
}

func (s *` + pascalNameFormat + `Service) Get() {
	s.` + pascalNameFormat + `Repository.Find()
}

func (s *` + pascalNameFormat + `Service) Post() {
	s.` + pascalNameFormat + `Repository.Store()
}

func (s *` + pascalNameFormat + `Service) Put() {
	s.` + pascalNameFormat + `Repository.Update()
}

func (s *` + pascalNameFormat + `Service) Delete() {
	s.` + pascalNameFormat + `Repository.Delete()
}
`

	// module service test
	ModuleServiceTestTemplate = `package ` + snakeNameFormat + `

import (
	"testing"

	` + importDatabase + `
	` + importRepository + `
)

var ` + camelNameFormat + `Service *` + pascalNameFormat + `Service

func init() {
	db := database.GetDB()
	` + camelNameFormat + `Repository := repository_` + snakeNameFormat + `.New` + pascalNameFormat + `Repository(db)
	` + camelNameFormat + `Service = New` + pascalNameFormat + `Service(` + camelNameFormat + `Repository)
}

func TestGet(t *testing.T) {}

func TestPost(t *testing.T) {}

func TestPut(t *testing.T) {}

func TestDelete(t *testing.T) {}
`

	// service
	ServiceTemplate = `package ` + snakeNameFormat + `

type ` + pascalNameFormat + `Service struct{}

func New` + pascalNameFormat + `Service() *` + pascalNameFormat + `Service {
	return &` + pascalNameFormat + `Service{}
}

func (s *UserService) Get() {}

func (s *UserService) Post() {}

func (s *UserService) Put() {}

func (s *UserService) Delete() {}

`

	// service test
	ServiceTestTemplate = `package ` + snakeNameFormat + `

import (
	"testing"
)

var ` + camelNameFormat + `Service *` + pascalNameFormat + `Service

func init() {
	` + camelNameFormat + `Service = New` + pascalNameFormat + `Service()
}

func TestGet(t *testing.T) {}

func TestPost(t *testing.T) {}

func TestPut(t *testing.T) {}

func TestDelete(t *testing.T) {}
`

	// repository
	RepositoryTemplate = `package ` + snakeNameFormat + `

import (
	` + importDatabase + `
)

type ` + pascalNameFormat + `Repository struct {
	DB *database.DB
}

func New` + pascalNameFormat + `Repository(db *database.DB) *` + pascalNameFormat + `Repository {
	return &` + pascalNameFormat + `Repository{
		DB: db,
	}
}

func (r *` + pascalNameFormat + `Repository) Find() {}

func (r *` + pascalNameFormat + `Repository) Store() {}

func (r *` + pascalNameFormat + `Repository) Update() {}

func (r *` + pascalNameFormat + `Repository) Delete() {}
`

	// repository test
	RepositoryTestTemplate = `package ` + snakeNameFormat + `

import (
	"testing"

	` + importDatabase + `
)

var ` + camelNameFormat + `Repository *` + pascalNameFormat + `Repository

func init() {
	db := database.GetDB()
	` + camelNameFormat + `Repository = New` + pascalNameFormat + `Repository(db)
}
	
func TestFind(t *testing.T) {}

func TestStore(t *testing.T) {}

func TestUpdate(t *testing.T) {}

func TestDelete(t *testing.T) {}
`

	// model
	ModelTemplate = `package ` + snakeNameFormat + `

type ` + pascalNameFormat + ` struct {
	CreateAt string ` + "`json:\"create_at\"`" + `
	UpdateAt string ` + "`json:\"update_at\"`" + `
}
`
)

func snakeToPascal(str string) string {
	words := strings.Split(str, "_")
	for i, word := range words {
		words[i] = strings.ToTitle(word[:1]) + word[1:]
	}
	return strings.Join(words, "")
}

func snakeToCamel(str string) string {
	words := strings.Split(str, "_")
	for i, word := range words {
		if i == 0 {
			continue
		}
		words[i] = strings.ToTitle(word[:1]) + word[1:]
	}
	return strings.Join(words, "")
}

func GetTemplate(templateType enum.GenType, name string) string {
	template := ""
	pascalName := snakeToPascal(name)
	camelName := snakeToCamel(name)

	switch templateType {
	case enum.GenTypeController:
		template = ControllerTemplate
	case enum.GenTypeControllerTest:
		template = ControllerTestTemplate
	case enum.GenTypeModuleController:
		template = ModuleControllerTemplate
	case enum.GenTypeModuleControllerTest:
		template = ModuleControllerTestTemplate
	case enum.GenTypeService:
		template = ServiceTemplate
	case enum.GenTypeServiceTest:
		template = ServiceTestTemplate
	case enum.GenTypeModuleService:
		template = ModuleServiceTemplate
	case enum.GenTypeModuleServiceTest:
		template = ModuleServiceTestTemplate
	case enum.GenTypeRepository:
		template = RepositoryTemplate
	case enum.GenTypeRepositoryTest:
		template = RepositoryTestTemplate
	case enum.GenTypeModel:
		template = ModelTemplate
	}

	template = strings.ReplaceAll(template, snakeNameFormat, name)
	template = strings.ReplaceAll(template, pascalNameFormat, pascalName)
	template = strings.ReplaceAll(template, camelNameFormat, camelName)

	return template
}
