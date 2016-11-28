# go-foreman
Go library for accessing the Foreman API. This Foreman https://en.wikipedia.org/wiki/Foreman_(software).

The general design and implementation (and code as well) of this comes from the [go-github](https://github.com/google/go-github) library from Google.

## Status
There's still a LOT to do. I have a need for this at my current job and I'm sharing it here in hopes that someone else has the same need.

## Integration testing
Integration tests for a library like this can be hard because the only real way to do integration tests is to do it against a Foreman instance. I'd like
to be able to use [this](https://github.com/ory-am/dockertest/tree/v3) to spin up a Foreman Docker container for ease of use.

## Can I help?
Of course! I love Pull Requests! I've linked above for how the general design of the library should work and [here](https://theforeman.org/api/1.13/index.html)
are the API docs for Foreman. Just find something you want to implement and do it!
