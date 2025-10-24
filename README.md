# A Basic HTTP Server


API versioning
APIs which support real-world businesses and users often need to change their functionality
and endpoints over time — sometimes in a backwards-incompatible way. So, to avoid
problems and confusion for clients, it’s a good idea to always implement some form of API
versioning.
There are two common approaches to doing this:
1. By prefixing all URLs with your API version, like /v1/healthcheck or /v2/healthcheck.
2. By using custom Accept and Content-Type headers on requests and responses to
convey the API version, like Accept: application/vnd.greenlight-v1 .
From a HTTP semantics point of view, using headers to convey the API version is the ‘purer’
approach. But from a user-experience point of view, using a URL prefix is arguably better. It
makes it possible for developers to see which version of the API is being used at a glance,
and it also means that the API can still be explored using a regular web browser (which is
harder if custom headers are required).
Throughout this book we’ll version our API by prefixing all the URL paths with /v1/ — just
like we did with the /v1/healthcheck endpoint in this chapter.
