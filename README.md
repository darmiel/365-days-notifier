# 365-days-notifier
For more than half a year, I've been trying to get at least in touch with code or the open source community every day. 
This little tool is supposed to remind me every day to fulfill this goal, should I not have done anything ~~productive~~ on GitHub that day.

For this, this tool is called twice a day via a [GitHub action](https://github.com/darmiel/365-days-notifier/blob/master/.github/workflows/check.yaml) per `go run`, 
so I don't have to worry about having my own server and the uptime of the service.

### How?
The tool calls the GitHub API on startup (https://api.github.com/users/darmiel/events) and then checks if there is a relevant event in the GitHub feed today.
This API returns the last X events of a user:
```json
[
  {
    "id": "18459505751",
    "type": "CreateEvent", // productive!
    ...
    "public": true,
    "created_at": "2021-10-15T21:45:00Z"
  },
  {
    "id": "18424454329",
    "type": "WatchEvent",
    ...
    "public": true,
    "created_at": "2021-10-13T20:03:22Z"
  },
  {
    "id": "18368518379",
    "type": "PushEvent", // productive
    ...
    "public": true,
    "created_at": "2021-10-10T13:36:49Z"
  },
  ...
```

---

### Pushover
If no activity is found that day, I get bugged via push message on my phone, for which I use the super duper app [Pushover](https://pushover.net/):
![Push+](https://user-images.githubusercontent.com/71837281/137634896-52366845-b7ff-461a-8d96-ec2b54481269.jpg)
