# walkthrough demo


We are using Azure VM (for this i am using my personal dev env).


Azure mounts debugfs automatically.

## prelog
`sudo -s` 
everything is sudo in this demo

`cd /sys/kernel/debug/tracing/`
this is tracefs

`tree -d 1 ./events`
events are defined per `subsystem` for example all kernel block i/o are under `./events/block`


## explore one event

`tree -d 2 ./events/sched`
lists all the events under scheduler (YMMV).

`cat ./events/sched/sched_wakeup/format`
each event has a format. The format is a structure. You can view it on `format` file under the event 
definition.


`echo 1 > ./events/sched/enable`
enables all sched events  (0 to disable)

`echo 1 > ./events/sched/<event>/enable`
enables specific event

## listing and filtering

there are two ways (for now) to `listen` to events via file or via a pipeline like file.

> split screen

> because of the high throughput of events, we need to filter the output. Let us filter
> keep in mind that we are filtering at the entire category. The field we choose 
> need to be in every data structure for all event types.

split 1:
`cat ./trace_pipe` # live events

`echo 'pid == 5249' > ./events/sched/sched_wake/filter` # pid filter
`echo 1 > ./events/sched/enable` # enable all scheduler events

.. wait 1 or 2 sec
`echo 0 > ./events/sched/enable` # disable, pipeline still spitting data because of how big the buffer is

buffer size can be configured by writing to files at the `/sys/kernel/debug/tracing` directory.

> `pid` is specifically available as a global filter via 
```
# cd /sys/kernel/debug/tracing
# echo $$ > set_event_pid
```

## Bonus Rounds: Triggers and histograms. 
Triggers are actions (a la cart, not custom) that runs when an events occurs (filter applied). triggers can enable or disable other traces, get stacktrace. triggers can be used to create histograms. 

details: https://www.kernel.org/doc/Documentation/trace/events.txt
to understand more about adjusting buffers, trace markers, trace files etc https://www.kernel.org/doc/Documentation/trace/ftrace.txt

