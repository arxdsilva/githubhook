webhook
===============================================

Golang parser for [webhooks][webhook]. Not a server, though it could
be integrated with one.

[![Build Status](https://travis-ci.org/arxdsilva/webhook.svg?branch=master)](https://travis-ci.org/arxdsilva/webhook)
[![Go Report Card](https://goreportcard.com/badge/github.com/arxdsilva/webhook)](https://goreportcard.com/report/github.com/arxdsilva/githubhook)
[![GoDoc](https://godoc.org/github.com/arxdsilva/githubhook?status.svg)](https://godoc.org/github.com/arxdsilva/githubhook)

Installation
-----------------------------------------------

```ShellSession
$ go get gopkg.in/arxdsilva/webhook.v0
```

Usage
-----------------------------------------------

Given an incoming `*http.Request` representing a webhook signed with a `secret`,
use `webhook` to validate and parse its content:

```go
secret := []byte("don't tell!")
hook, err := webhook.Parse(secret, req)
```

Plays nicely with the [google/go-github][gh-go-github] client!

```go
evt := github.PullRequestEvent{}
if err := json.Unmarshal(hook.Payload, &evt); err != nil {
  fmt.Println("Invalid JSON?", err)
}
```

[webhook]: https://developer.github.com/webhooks/
[gh-go-github]: https://github.com/google/go-github
