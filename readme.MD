## Ordered UUID 

<i>inspired by Str::orderedUUID() on laravel</i>

#### Installation

```go
go get -u github.com/calasteo/uuid
```

#### usage
```go
type (
    User struct {
	   ID   string `json:"id"`
	   Name string `json:"name"`
    }   
)

user:= &User{}
user.ID = uuid.GenerateOrderedUUID()
user.Name = "John Doe"

fmt.Printf("%+v\n", user)

```
