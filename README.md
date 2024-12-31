# go-gen-root

A simple example of an automated composition root generator for Go projects.

## Usage

- Install the tool
  - `go install -v github.com/sealbro/go-gen-root@latest`
- Add `@inject` comment to your struct constructor function
  - Function must return only reference from struct instance
- Run generator with next parameters as example
  - `go-gen-root generate -p examples/app-with-deps -m github.com/sealbro/go-gen-root -o examples/app-with-deps/cmd/app/di.go`

## Features

- [x] Generate composition root App struct
- [x] Sorted variables from zero parameter to many
- [x] Throw panic if constructor function has not injected parameters
- [ ] Register interface as declaration and bind with reference, additionally check if interface is implemented
- [ ] Use instances initializes manually and register them in App struct
- [ ] Return `error` from constructor function and correctly handle it
- [ ] Configure generator tool with flags (module name, output path, package name)