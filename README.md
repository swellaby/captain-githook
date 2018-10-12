# captain-githook 
Configuration driven, cross-platform, git hook utility geared for Go codebases.  

****** Functional, but still in Beta ****** 

[![Linux CI Badge][linux-ci-badge]][linux-ci-url]
[![Mac CI Badge][mac-ci-badge]][mac-ci-url]
[![Windows CI Badge][windows-ci-badge]][windows-ci-url]  

[![Test Results Badge][tests-badge]][sonar-tests-url]
[![Codecov Badge][codecov-badge]][codecov-url]
[![Sonar Quality Gate Badge][sonar-quality-gate-badge]][sonar-url]  

## About
Git hooks are scripts/commands that git can execute on certain events which helps you automate tasks, enforce quality and consistency in your code, and more. `captain-githook` allows you to utilize any and all git hooks in your repository (`commit-msg`, `pre-commit`, `pre-push`, etc.) via a simple configuration file. 

Yes, there are other git hook utilities out there (in fact we love, recommend, and use [husky][husky-url] in our JavaScript/TypeScript projects that use [npm][npm-url]). However, we made `captain-githook` because we wanted a git hook utility for our Go codebases that was cross-platform, easily configurable, and that wasn't *dependent* on another, non-Go framework/runtime.

## Installation
We'll be adding binary releases shortly, but for now you'll need to have your [Go][go-download-url] environment setup and use `go get` i.e.:

```sh
go get -u github.com/swellaby/captain-githook
```

This will ensure that the captain-githook executable is available on your system for initializing your repositories (creating git hook and config files) as well as for executing your defined hook scripts.

## Initialize
Run the `init` sub-command within a git repository to create the git hook files and create the captain-githook config file:

```sh
captain-githook init
```

If you'd prefer to use a different name for the config file, you can specify your desired config file name as a flag (`--config-filename` or `--f`) on the `init` command. For example:

```sh
captain-githook init --config-filename .captain-githookrc.json
```

## Configure
You specify which scripts/commands you want to run for each git hook in the captain-githook configuration file included in your repository.

### Config File Names

## Remove
Not yet implemented

[githooks-docs-url]: https://git-scm.com/docs/githooks
[go-download-url]: https://golang.org/dl/
[husky-url]: https://www.npmjs.com/package/husky
[npm-url]: https://www.npmjs.com/get-npm
[linux-ci-badge]: https://dev.azure.com/swellaby/OpenSource/_apis/build/status/captain-githook/captain-githook-PR-Linux?branchName=master&label=linux%20build
[linux-ci-url]: https://dev.azure.com/swellaby/OpenSource/_build/latest?definitionId=25
[mac-ci-badge]: https://dev.azure.com/swellaby/OpenSource/_apis/build/status/captain-githook/captain-githook-PR-Mac?branchName=master&label=mac%20build
[mac-ci-url]: https://dev.azure.com/swellaby/OpenSource/_build/latest?definitionId=26
[windows-ci-badge]: https://dev.azure.com/swellaby/OpenSource/_apis/build/status/captain-githook/captain-githook-PR-Windows?branchName=master&label=windows%20build
[windows-ci-url]: https://dev.azure.com/swellaby/OpenSource/_build/latest?definitionId=24
[codecov-badge]: https://img.shields.io/codecov/c/github/swellaby/captain-githook.svg
[codecov-url]: https://codecov.io/gh/swellaby/captain-githook
[tests-badge]: https://img.shields.io/appveyor/tests/swellaby/captain-githook.svg?label=unit%20tests
[sonar-quality-gate-badge]: https://sonarcloud.io/api/project_badges/measure?project=swellaby%3Acaptain-githook&metric=alert_status
[sonar-url]: https://sonarcloud.io/dashboard?id=swellaby%3Acaptain-githook
[sonar-tests-url]: https://sonarcloud.io/component_measures?id=swellaby%3Acaptain-githook&metric=tests
