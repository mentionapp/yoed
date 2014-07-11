#Yo'ed

Yo'ed is an extensible hub, written in Go, to dispatch actions when someone yo you.

#Usage

You need the [Go runtime](http://golang.org/doc/install) installed.

```go
go get github.com/mentionapp/yoed
cd $GOPATH/src/github.com/mentionapp/yoed
cp config.json.dist config.json
$GOPATH/bin/yoed
```

#Configuration

The configuration is specified in the `config.json` file.

### Listen

Specify the URL and port to listen to. This is a URL you specify in your [Yo's API dashboard](http://developer.justyo.co/).

### Handlers

This is the enabled handlers list. Each handler has its own configuration. See below for more details about bundled handlers.

#Bundled Handlers

Yo'ed is shipped with two default handlers:

* Slack
* Yo-back

All Yo-ed's handlers must satisfy this simple Go interface:
```go
type yoedHandler interface {
	Handle(username string)
}
```

### Slack handler

It posts an automatic message on a specified [Slack room](slack.com).
The only configuration parameter is the `webhook_url` URL you get from the Slack's settings.

### Yo-back

It simply re-posts a yo to the user who has Yo you. 
The only configuration parameter is the `api_token`.
