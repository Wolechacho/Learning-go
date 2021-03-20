//+build wireinject
package dependencyInjection

import (
	"github.com/google/wire"
)

//this should be per client
func InitializePersonResource() *PersonResource {
	wire.Build(NewDatastore, NewLogger, NewPersonResource)
	return &PersonResource{}
}
