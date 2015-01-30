# Intercom-Go

Thin client for the [Intercom](https://www.intercom.io) API.

## Install

`go get github.com/intercom/intercom-go`

## Usage

### Getting a Client

The first step to using Intercom's Go client is to create a client object, using your App ID and Api Key from your [settings](http://app.intercom.io/apps/api_keys).

```go
import (
  "github.com/intercom/intercom-go"
  "github.com/intercom/intercom-go/client"
)

ic := intercom.NewClient("appID", "apiKey")
```

This client can then be used to make requests.

#### Client Options

The client can be configured with different options by calls to `ic.Option`:

```go
ic.Option(intercom.TraceHTTP(true)) // turn http tracing on
ic.Option(intercom.BaseURI("http://intercom.dev")) // change the base uri used, useful for testing
```

or combined:

```go
ic.Option(intercom.TraceHTTP(true), intercom.BaseURI("http://intercom.dev"))
```

### Users

#### Save

```go
user := c.User
user.UserID = "27"
user.Email = "test@example.com"
user.Name = "InterGopher"
user.SignedUpAt = 1416750500
user.CustomAttributes = map[string]interface{}{"is_cool": true}
user, err := user.Save()
```

* One of `UserID`, or `Email` is required.
* `SignedUpAt` (optional), like all dates in the client, must be an integer(32) representing seconds since Unix Epoch.

#### Find

```go
user, err := c.User.FindByID("46adad3f09126dca")
```

```go 
user, err := c.User.FindByUserID("27")
```

```go
user, err := c.User.FindByEmail("test@example.com")
```

#### List

```go
user_list, err := c.Users.List(client.PageParams{Page: 2})
user_list.Pages // page information
user_list.Users // []User
```


### Events

#### Save
  
```go
event := ic.Event
event.UserId = "27"
event.EventName = "bought_item"
event.CreatedAt = int32(time.Now().Unix())
event.Metadata = map[string]interface{}{"item_name": "PocketWatch"}
err := event.Save()
```

* One of `UserID`, or `Email` is required.
* `EventName` is required.
* `CreatedAt` is optional, must be an integer representing seconds since Unix Epoch. Will be set to _now_ unless given.
* `Metadata` is optional, and can be constructed using the helper as above, or as a passed `map[string]interface{}`.
  

### Errors

Errors may be returned from some calls. Errors returned from the API will implement `client.IntercomError` and can be checked:

```go
_, err := c.User.FindByEmail("doesnotexist@intercom.io")
if herr, ok := err.(client.IntercomError); ok && herr.GetCode() == "not_found" {
  fmt.Print(herr)
}
```

----

#### Old Stuff

----

### Notes

#### New

```go
note, err := c.Notes.New(intercom.NoteParams{
  UserId: "27",
  Body: "Unicorn Developer",
  AdminId: "1457"
})
```

* One of `Id`, `UserId`, or `Email` is required.
* `Body` is required.
* `AdminId` is optional.

#### Find

```go
note, _ := c.Notes.Find(intercom.NoteParams{
  Id: "87",
})

log.Printf(note)
// [intercom] note { id: 87, body: "my note" }
```
  
  * `Id` is required.



### Admins


#### List

```go
admin_list, err := c.Admins.List()
admin_list.Admins // list of admins
```

### Conversations


#### List

```go
conversation_list, err := c.Conversations.List(intercom.PageParams{}) // no paging; therefore first page
conversation_list.Pages // page information
conversation_list.Conversations // conversations
```

or with paging...

```go
user_list, err := c.Users.List(intercom.PageParams{ Page: 2})
```

#### List By Admin

```go
conversation_list, err := c.Conversations.ListForAdmin(intercom.PageParams{
  Page: 2,
}, intercom.AdminParams{
  Id:   "2793",
  Open: intercom.Bool(true),
})
```

* Id is required for Admin
* Open is optional, passing `intercom.Bool(true)` will return only open conversations (or false == closed conversations)

#### List By User

```go
conversation_list, err := c.Conversations.ListForUser(intercom.PageParams{
  Page: 4,
}, intercom.UserParams{
  Id:     "54713d0c8a68188189000013",
  Unread: intercom.Bool(true),
})
```

* `Id`, `UserId`, or `Email` is required.
* Unread is optional, passing `intercom.Bool(true)` will return only unread conversations for that User.

#### Find

```go
convo, err := c.Conversations.Find(intercom.ConversationParams{
  Id:     "1315",
})
```

* `Id` is required.

### On Bools

Due to the way Go represents the zero value for a bool, it's necessary to pass pointers to bool instead in some places.

The helper `intercom.Bool(true)` creates these for you.

### Errors

All operations can return errors in addition to structured data. These errors should be checked.
