# disism saikan 

IPFS-based streaming media service.

## Dev

Go 1.22 is required.

### Run
```shell
export JWT_SECRET="disism.jwt.1234567890.secret@dev"
export SERVICE_ENDPOINT="127.0.0.1:8032"

make run
```

## Plan
In the future, saikan will support activitypub for playlist sharing, so users can subscribe to someone's playlist to receive notifications of updates to that playlist.
Users can subscribe to someone's playlist to be notified of updates to that playlist, and they can also Fork a playlist to add the tracks they want.
