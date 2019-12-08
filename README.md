# LogFlow

A logsurfer replacement inspired by data-flow based paradigms.

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
which prints lines of text to stdout.

            +---------------------+           +---------------------+
            | ReadFile            |           | Print               |
            |                     |           |                     |
            |               Error |           |                     |
    In ---->+ Filename            |           |                     |
            |                Line +---------->+ In                  |
            +---------------------+           +---------------------+

Run it:
```
cd go/src/terreactive.ch/LogFlow
go run mytest/poc0.go
```

## poc1
Builds on poc0 and adds more components, allowing the construction
of networks that really processes real logs.

The LineToStream component serves as an input of real-world logs into our
Log Flow system. It accepts lines of string (log line) and converts them into
LogStream objects, the base class of objects that are used throughout the Log
Flow system. Component LineToStream simply assigns the log line to the Raw
field of LogStream, no further parsing is done. The incoming time is assigned
from the wall clock.

The StreamDump component prints the entire contents of the LogStream object
in a multi-line, human readable form. This is meant for developing and 
debugging.

The poc1 executable executes this simple pipeline:

            +---------------------+   +-----------------+   +-----------------+
            | ReadFile            |   | LineToStream    |   | StreamDump      |
            |                     |   |                 |   |                 |
            |               Error |   |                 |   |                 |
    In+---->+ Filename            |   |                 |   |                 |
            |                Line +-->+ In          Out +-->+ In              |
            +---------------------+   +-----------------+   +-----------------+

