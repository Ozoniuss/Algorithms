DURATION = 30


# we need this for comparison purposes.
def convert_hours_to_minutes(hour):
    hours, minutes = hour.split(":")
    return int(hours)*60 + int(minutes)

def convert_minutes_to_hours(mins):
    hours = mins // 60
    minutes = mins % 60
    if minutes < 10:
        return f"{hours}:0{minutes}"
    else:
        return f"{hours}:{minutes}"

print(convert_hours_to_minutes("10:30"))
print(convert_minutes_to_hours(700))

def read(file):
    meetings = []
    with open(file) as f:
        for line in f:
            hours = line.strip().split(', ')
            meetings.append(hours)

    interval = meetings.pop()
    start = interval[0]
    start = start.split(': ')[1]
    stop = interval[1]

    interval = [start, stop]

    return meetings, interval

m1, i1 = read("calendar1.in")
m2, i2 = read("calendar2.in")

def find_avaialbe_timeline(scheduled_meetings, timeline):
    available_meeting_times = []
    current_time = timeline[0]
    for schedule in scheduled_meetings:
        start = schedule[0]
        end = schedule[1]

        if current_time == start:
            current_time = end
            continue

        if (convert_hours_to_minutes(start) - convert_hours_to_minutes(current_time) >= DURATION):
            available_meeting_times.append([current_time, start])
        current_time = end

    if (convert_hours_to_minutes(timeline[1]) - convert_hours_to_minutes(current_time)) >= DURATION:
        available_meeting_times.append([current_time, timeline[1]])

    return available_meeting_times

at1 = find_avaialbe_timeline(m1, i1)
at2 = find_avaialbe_timeline(m2, i2)
print(f"at1: {at1}")
print(f"at2: {at2}")
def merge_timetables(available1, available2):
    we_can_meet=[]
    for interval1 in available1:
        for interval2 in available2:
            start1, end1 = interval1[0], interval1[1]
            start2, end2 = interval2[0], interval2[1]
            meeting_start_time = max(convert_hours_to_minutes(start1), convert_hours_to_minutes(start2))
            meeting_end_time = min(convert_hours_to_minutes(end1), convert_hours_to_minutes(end2))

            if meeting_start_time < meeting_end_time:
                we_can_meet.append([convert_minutes_to_hours(meeting_start_time),
                                   convert_minutes_to_hours(meeting_end_time)])

    return we_can_meet

print(merge_timetables(at1, at2))