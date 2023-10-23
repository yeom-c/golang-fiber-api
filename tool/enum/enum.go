package enum

type GenType string

const (
	GenTypeController           GenType = "controller"
	GenTypeControllerTest       GenType = "controllerTest"
	GenTypeModuleController     GenType = "moduleController"
	GenTypeModuleControllerTest GenType = "moduleControllerTest"
	GenTypeService              GenType = "service"
	GenTypeServiceTest          GenType = "serviceTest"
	GenTypeModuleService        GenType = "moduleService"
	GenTypeModuleServiceTest    GenType = "moduleServiceTest"
	GenTypeRepository           GenType = "repository"
	GenTypeRepositoryTest       GenType = "repositoryTest"
	GenTypeModel                GenType = "model"
)
