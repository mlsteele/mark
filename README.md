# Mark

Pipe logs through `mark` to see short bursts of log lines grouped together with newlines on pauses.

## Installation

```shell
go get github.com/mlsteele/mark
go install github.com/mlsteele/mark
# make sure $GOPATH/bin is in your $PATH
```

## Usage

```shell
# to tail a file:
tail -f -n0 /path/to/some.log | mark
# for a process (+stderr):
someprog 2>&1 | mark
```
