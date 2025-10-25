# Encoding Structs

- Here you will see how to use structs as json in short
- How to change json key (because by default json keys will look exactly how structs key looks) by useing *struct tags* in data/movies.go
eg
```go
type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Title     string    `json:"title"`
	Year      int32     `json:"year"`
	Runtime   int32     `json:"runtime"`
	Genres    []string  `json:"genres"`
	Version   int32     `json:"version"`
}
```
- How to hide structs field in json object by using the omitempty and - struct tag directives.

The - (hyphen) directive can be used when you never want a particular struct field to
appear in the JSON output. This is useful for fields that contain internal system information
that isn’t relevant to your users, or sensitive information that you don’t want to expose (like
the hash of a password).

In contrast the omitempty directive hides a field in the JSON output if and only if the struct
field value is empty, where empty is defined as being:
- Equal to false, 0, or ""
- An empty array, slice or map
- A nil pointer or a nil interface value

- The string tag to return an (let's say) int value to string in json response eg
```go

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"` 
    Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Runtime   int32     `json:"runtime,omitempty,string"` //Add the string directive
	Genres    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version"`
}

Now the resulting json response the runtime will appear a string instead of a number
```
