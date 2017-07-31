# too

the opposite of `tee`.

```
-> command 1 stream ─┐
-> command 2 stream ─┤
                     └─ stdout/stderr/SIGINT to kill both
```

# why?

When you need to start 2 file watchers, and kill them by 1 signal, like this.

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
% too
> rails server # return key
> npm start-webpack # return key
> # return key again

[0] ... # Rails log here
[1] ... # npm log here

# To kill both, just Ctrl+C once!
```
