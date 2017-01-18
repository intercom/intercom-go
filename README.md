# Intercom-Go

[![Build Status](https://travis-ci.org/intercom/intercom-go.svg)](https://travis-ci.org/intercom/intercom-go)

Thin client for the [Intercom](https://www.intercom.io) API.

_Currently in beta, though breaking API changes are not expected._

## Install

`go get gopkg.in/intercom/intercom-go.v1`

## Usage

### Getting a Client

The first step to using Intercom's Go client is to create a client object, using your App ID and Api Key from your [settings](http://app.intercom.io/apps/api_keys).

```go
import (
	`gopkg.in/intercom/intercom-go.v1`
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
		{ID: "5"},
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

### Contacts

#### Find

```go
contact, err := ic.Contacts.FindByID("46adad3f09126dca")
```

```go
contact, err := ic.Contacts.FindByUserID("27")
```

#### List

```go
contactList, err := ic.Contacts.List(intercom.PageParams{Page: 2})
contactList.Pages // page information
contactList.Contacts // []Contact
```

```go
contactList, err := ic.Contacts.ListByEmail("test@example.com", intercom.PageParams{})
```

#### Create

```go
contact := intercom.Contact{
	Email: "test@example.com",
	Name: "SomeContact",
	CustomAttributes: map[string]interface{}{"is_cool": true},
}
savedContact, err := ic.Contacts.Create(&contact)
```

* No identifier is required.
* Set values for UserID will be ignored (consider creating _Users_ instead)

#### Update

```go
contact := intercom.Contact{
	UserID: "abc-13d-3",
	Name: "SomeContact",
	CustomAttributes: map[string]interface{}{"is_cool": true},
}
savedContact, err := ic.Contacts.Update(&contact)
```

* ID or UserID is required.
* Will not create new contacts.

#### Convert

Used to convert a Contact into a User

```go
contact := intercom.Contact{
	UserID: "abc-13d-3",
}
user := intercom.User{
	Email: "myuser@signedup.com",
}
savedUser, err := ic.Contacts.Convert(&contact, &user)
```

* If the User does not already exist in Intercom, the Contact will be uplifted to a User.
* If the User does exist, the Contact will be merged into it and the User returned.

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
	UserID: "27",
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
taggingList := intercom.TaggingList{Name: "GoTag", Users: []intercom.Tagging{{UserID: "27"}}}
savedTag, err := ic.Tags.Tag(&taggingList)
```

A `Tagging` can identify a User or Company, and can be set to `Untag`:

```go
taggingList := intercom.TaggingList{Name: "GoTag", Users: []intercom.Tagging{{UserID: "27", Untag: intercom.Bool(true)}}}
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

### Messages

#### New Admin to User/Contact Email

```go
msg := intercom.NewEmailMessage(intercom.PERSONAL_TEMPLATE, intercom.Admin{ID: "1234"}, intercom.User{Email: "test@example.com"}, "subject", "body")
savedMessage, err := ic.Messages.Save(&msg)
```

Can use intercom.PLAIN_TEMPLATE too, or replace the intercom.User with an intercom.Contact.

#### New Admin to User/Contact InApp

```go
msg := intercom.NewInAppMessage(intercom.Admin{ID: "1234"}, intercom.Contact{Email: "test@example.com"}, "body")
savedMessage, err := ic.Messages.Save(&msg)
```

#### New User Message

```go
msg := intercom.NewUserMessage(intercom.User{Email: "test@example.com"}, "body")
savedMessage, err := ic.Messages.Save(&msg)
```

### Conversations

### Find Conversation

```go
convo, err := intercom.Conversations.Find("1234")
```

### List Conversations

#### All

```go
convoList, err := intercom.Conversations.ListAll(intercom.PageParams{})
```

#### By User

Showing all for user:

```go
convoList, err := intercom.Conversations.ListByUser(&user, intercom.SHOW_ALL, intercom.PageParams{})
```

Showing just Unread for user:

```go
convoList, err := intercom.Conversations.ListByUser(&user, intercom.SHOW_UNREAD, intercom.PageParams{})
```

#### By Admin

Showing all for admin:

```go
convoList, err := intercom.Conversations.ListByAdmin(&admin, intercom.SHOW_ALL, intercom.PageParams{})
```

Showing just Open for admin:

```go
convoList, err := intercom.Conversations.ListByAdmin(&admin, intercom.SHOW_OPEN, intercom.PageParams{})
```

Showing just Closed for admin:

```go
convoList, err := intercom.Conversations.ListByAdmin(&admin, intercom.SHOW_CLOSED, intercom.PageParams{})
```

### Reply

User reply:

```go
convo, err := intercom.Conversations.Reply("1234", &user, intercom.CONVERSATION_COMMENT, "my message")
```
User reply with attachment:

```go
convo, err := intercom.Conversations.ReplyWithAttachmentURLs("1234", &user, intercom.CONVERSATION_COMMENT, "my message", string[]{"http://www.example.com/attachment.jpg"})
```

User reply that opens:

```go
convo, err := intercom.Conversations.Reply("1234", &user, intercom.CONVERSATION_OPEN, "my message")
```

Admin reply:

```go
convo, err := intercom.Conversations.Reply("1234", &admin, intercom.CONVERSATION_COMMENT, "my message")
```

Admin note:

```go
convo, err := intercom.Conversations.Reply("1234", &admin, intercom.CONVERSATION_NOTE, "my message to just admins")
```

### Open and Close

Open:

```go
convo, err := intercom.Conversations.Open("1234", &openerAdmin)
```

Close:

```go
convo, err := intercom.Conversations.Close("1234", &closerAdmin)
```

### Assign

```go
convo, err := intercom.Conversations.Assign("1234", &assignerAdmin, &assigneeAdmin)
```

### Bulk

Bulk operations are supported through this package, see [the documentation](https://doc.intercom.io/api/#bulk-apis) for details.

New user bulk job, posts a user and deletes another:

```go
jobResponse := ic.Jobs.NewUserJob(intercom.NewUserJobItem(&user, intercom.JOB_POST), intercom.NewUserJobItem(&userTwo, intercom.JOB_DELETE))
```

Append to an existing user job:

```go
jobResponse := ic.Jobs.AppendUsers("job_5ca1ab1eca11ab1e", intercom.NewUserJobItem(&user, intercom.JOB_POST))
```

New event bulk job:

```go
jobResponse := ic.Jobs.NewEventJob(intercom.NewEventJobItem(&event), intercom.NewEventJobItem(&eventTwo))
```

Append to an existing event job:

```go
jobResponse := ic.Jobs.AppendEvents("job_5ca1ab1eca11ab1e", intercom.NewEventJobItem(&eventTwo))
```

Find a Job:

```go
jobResponse := ic.Jobs.Find("job_5ca1ab1eca11ab1e")
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
	Patch(string, interface{}) ([]byte, error)
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
