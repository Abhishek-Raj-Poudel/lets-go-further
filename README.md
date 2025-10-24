# Fixed-Format JSON
In this section of the book you’ll learn:
- How to send JSON responses from your REST API (including error responses).
- How to encode native Go objects into JSON using the *encoding/json* package.
- Different techniques for customizing how Go objects are encoded to JSON — first by
- using struct tags, and then by leveraging the json.Marshaler interface.
- How to create a reusable helper for sending JSON responses, which will ensure that all
your API responses have a sensible and consistent structure.
