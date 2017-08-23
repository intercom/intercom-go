# Intercom-Go

[![Build Status](https://travis-ci.org/intercom/intercom-go.svg)](https://travis-ci.org/intercom/intercom-go)

Thin client for the [Intercom](https://www.intercom.io) API.

## Install

`go get gopkg.in/intercom/intercom-go.v2`

[![docker_image 1](https://cloud.githubusercontent.com/assets/15954251/17524401/5743439e-5e56-11e6-8567-d3d9da1727da.png)](https://hub.docker.com/r/cathalhoran/intercom-go/) <br>
Try out our [Docker Image (Beta)](https://hub.docker.com/r/cathalhoran/intercom-go/) to help you get started more quickly. <br>
It should make it easier to get setup with the SDK and start interacting with the API. <br>
(Note, this is in Beta and is for testing purposes only, it should not be used in production)

## Usage

### Getting a Client

```go
import (
	intercom "gopkg.in/intercom/intercom-go.v2"
)
// You can use either an an OAuth or Access Token
ic := intercom.NewClient("access_token", "")
```
This client can then be used to make requests.

If you already have an access token you can find it [here](https://app.intercom.com/developers/_). If you want to create or learn more about access tokens then you can find more info [here](https://developers.intercom.io/docs/personal-access-tokens).

If you are building a third party application you can get your OAuth token by [setting-up-oauth](https://developers.intercom.io/page/setting-up-oauth) for Intercom.
You can use the [Goth library](https://github.com/markbates/goth) which is a simple OAuth package for Go web aplicaitons and supports Intercom to more easily implement Oauth.

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
	SignedUpAt: int64(time.Now().Unix()),
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
userList, err := ic.Users.Scroll("")
scrollParam := userList.ScrollParam
userList, err := ic.Users.Scroll(scrollParam)
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
contactList, err := ic.Contacts.Scroll("")
scrollParam = contactList.ScrollParam
contactList, err := ic.Contacts.Scroll(scrollParam)
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
	CreatedAt: int64(time.Now().Unix()),
	Metadata: map[string]interface{}{"item_name": "PocketWatch"},
}
err := ic.Events.Save(&event)
```

* One of `UserID`, `ID`, or `Email` is required (With leads you need to use ID).
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
segments, err := segmentList.Segments
```

#### Find

```go
segment, err := ic.Segments.Find("abc312daf2397")
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


### Pull Requests

- **Add tests!** Your patch won't be accepted if it doesn't have tests.

- **Document any change in behaviour**. Make sure the README and any other
  relevant documentation are kept up-to-date.

- **Create topic branches**. Don't ask us to pull from your master branch.

- **One pull request per feature**. If you want to do more than one thing, send
  multiple pull requests.

- **Send coherent history**. Make sure each individual commit in your pull
  request is meaningful. If you had to make multiple intermediate commits while
  developing, please squash them before sending them to us.
