# JSON Encoding

In this we look at how to encode native Go objects
(like maps, structs and slices) to JSON.

we will use encode/json package to do that and use json json.Marshal() function.
The function's signature looks like this

```go
func Marshal(v interface{}) ([]byte, error)
```

> Note : The v parameter in the above method has the type interface{} (known as the
empty interface). This effectively means that we’re able to pass any Go type to
Marshal() as the v parameter.
