# LogFlow

A logsurfer replacement inspired by data-flow based pardigms.

using goflow version 0.1 because the current
development branch is unstable and undocumented.


## Setup
```
go get github.com/trustmaster/goflow
go get github.com/gorilla/websocket
go get github.com/nu7hatch/gouuid
cd github.com/trustmaster/goflow
git checkout 0.1
```

## poc0
This is the first simple network that actually runs for me.
It has only two components: ReadFile which reads a text file
line by line (accepting filenames at its input), and Print,
which prints lines of text to stdout. Run it:
```
cd go/src/terreactive.ch/LogFlow
go run mytest/poc0.go
