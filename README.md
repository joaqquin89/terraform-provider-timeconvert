# terraform-provider-time-convert

# the purpose in this provider is to convert from hours to milliseconds , nanoseconds and seconds

the usage in for this provider are the nexts steps:


- the resources can be used in this provider are timeconvert_nanoseconds , timeconvert_seconds or timeconvert_nanoseconds

-  time must be setting in Xh(x is an integer that means hours ) or can be XhYm (Y means the minutes), i.e: 3h means 3 hours, 3h20m means 3 hours and 20 minutes.

- the output value for nanoseconds is ns, for seconds is secs and
for milliseconds is  ms

i.e:

    resource "timeconvert_nanoseconds" "main" {
        time = "1h"
    }

    output "seconds" {
    value = timeconvert_nanoseconds.main.ns
    }

the example above, try to get the amount of nanoseconds that exists in 1 hour