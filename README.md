# tpcli

tpcli is small cmd tool for accessing target process entity data inside terminal. This tool was build while learning concepts of golang, so it is very simple tool that is for now restricted only to GET methods.

[Target Process](http://www.targetprocess.com/) is Visual project management software, that helps you Plan and track any process, including scrum, kanban and your own.


## Installation

Compile from source or [download](Downloads/) binary for your platform.

## Usage

To use tpcli you have to provide username, password and url data for your Target Process api endpoint. This can be done with flags or with environment variables.

### Setting account data with flags

```sh
$ tpcli --username jack@acme.com --password AstrongPass --url http://acme.tpondemand.com/api/v1/ .....
```

### Setting account data with Environment variables

```sh
$ export TPCLI_USERNAME=jack@acme.com
$ export TPCLI_PASSWORD=AstrongPass
$ export TPCLI_URL=http://acme.tpondemand.com/api/v1/
```
And the You can use tpcli without global flags. For example:

```sh
$ tpcli bug 808
```

### Getting Entity Data
For now only Bug, User Story and Task entities are supported. Entity data can be displayed in 3 different formats s (small), m (medium) and l (large). Format can be spacified using --tempate flag

#### Getting bug data
```sh
$ tpcli bug 808
$ tpcli bug 808 -template l
$ tpcli --username jack@acme.com --password AstrongPass --url http://acme.tpondemand.com/api/v1/ bug 808 -t m
```

#### Getting user story data
```sh
$ tpcli story 102
$ tpcli story 102 -template l
$ tpcli --username jack@acme.com --password AstrongPass --url http://acme.tpondemand.com/api/v1/ story 102 -t m
```

#### Getting task data
```sh
$ tpcli task 197
$ tpcli task 197 -template l
$ tpcli --username jack@acme.com --password AstrongPass --url http://acme.tpondemand.com/api/v1/ task 197 -t m
```

## To do
- [ ] Write tests
- [ ] Adjust terminal output
- [ ] Add more data and templates
- [ ] Add POST and PUT requests


## License

MIT

## Author

Ivan Padavic (@ipadavic)
