# go-utils

go routine utilities to increase developing performance.

 - [binds](binds)
	- Populate( map[string]interface{}, *struct ) -> error
	- Mapping( *struct ) -> map[string]interface{}
 - [slices](slices)
	- SyncEach( slice, func(element, index, slice) -> error ) -> slice, error
	- AsyncEach( slice, func(element, index, slice) -> error, num_workers ) -> slice, error
	- Filter( slice, func(element, index, slice) -> bool ) -> slice, error

## binds
binding go struct with `map[string]interface{}` type data.

```go
import "github.com/pengtaikorea-adtech/binds"

// SampleStruct struct to bind.
// auto-mounting values by field tags
type SampleStruct struct {
	// if required value not specified, 
	ID string `binds:"id" required`
	Name string `binds:"name"`
	Ignored bool
	Composite map[string]interface{} `binds:"info"`
}

var sampleData binds.Map // equivalent to map[string]interface{}
var sampleInstance SampleStruct
// 
// ... load data
// sampleData = yaml.Unmarshal or json.Unmarshal ...
//
err := binds.Populate(sampleData, &sampleInstance)
if err != nil {
	// do exception handling here
}
```

## slices
refer javascript .forEach, .map, .reduce, and .filter functionality.

```go

```

## 
