<img align="right" float="right" src="https://user-images.githubusercontent.com/71837281/137636053-f090fa9a-7a70-4028-a95b-89810263caac.png" height="256px" width="237px">

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

### Use it yourself
Do you have a similar goal or want a daily reminder to be ~~productive~~? 
1. **Fork** this repository and change the username to your username in the [workflow](https://github.com/darmiel/365-days-notifier/blob/master/.github/workflows/check.yaml) file.
    * You can also customize the cronjob here
2. Create the **secrets** `PUSH_APP_KEY` and `PUSH_RECIPIENT_KEY` in the repository settings. 
    * The secret `GH_PAT` is only required if you want to include private activity.
3. Go to **Actions** tab and enable actions
4. Done!

### Pushover
If no activity is found that day, I get bugged via push message on my phone, for which I use the super duper app [Pushover](https://pushover.net/):
![Push+](https://user-images.githubusercontent.com/71837281/137634896-52366845-b7ff-461a-8d96-ec2b54481269.jpg)
