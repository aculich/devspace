package deploy

import (
	"github.com/devspace-cloud/devspace/pkg/devspace/config/generated"
)

// Interface defines the common interface used for the deployment methods
type Interface interface {
	Delete() error
	Status() ([][]string, error)
	Deploy(generatedConfig *generated.Config, isDev, forceDeploy bool) error
}
