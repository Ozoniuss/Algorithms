# Calendar organizer

Given a list of schedules for two people, find the time intervals in which they could meet. Assume intervals are given in correct order, there are no conflicts between intervals, and the intervals are between the correct schedule given as the `interval`.

A sample input can be found in the calendar files: [calendar 1](calendar1.in) for the first person and [calendar 2](calendar2.in) for the second person.

```
# Person 1

10:00, 11:30
12:30, 14:30
14:30, 15:00
16:00, 17:00
interval: 10:00, 18:30
```

```
# Person 2

9:00, 10:30
12:00, 13:00
16:00, 18:00
interval: 9:00, 20:00
```

The expected output with the possible schedule times can be found in file [expected output](expected-output.txt).

```
['11:30', '12:00'], ['15:00', '16:00'], ['18:00', '18:30']
```

This problem is taken from a mock interview from the popular channel Tech with Tim: https://www.youtube.com/watch?v=kbwk1Tw3OhE