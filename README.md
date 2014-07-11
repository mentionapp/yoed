YO'ed
=====

YO'ed is a simple server to get notifyed when someone YO's your YOAPI account, and handle the YO in some way.

Handlers
--------

YO'ed currently supports the following handlers:

 - slack: posts on a slack channel, using the 'incomming webhooks' integration
 - yoback: sends a YO to the sender

Usage
-----

```
$ go get github.com/mentionapp/yoed
$ cp config.dist.json ./config.json
$ vim ./config.json
$ $GOPATH/bin/yoed
```

