# hoofli

Generate PlantUML diagrams from Chrome or Firefox network inspections

[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/dnnrly/hoofli)](https://github.com/dnnrly/hoofli/releases/latest)
[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/dnnrly/hoofli/Release%20workflow)](https://github.com/dnnrly/hoofli/actions?query=workflow%3A%22Release+workflow%22)
[![report card](https://goreportcard.com/badge/github.com/dnnrly/hoofli)](https://goreportcard.com/report/github.com/dnnrly/hoofli)
[![godoc](https://godoc.org/github.com/dnnrly/hoofli?status.svg)](http://godoc.org/github.com/dnnrly/hoofli)
[![codecov](https://codecov.io/gh/dnnrly/hoofli/branch/main/graph/badge.svg?token=7SK2qu0f8f)](https://codecov.io/gh/dnnrly/hoofli)

![GitHub watchers](https://img.shields.io/github/watchers/dnnrly/hoofli?style=social)
![GitHub stars](https://img.shields.io/github/stars/dnnrly/hoofli?style=social)
[![Twitter URL](https://img.shields.io/twitter/url?style=social&url=https%3A%2F%2Fgithub.com%2Fdnnrly%2Fhoofli)](https://twitter.com/intent/tweet?url=https://github.com/dnnrly/hoofli)

This tool reads browser HAR files stored on your local disk and transforms them into
PlantUML formatted files. You will need to download PlantUML from https://plantuml.com/
or use the package management tool of your choice

### Installing

```bash
$ go install github.com/dnnrly/hoofli/cmd/hoofli
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

* `install` -- install hoofli from the current working tree
* `build` -- build hoofli
* `clean` -- remove build artifacts from the working tree
* `clean-deps` -- remove dependencies in the working tree
* `test-deps` -- ci target - install test dependencies
* `build-deps` -- ci target - install build dependencies
* `deps` -- ci target - install build and test dependencies
* `test` -- run unit tests with tparse prettifying
* `acceptance-test` -- run acceptance tests on built hoofli
* `ci-test` -- ci target - run unit tests
* `lint` -- run linting
* `release` -- ci target - release hoofli
* `update` -- update dependencies
* `help` -- Show this help.


## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/dnnrly/hoofli/tags). 

## Authors

* **Pascal Dennerly** - *Initial work* - [dnnrly](https://github.com/dnnrly)

See also the list of [contributors](https://github.com/dnnrly/hoofli/contributors) who participated in this project.

## License

This project is licensed under the Apache 2.0 License - see the [LICENSE](LICENSE) file for details

## Acknowledgments

* Important people here
