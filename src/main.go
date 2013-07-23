// Copyright 2011 The goauth2 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This program makes a call to the specified API, authenticated with OAuth2.
// a list of example APIs can be found at https://code.google.com/oauthplayground/
package main

import (
	"anaconda"
	
	"fmt"
	"net/url"
)


func main() {
	anaconda.SetConsumerKey("")
	anaconda.SetConsumerSecret("")
	api := anaconda.NewTwitterApi("", "")
	
	v := url.Values{}
	v.Set("count", "30")
	
	searchResult, _ := api.GetDirectMessagesShow(v);
	
	
	fmt.Print(searchResult);
}   
