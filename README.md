# API Endpoints and RESTful Routing
When you’re building an API with endpoints like this in Go, one of the first hurdles you’ll
meet is the fact that http.ServeMux — the router in the Go standard library — is quite
limited in terms of its functionality. In particular it doesn’t allow you to route requests to
different handlers based on the request method (GET, POST, etc.), nor does it provide
support for clean URLs with interpolated parameters.
Although you can work-around these limitations, or implement your own router, generally
it’s easier to use one of the many third-party routers that are available instead.
In this chapter we’re going to integrate the popular httprouter package with our
application. Most importantly, httprouter is stable, well-tested and provides the
functionality we need — and as a bonus it’s also extremely fast thanks to its use of a radix
tree for URL matching. If you’re building a REST API for public consumption, then
httprouter is a solid choice.

install it like 
```sh
go get github.com/julienschmidt/httprouter
```
