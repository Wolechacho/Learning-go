Module is go support for dependency management.
A module by definition is a collection of related packages with go.mod at its root

The go.mod file defines the
Module import path.
Dependency requirements of the module for a successful build. It defines both project’s dependencies requirement and also locks them to their correct version

go.sum file which contains the cryptographic hash of bits of all project’s dependent modules

Modules was introduced v1.13

When GO111MODULE=off, then go get will behave in the old way where it will download the dependency in the $GOPATH/src folder.

When GO111MODULE=on, then go get will behave in a new way and all the modules will get downloaded in the $GOPATH/pkg/mod/cache folder with versioning.

When GO111MODULE=auto, then

When running go get outside the $GOPATH/src folder, then will behave as if it is GO111MODULE=on
When running go get inside the $GOPATH/src folder, then will behave as if it is GO111MODULE=off

On running ‘go install‘ it creates the binary with the name as the last part of the module import that contains the .go file belonging to the main package

Within the same package, all variables, functions, constants are accessible between different .go files belonging to the package

package name != directory name
YOU USE THE PACKAGE NAME TO IMPORT NOT THE DIRECTORY NAME

init() function is a special function that is used to initialize global variables of a package. These functions are executed when the package is initialized

Init function is optional
Init function does not take any argument
Init function does not have any return value.
Init function is called implicitly. Since it is called implicitly, init function cannot reference it from anywhere.
There can be multiple init() functions within the same source file

go mod download - This command is used to pre download all the dependency before the application is run


a blank import of a package is used when

The imported package is not being used in the current program
But we intend to import that package so that the init function in the GO source files belonging to that package can be called and initialization of variables in that package can be done properly

Example
So basically a blank import is used when a package is solely imported for its side effects. As an example MySQL package is used as a blank import for its side-effect of registering the MySQL driver as a database driver in the init() function of MySQL package, without importing any other functions


COMMANDS
go mod download
go mod tidy
go install
go get github.com/pborman/uuid

