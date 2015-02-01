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
)

ic := intercom.NewClient("appID", "apiKey")
```

This client can then be used to make requests.

#### Client Options

The client can be configured with different options by calls to `ic.Option`:

```go
ic.Option(intercom.TraceHTTP(true)) // turn http tracing on
ic.Option(intercom.BaseURI("http://intercom.dev")) // change the base uri used, useful for testing
ic.Option(intercom.SetHTTPClient(myHTTPClient)) // set a new HTTP client, see below for more info
```

or combined:

```go
ic.Option(intercom.TraceHTTP(true), intercom.BaseURI("http://intercom.dev"))
```

### Users

#### Save

```go
user := intercom.User{
  UserID: "27",
  Email: "test@example.com",
  Name: "InterGopher",
  SignedUpAt: int32(time.Now().Unix()),
  CustomAttributes: map[string]interface{}{"is_cool": true},
}
saved_user, err := ic.Users.Save(&user)
```

* One of `UserID`, or `Email` is required.
* `SignedUpAt` (optional), like all dates in the client, must be an integer(32) representing seconds since Unix Epoch.

#### Find

```go
user, err := ic.Users.FindByID("46adad3f09126dca")
```

```go 
user, err := ic.Users.FindByUserID("27")
```

```go
user, err := ic.Users.FindByEmail("test@example.com")
```

#### List

```go
user_list, err := ic.Users.List(client.PageParams{Page: 2})
user_list.Pages // page information
user_list.Users // []User
```


### Events

#### Save
  
```go
event := intercom.Event{
  UserId: "27",
  EventName: "bought_item",
  CreatedAt: int32(time.Now().Unix()),
  Metadata: map[string]interface{}{"item_name": "PocketWatch"},
}
err := ic.Events.Save(&event)
```

* One of `UserID`, or `Email` is required.
* `EventName` is required.
* `CreatedAt` is optional, must be an integer representing seconds since Unix Epoch. Will be set to _now_ unless given.
* `Metadata` is optional, and can be constructed using the helper as above, or as a passed `map[string]interface{}`.

### Conversations

#### Find

```go
conversation, err := ic.Conversations.Find("12")
```

### Errors

Errors may be returned from some calls. Errors returned from the API will implement `intercom.IntercomError` and can be checked:

```go
_, err := ic.Users.FindByEmail("doesnotexist@intercom.io")
if herr, ok := err.(intercom.IntercomError); ok && herr.GetCode() == "not_found" {
  fmt.Print(herr)
}
```

### HTTP Client

The HTTP Client used by this package can be swapped out for one of your choosing, with your own configuration, it just needs to implement the HTTPClient interface:

```go
type HTTPClient interface {
  Get(string, interface{}) ([]byte, error)
  Post(string, interface{}) ([]byte, error)
}
```

It'll probably need to hold `appId`, `apiKey` and `baseURI` values. See the provided client for an example. Then create an Intercom Client and inject the HTTPClient:

```go
ic := intercom.Client{}
ic.Option(intercom.SetHTTPClient(myHTTPClient))
// ready to go!
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
