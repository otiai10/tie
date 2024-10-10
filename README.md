# tie

[![Go](https://github.com/otiai10/tie/actions/workflows/go.yml/badge.svg)](https://github.com/otiai10/tie/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/otiai10/tie/branch/master/graph/badge.svg)](https://codecov.io/gh/otiai10/tie)
[![Go Report Card](https://goreportcard.com/badge/github.com/otiai10/tie)](https://goreportcard.com/report/github.com/otiai10/tie)

The opposite of `tee`, merges multiple command io stream and controls like 1 command.

```
-> command 1 stream ─┐
-> command 2 stream ─┤
                     └─ stdout/stderr/SIGINT to kill both
```

# install

```sh
% go get -u github.com/otiai10/tie
# then just hit `tie`
```

# why?

When you need to start 2 file watchers, and kill them at the same time, like this

```sh
% nohup rails server &
% nohup npm start-webpack &
# then
% pkill rails
% pkill start-webpack
# <- annoying :(
```

# usage

```sh
% tie
> rails server # return key
> npm start-webpack # return key
> # return key again

[0] ... # Rails log here
[1] ... # npm log here

# To kill both, just Ctrl+C once!
```

# examples

by interactive mode

![i](https://user-images.githubusercontent.com/931554/28806719-843a9ffe-76ac-11e7-80c0-13b378ecf7c4.png)

by one-line mode

![o](https://user-images.githubusercontent.com/931554/28806757-aef046c2-76ac-11e7-8d47-be1f2b299fbb.png)

in both ways, you can kill all the commands by just 1 Ctrl+C
