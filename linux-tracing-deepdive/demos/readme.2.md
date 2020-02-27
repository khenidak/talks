# meet perf

## quick commands

`perf list`

lists everything, events are listed by `category:event`. This includes software events
tracepoints, metrics etc.


`perf list block:*`
filtering becomes really handy (or just `grep` it). the above filters for block i/o events

## state and top
`perf state ./go_app/main` # PMU is h/w events, arch specific
quick and dirty stating a process. You can do with pids as well


`perf top -a --pid=<PID>`
top like view of a single process. kernel and userspace


## recording and viewing
`perf record -e sched:sched_wakeup -a (executable)`
recording same event as previous demo. 

`perf script`

## tracing libc

`perf probe sdt_libc:memory_sbrk_more` # You don't need that on new Perf 
This enables the user prop

`perf report`
mostly you will be using `perf script` or `perf report --stdio`


