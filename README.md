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
savedUser, err := ic.Users.Save(&user)
```

* One of `UserID`, or `Email` is required.
* `SignedUpAt` (optional), like all dates in the client, must be an integer(32) representing seconds since Unix Epoch.

##### Adding/Removing Companies

Adding a Company:

```go
companyList := intercom.CompanyList{
  Companies: []intercom.Company{
    intercom.Company{ID: "5"},
  },
}
user := intercom.User{
  UserID: "27",
  Companies: &companyList,
}
```

Removing is similar, but adding a `Remove: intercom.Bool(true)` attribute to a company.

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
userList, err := ic.Users.List(intercom.PageParams{Page: 2})
userList.Pages // page information
userList.Users // []User
```

```go
userList, err := ic.Users.ListBySegment("segmentID123", intercom.PageParams{})
```

```go
userList, err := ic.Users.ListByTag("42", intercom.PageParams{})
```

#### Delete

```go
user, err := ic.Users.Delete("46adad3f09126dca")
```

### Companies

#### Save

```go
company := intercom.Company{
  CompanyID: "27",
  Name: "My Co",
  CustomAttributes: map[string]interface{}{"is_cool": true},
  Plan: &intercom.Plan{Name: "MyPlan"},
}
savedCompany, err := ic.Companies.Save(&company)
```

* `CompanyID` is required.

#### Find

```go
company, err := ic.Companies.FindByID("46adad3f09126dca")
```

```go 
company, err := ic.Companies.FindByCompanyID("27")
```

```go
company, err := ic.Companies.FindByName("My Co")
```

#### List

```go
companyList, err := ic.Companies.List(intercom.PageParams{Page: 2})
companyList.Pages // page information
companyList.Companies // []Companies
```

```go
companyList, err := ic.Companies.ListBySegment("segmentID123", intercom.PageParams{})
```

```go
companyList, err := ic.Companies.ListByTag("42", intercom.PageParams{})
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


### Admins

#### List

```go
adminList, err := ic.Admins.List()
admins := adminList.Admins
```

### Tags

#### List

```go
tagList, err := ic.Tags.List()
tags := tagList.Tags
```

#### Save

```go
tag := intercom.Tag{Name: "GoTag"}
savedTag, err := ic.Tags.Save(&tag)
```

`Name` is required. Passing an `ID` will attempt to update the tag with that ID.

#### Delete

```go
err := ic.Tags.Delete("6")
```

#### Tagging Users/Companies

```go
taggingList := intercom.TaggingList{Name: "GoTag", Users: []intercom.Tagging{intercom.Tagging{UserID: "27"}}}
savedTag, err := ic.Tags.Tag(&taggingList)
```

A `Tagging` can identify a User or Company, and can be set to `Untag`:

```go
taggingList := intercom.TaggingList{Name: "GoTag", Users: []intercom.Tagging{intercom.Tagging{UserID: "27", Untag: intercom.Bool(true)}}}
savedTag, err := ic.Tags.Tag(&taggingList)
```

### Segments

#### List

```go
segmentList := ic.Segments.List()
segments := segmentList.Segments
```

#### Find

```go
segment := ic.Segments.Find("abc312daf2397")
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
  Delete(string, interface{}) ([]byte, error)
}
```

It'll probably need to work with `appId`, `apiKey` and `baseURI` values. See the provided client for an example. Then create an Intercom Client and inject the HTTPClient:

```go
ic := intercom.Client{}
ic.Option(intercom.SetHTTPClient(myHTTPClient))
// ready to go!
```

### On Bools

Due to the way Go represents the zero value for a bool, it's necessary to pass pointers to bool instead in some places.

The helper `intercom.Bool(true)` creates these for you.
