# hoofli

Generate PlantUML diagrams from Crome network inspections

[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/dnnrly/hoofli)](https://github.com/dnnrly/hoofli/releases/latest)
[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/dnnrly/hoofli/Release%20workflow)](https://github.com/dnnrly/hoofli/actions?query=workflow%3A%22Release+workflow%22)
[![report card](https://goreportcard.com/badge/github.com/dnnrly/hoofli)](https://goreportcard.com/report/github.com/dnnrly/hoofli)
[![godoc](https://godoc.org/github.com/dnnrly/hoofli?status.svg)](http://godoc.org/github.com/dnnrly/hoofli)

![GitHub watchers](https://img.shields.io/github/watchers/dnnrly/hoofli?style=social)
![GitHub stars](https://img.shields.io/github/stars/dnnrly/hoofli?style=social)
[![Twitter URL](https://img.shields.io/twitter/url?style=social&url=https%3A%2F%2Fgithub.com%2Fdnnrly%2Fhoofli)](https://twitter.com/intent/tweet?url=https://github.com/dnnrly/hoofli)


## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

What things you need to install the software and how to install them

```
Give examples
```

### Installing

```bash
$ git clone http://github.com/dnnrly/hoofli.git
$ cd hoofli
$ make install
```

### Running Unit Tests

```bash
$ make test
```

### Running Acceptance tests

```bash
$ make deps
$ make build acceptance-test
```

## Important `make` targets

* `deps` - downloads all of the deps you need to build, test, and release
* `install` - installs your application
* `build` - builds your application
* `test` - runs unit tests
* `ci-test` - run tests for CI validation
* `acceptance-test` - run the acceptance tests
* `lint` -  run linting
* `update` - update Go dependencies
* `clean` - clean project dependencies
* `clean-deps` - remove all of the build dependencies too


## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/dnnrly/hoofli/tags). 

## Authors

* **Your name here** - *Initial work* - [dnnrly](https://github.com/dnnrly)

See also the list of [contributors](https://github.com/dnnrly/hoofli/contributors) who participated in this project.

## License

This project is licensed under the Apache 2.0 License - see the [LICENSE](LICENSE) file for details

## Acknowledgments

* Important people here
