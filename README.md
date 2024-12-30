# go-gen-root

Composition root generator

## Usage

- Install the tool
  - `go  install -v github.com/sealbro/go-gen-root@latest`
- Add `@inject` comment to your struct constructor function
  - Function must return only reference from struct instance
- Run generator with next parameters
  - TODO: Add parameters description

## Features

- [x] Generate composition root App struct
- [x] Sorted variables from zero parameter to many
- [x] Throw panic if constructor function has not injected parameters
- [ ] Use instances initializes manually and register them in App struct
- [ ] Return `error` from constructor function and correctly handle it
- [ ] Configure generator tool with flags (module name, output path, package name)