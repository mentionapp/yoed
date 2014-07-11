#Yo'ed

Yo'ed is an extensible hub, written in Go, to dispatch actions when someone yo you.

#Usage

You need the [Go tools](http://golang.org/doc/install) installed.

```go
go get github.com/mentionapp/yoed
cd $GOPATH/src/github.com/mentionapp/yoed
cp config.json.dist config.json
$GOPATH/bin/yoed
```

#Configuration

The configuration is specified in the `config.json` file.

### Listen

Specify the address and port to listen to. Change the callback URL in [Yo's API dashboard](http://developer.justyo.co/) to point to `http://.../yoed`.

### Handlers

This is the enabled handlers list. Each handler has its own configuration. See below for more details about handlers.

#Bundled Handlers

Yo'ed implements two handlers:

* Slack
* Yo-back

### Slack handler

It posts an automatic message on a specified [Slack](https://slack.com) room.
The only configuration parameter is the `webhook_url` URL you get from the Slack's [Incoming WebHooks](https://slack.com/services/new/incoming-webhook) integration settings.

### Yo-back

It simply re-posts a YO to users who YO you.
The only configuration parameter is the `api_token`.

### Custom handlers

Custom handlers can be added easily. They must be compatible with the `yoedHandler` interface:

```go
type yoedHandler interface {
	Handle(username string)
}
```
