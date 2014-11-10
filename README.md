# Intercom-Go

Offical bindings for the [Intercom](https://www.intercom.io) API

## Install

`go get github.com/intercom/intercom-go`

Dependencies:

* [https://github.com/franela/goreq](https://github.com/franela/goreq)

## Usage

### Getting a Client

The first step to using Intercom's Go client is to create a client object, using your App ID and Api Key from your [settings](http://app.intercom.io/apps/api_keys).

```go
ic := intercom.GetClient("appID", "apiKey")
```

This client can then be used to make requests.

#### Client Options

The client can be configured with different options by calls to `ic.Option`:

```go
ic.Option(intercom.TraceHttp(true)) // turn http tracing on
ic.Option(intercom.BaseUri("http://intercom.dev")) // change the base uri used, useful for mocking
```

or combined:

```go
ic.Option(intercom.TraceHttp(true), intercom.BaseUri("http://intercom.dev"))
```

### Events

#### New

Events are created using an `EventParams` structure:

```go
err := ic.Events.New(&intercom.EventParams{
  UserId:     "27",
  EventName: "bought_item",
  Metadata:  intercom.CreateMetadata().Add("item_id", 22).Add("item_name", "PocketWatch"),
})
```

One of `Id`, `UserId`, or `Email` is required.

`EventName` is required.

`Metadata` is optional, and can be constructed using the helper as above, or as a passed `map[string]interface{}`.

### Errors

All operations can return errors in addition to structured data. These errors should be checked.