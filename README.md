# Installation
This application is intended to be run on a re-occurring basis using cron, but
cron is a bit peculiar by default:
- it uses /bin/sh not /bin/bash
- most environment variables are not available

1. Open up crontab editor

```
crontab -e
```

2. Define HIPCHAT_TOKEN & add an entry for hibye

```crontab
HIPCHAT_TOKEN=<YOUR-HIPCHAT-TOKEN>

* * * * * ~/bin/go/hibye
```
