# Go basic commands

- Create a module using go mod init <module-name - location on the internet where you'll go to download that module>.
Modules are the unit of distribution and versioning of software

- Go programs are organized into packages. A package is a collection of source files in the same directory that are compiled together. Functions, types, variables, and constants defined in one source file are visible to all other source files within the same package.

- A module is a collection of related Go packages that are released together. A module is defined by a directory tree of Go source files with a go.mod file in the tree's root directory. go.mod declares the module path: the import path prefix for all packages within the module.

- A package's import path is its module path joined with its subdirectory within the module.

go install learn/testing/gotests/go-basics

- This command builds the go files in <learn/testing/gotests/go-basics> package, producing an executable binary. It then installs that binary as $HOME/go/bin/go-basics

- For convenience, go commands accept paths relative to the working directory, and default to the package in the current working directory if no other path is given. So in root directory (learn/testing/gotests), the following commands are all equivalent:

$ go install learn/testing/gotests/go-basics
$ go install ./go-basics

or same as the following command in learn/testing/gotests/go-basics directory
$ go install

- go build
compiles the packages, along with their dependencies, but it doesn't install the results. Instead it saves the compiled package in the local build cache. If we want our program to run again, we don’t have to compile the program again, we simply run the executable file.

- The go install command compiles and installs the packages, along with their dependencies. Basically you can use go build as a check that the packages can be built (along with their dependencies) while go install also (permanently) installs the results in the proper folder ($GOPATH/pkg).

- go build -o DirectoryPath
parameter -o followed by go build allows to output binary to a specified location

- go run <name of the Go program>
combines both the compilation and execution of code. This allows to quickly check the output of updated code.
Unlike go build, go run will NOT create an executable file in our current folder.

- we can also sun the program by just typing the binary executable name for this package, i.e. use commannd go-basics to run hello.go
- go mod tidy
command adds missing module requirements for imported packages and removes requirements on modules that aren't used anymore. Module dependencies are automatically downloaded to the pkg/mod subdirectory of the directory indicated by the GOPATH environment variable (github.com/google/go-cmp in this example)

- The downloaded contents for a given version of a module are shared among all other modules that require that version, so the go command marks those files and directories as read-only. To remove all downloaded modules, you can pass the -modcache flag to go clean:
$ go clean -modcache

- GOROOT is for compiler/tools that comes from go installation
- GOPATH is for your own go projects / 3rd party libraries (downloaded with "go get"). If the environment variable is unset, GOPATH defaults to a subdirectory named "~/go". It is used to resolve import statements when not using go.mod. It is used to store downloaded source code in GOPATH/pkg and compiled commands in GOPATH/bin.

- The package `main` tells the Go compiler that the package should compile as an executable program instead of a shared library. During build, main() func inside main package is searched and teh built binary starts there. It is the entry point of the program. Each individual program only has one main package
