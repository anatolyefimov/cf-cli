# Codeforces command-line interface
 [![anatolyefimov](https://circleci.com/gh/anatolyefimov/cf-cli.svg?style=svg)]()


## Install
```
$ export PATH=$PATH:$(go env GOBIN)
$ go build $(go env GOBIN)/cf
```
## Usage

```
сf login                               Login to Codeforces. Note that password is entered in silence mode.
cf enter <contest-id>                  Entering the contest with this id.
сf submit <id|index> <path-to-source>  The command send the solution to the problemset, 
                                       if before that the user entered the contest using 'cf enter', 
                                       otherwise  it will send by the index of the corresponding contest.
cf help                                Get help
```