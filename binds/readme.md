# Binds

stuct bindings

```go

/* THIS IS SAMPLE DATA */

type Bindable struct {
	Name	string	`binds:"name"`
	Age 	uint	`binds:"age"`
	IsChild	bool	`binds:"child"`

	Ignored  string
	Ignored2 string
}


b := Bindable{Name: "joe", Age: 30, IsChild: false }

```

## Serialize

create the data map

```go

// Serialize(instance interface{}) map[string]interface{}
/* *** Expecting result: 
v := map[string]interface{}{
	"name": "joe",
	"age": uint(30),
	"child": false,
}
*** */

v := Serialize(b)

```

## Populate

create the instance, based on tmplInstance struct type, with map data.

```go

// Populate(data map[string]interface{}, tmplInstance interface{}) interface{}

d := Populate(v, b)
// v === d here

```

