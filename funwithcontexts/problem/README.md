# Problem statement
We require an uptime check utility \
The util should asynchronously validate each url in a provided list every x seconds \
We should be able to stop the util, preventing further uptime checks.

Example URLs that could be used:
```
aroundhome.de
google.com
facebook.com
prosiebensat1.de
```

For the first iteration use an auto-canceling context (timeout/deadline) to signal when to stop.

## Bonus
- Handle termination signals like `CTRL+C` 
- Gracefully shutdown (stop each concurrent routine individually, exit once all are stopped)
