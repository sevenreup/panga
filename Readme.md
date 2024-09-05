# Panga

Panga is a simple and efficient (not really) scaffolding tool designed to help developers (me) quickly set up project structures and generate files for various types of projects (just like 2).

## Features

- Scaffold different project types
- Generate individual files based on templates
- Easy-to-use command-line interface
- Customizable templates

## Usage

For example creating a simple scaffold for a golang project

```bash
panga new go --template=go-sveltekit --name=hello --module=github.com/sample/hello
```

## Template

The templates are located in the `./templates/` folder (Should also have a way to add custom paths for templates).

The template needs `template.yaml` that defines how the scaffold is generated.

### Templating scratch pad
- should have a presets section of the params
- should be able to run commands (structure this like github actions)
