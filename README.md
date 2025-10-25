# Formatting and Enveloping Responses

The response is not readable(somewhat) in curl so instead of using `json.Marshal` so we can ue `json.MarshalIndent()`

- it's lets performative and takes more ram (although in the grand scheme of things is abizmel) so we will be sticking with `json.Marshal`. And also i will just use postman instead of curl

> Note: Behind the scenes json.MarshalIndent() works by calling json.Marshal() as normal, then running the JSON through the standalone json.Indent() function to add the whitespace. There’s also a reverse function available, json.Compact(), which you can use to remove whitespace from JSON.



