# Godo

## Description
A task manager and todo list integrated into your command line as both text and UI.

## Getting started 
First run:
```
go get github.com/C3NZ/godo
```

And then from there, you can:
```
godo
```
To get started!

## Goals
* To learn more about termui (moved to clui for creating the UI)
* get better at making command line tools with golang (utilized cobra)

## Visions
* Implement gocal for both CLI and UI
* Create stats based on gocal events and godo list events

## How to run
```
godo
```
Will give you basic usage of how to run the commandline and todo list program.
Currently the `godo gocall` command requires a `token.json` file from googles developer
platform in order to connect to the go calender.

In the future, this will be automated and the user will no longer need to go to the
google developer console to get the token file.

## How to use
* `godo *` - General usage and help
* `godo add some task`  - Adds a task to the todo list
* `godo list` - list all of the todo list items 
* `godo do [ID]` - complete a task where ID is the number of the task
* `godo ui` - start the UI application which now visualizes the todo list right within your terminal!

## Resources
[clui](https://github.com/VladimirMarkelov/clui)
[bolt](https://github.com/boltdb/bolt)
[cobra](https://github.com/spf13/cobra)
