package models

import "testing"

func TestUserGet(t *testing.T) {
	u := &User{
		OpenId: "aaa",
	}
	if u.Get("openid") != "aaa" {
		t.Fatalf("expect openid \"aaa\", but get %v", u.Get("openid"))
	}
}
