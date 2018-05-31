// Copyright 2015 The intercom-go AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package intercom-go provides a thin client for the Intercom API: http://developers.intercom.com/reference.

The first step to using Intercom's Go client is to create a client object, using your App ID and Api Key from your [settings](http://app.intercom.io/apps/api_keys).

  import (
    "gopkg.in/intercom/intercom-go.v2"
  )
  ic := intercom.NewClient("appID", "apiKey")

The client can be configured with different options by calls to Option:

  ic.Option(intercom.TraceHTTP(true)) // turn http tracing on
  ic.Option(intercom.BaseURI("http://intercom.dev")) // change the base uri used, useful for testing
  ic.Option(intercom.SetHTTPClient(myHTTPClient)) // set a new HTTP client

Errors

Errors may be returned from some calls. Errors returned from the API will implement `intercom.IntercomError` and can be checked:

  _, err := ic.Users.FindByEmail("doesnotexist@intercom.io")
  if herr, ok := err.(intercom.IntercomError); ok && herr.GetCode() == "not_found" {
    fmt.Print(herr)
  }

HTTP Client

The HTTP Client used by this package can be swapped out for one of your choosing, with your own configuration, it just needs to implement the HTTPClient interface:

  type HTTPClient interface {
    Get(string, interface{}) ([]byte, error)
    Post(string, interface{}) ([]byte, error)
    Delete(string, interface{}) ([]byte, error)
  }

The client will probably need to work with `appId`, `apiKey` and `baseURI` values. See the provided client for an example. Then create an Intercom Client and inject the HTTPClient:

  ic := intercom.Client{}
  ic.Option(intercom.SetHTTPClient(myHTTPClient))
  // ready to go!

On Bools

Due to the way Go represents the zero value for a bool, it's necessary to pass pointers to bool instead in some places. The helper `intercom.Bool(true)` creates these for you.

Pagination

For many resources, pagination should be applied through the use of a PageParams object passed into List() functions.


  pageParams := PageParams{
    Page: 2,
    PerPage: 10,
  }
  ic.Users.List(pageParams)

*/
package intercom
