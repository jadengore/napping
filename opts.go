// Copyright (c) 2012-2013 Jason McVetta.  This is Free Software, released
// under the terms of the GPL v3.  See http://www.gnu.org/copyleft/gpl.html for
// details.  Resist intellectual serfdom - the ownership of ideas is akin to
// slavery.

package napping

import (
	"net/http"
	"net/url"
)

type Opts struct {
	Userinfo *url.Userinfo // Optional username/password to authenticate this request
	Header   *http.Header  // HTTP Headers to use (will override defaults)
	// HTTP status code we expect the server to return on a successful request.
	// If ExpectedStatus is non-zero and server returns a different code, Send()
	// will return a BadStatus error.
	ExpectedStatus int
}

// update merges this Opts instance with another instance, preferring the other
// instance where there is disagreement.
func (this *Opts) update(other *Opts) *Opts {
	merged := this
	if other.Userinfo != nil {
		merged.Userinfo = other.Userinfo
	}
	h := *merged.Header
	for k, v := range *other.Header {
		h[k] = v
	}
	merged.Header = &h
	if other.ExpectedStatus != 0 {
		merged.ExpectedStatus = other.ExpectedStatus
	}
	return merged
}
