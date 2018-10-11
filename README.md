# captain-githook 
git hook utility for Go codebases  
****** Functional, but still in Beta ****** 

[![Linux CI Badge][linux-ci-badge]][linux-ci-url]
[![Mac CI Badge][mac-ci-badge]][mac-ci-url]
[![Windows CI Badge][windows-ci-badge]][windows-ci-url]  

[![Test Results Badge][tests-badge]][sonar-tests-url]
[![Codecov Badge][codecov-badge]][codecov-url]
[![Sonar Quality Gate Badge][sonar-quality-gate-badge]][sonar-url]  

## Installation
We'll be adding binary releases shortly, but for now you'll need to have your [Go][go-download-url] environment setup and use `go get` i.e.:

```sh
go get -u github.com/swellaby/captain-githook
```

## Initialize
Run the `init` sub-command within your git repository to initialize the git hooks and create the captain-githook config file:

```sh
captain-githook init
```

 
[linux-ci-badge]: https://swellaby.visualstudio.com/OpenSource/_apis/build/status/captain-githook/captain-githook-PR-Linux?branchName=master&label=linux%20build
[linux-ci-url]: https://swellaby.visualstudio.com/OpenSource/_build/latest?definitionId=25
[mac-ci-badge]: https://swellaby.visualstudio.com/OpenSource/_apis/build/status/captain-githook/captain-githook-PR-Mac?branchName=master&label=mac%20build
[mac-ci-url]: https://swellaby.visualstudio.com/OpenSource/_build/latest?definitionId=26
[windows-ci-badge]: https://swellaby.visualstudio.com/OpenSource/_apis/build/status/captain-githook/captain-githook-PR-Windows?branchName=master&label=windows%20build
[windows-ci-url]: https://swellaby.visualstudio.com/OpenSource/_build/latest?definitionId=24
[codecov-badge]: https://img.shields.io/codecov/c/github/swellaby/captain-githook.svg
[codecov-url]: https://codecov.io/gh/swellaby/captain-githook
[tests-badge]: https://img.shields.io/appveyor/tests/swellaby/captain-githook.svg?label=unit%20tests
[sonar-quality-gate-badge]: https://sonarcloud.io/api/project_badges/measure?project=swellaby%3Acaptain-githook&metric=alert_status
[sonar-url]: https://sonarcloud.io/dashboard?id=swellaby%3Acaptain-githook
[sonar-tests-url]: https://sonarcloud.io/component_measures?id=swellaby%3Acaptain-githook&metric=tests
[go-download-url]: https://golang.org/dl/