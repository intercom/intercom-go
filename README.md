# Intercom-Go

Offical bindings for the [Intercom](https://www.intercom.io) API

## Install

`go get github.com/intercom/intercom-go`

Dependencies:

* [https://github.com/franela/goreq](https://github.com/franela/goreq)
* [https://github.com/google/go-querystring](https://github.com/google/go-querystring)

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
  UserId:    "27",
  EventName: "bought_item",
  CreatedAt: int32(time.Now().Unix()),
  Metadata:  intercom.CreateMetadata().Add("item_id", 22).Add("item_name", "PocketWatch"),
})
```

* One of `Id`, `UserId`, or `Email` is required.
* `EventName` is required.
* `CreatedAt` is optional, must be an integer representing seconds since Unix Epoch. Will be set to _now_ unless given.
* `Metadata` is optional, and can be constructed using the helper as above, or as a passed `map[string]interface{}`.

### Notes

#### New

```go
note, err := c.Notes.New(&intercom.NoteParams{
  UserId: "27",
  Body: "Unicorn Developer",
  AdminId: "1457"
})
```

* One of `Id`, `UserId`, or `Email` is required.
* `Body` is required.
* `AdminId` is optional.

#### Find Single

```go
note, _ := c.Notes.Find(&intercom.NoteParams{
  Id: "87",
})

log.Printf(note.Author)
// [intercom] admin { id: "87", name: "Jamie", email: "somemail@intercom.io" }
```
  
  * One of `Id`, `UserId`, or `Email` is required.
  * `Body` is required.
  * `AdminId` is optional.

### Users

#### New

```go
usr, err := c.Users.New(&intercom.UserParams{
  UserId:          "27",
  Name:            "Gopher",
  RemoteCreatedAt: 1416750500,
  CustomData:  intercom.CreateCustomData().Add("vip_score", 22),
})
```

* One of `UserId`, or `Email` is required.
* `Name` is optional.
* `RemoteCreatedAt` is optional, must be an integer representing seconds since Unix Epoch.
* `CustomData` is optional, and can be constructed using the helper as above, or as a passed `map[string]interface{}`.

#### Find

```go
usr, err := c.Users.Find(&intercom.UserParams{
  Email: "jamie+jamie@intercom.io",
})
```

* One of `Id`, `UserId`, or `Email` is required.

#### List

```go
usr_list, err := c.Users.List()
```

### Errors

All operations can return errors in addition to structured data. These errors should be checked.