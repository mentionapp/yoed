# YO'ed

YO'ed is hub written in Go to dispatch actions when someone YO you.

## Usage

You need the [Go tools](http://golang.org/doc/install) installed.

```go
go get github.com/mentionapp/yoed
cd $GOPATH/src/github.com/mentionapp/yoed
cp config.json.dist config.json
$GOPATH/bin/yoed
```

Then change the callback URL in [Yo's API dashboard](http://developer.justyo.co/) to point to `http://your.server:port/yoed`.

## Configuration

The configuration is specified in the `config.json` file.

### `"listen"`

Specifies the address and port to listen on, e.g. `0.0.0.0:12345` to listen on port 12345 on any address.

### `"handlers"`

This is the enabled handlers list. Each handler has its own configuration. See below for more details about handlers.

#### Available Handlers

YO'ed comes with a few handlers:

##### Slack handler

Uses [Slack](https://slack.com)'s [Incoming WebHooks](https://slack.com/services/new/incoming-webhook) integration to YO in a room.

The only configuration parameter is the `webhook_url` URL you get while setting up the webhook.

##### YO-back

Sends back a YO to users who YO you.

The only configuration parameter is the `api_token`.

##### Custom handlers

Custom handlers can be added easily. They only have to be compatible with the `yoedHandler` interface:

```go
type yoedHandler interface {
	Handle(username string)
}
```
