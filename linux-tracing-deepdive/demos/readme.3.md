# dynamic tracing

> you can do the probes for kernel as well. This exercise is left for the reader. For examples
go to: http://www.brendangregg.com/perf.html




## listing
`perf probe -x <PATH> -F`
listing all symbols in an executable

The above is equal to `objdump -t <exec path>`
it basically reads the symbol table.

> You won't be able to do that for stripped binaries.


> while go apps are compiled to native code. It is significantly optimized by default. Hence our sample has `//go:noinline` paragma to stop inlining ohterwise we won't be able to see our symbols. You can achieve the save via gcflags `-n -l` but that goes for the entire code you are compiling



## recording a uprobe

`perf probe -x <exec> --add main.fA` # this enables the probe
`perf record -e probe_main:main -aRg <exec>` #record (you can also record for a pid)
`perf report -g --stdio` # report to stdio and show graph (you can report for a pid)




## more on uprobe
> for this to work, your exec needs to be compiled with debugging info `-g with gcc` 


(the C sample app has a `i` argument) 

`perf probe -x <C executable> --add 'fC i'`
add probe to function (notice how we are adding the parameter name, there is plenty more you can do there around return values, filters... for more RTFM

`perf record -aRg -e probe_main:fC <c executable>`
record..

`perf report -g --stdio` 
report


## removing probes

`perf probe -l` # lists all
`perf probe --del <probe name>` # removes a prob

> you can also use glob 
