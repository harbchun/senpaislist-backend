import datetime

date_time_str = "2013-04-13 10:00:00"
date_time_now = datetime.datetime.now()
date_time_obj = datetime.datetime.strptime(date_time_str, '%Y-%m-%d %H:%M:%S')
print(type(date_time_now), type(date_time_obj))

if date_time_obj < date_time_now:
    print('yes')