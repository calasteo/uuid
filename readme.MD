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
user.Name = "Karisma Prabakuncara"

fmt.Printf("%+v\n", user)

// &{ID:1709b84a-e322-ba08-4acb-e018eb39efa9 Name:Karisma Prabakuncara}

```
