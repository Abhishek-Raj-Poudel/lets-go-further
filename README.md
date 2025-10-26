# Managing Bad Requests

By default `json.decode()` gives good error output but that may not be enough especially if it's api used by the public. So we will be Triaging the Decode error, json.decode sends 5 types of error , we will be accounting for that in readJSON() function in helper.go
