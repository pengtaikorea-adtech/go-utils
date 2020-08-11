# go-util/slices

slice map/reduce lambda functions

## install: go get 

```bash
$ go get "github.com/pengtaikorea-adtech/go-utils/slices"
```

## Types/Constants

- **ElementType...** Slice Element Type const. to be used for "Map". 
type reflect.Type

  - ElementTypeInt
  - ElementTypeBool
  - ElementTypeString
  

## Snippet

refered by javascript map, forEach, filter, reduce

### Each

runs forEach

```go
import "github.com/pengtaikorea-adtech/go-utils/slices"

var targetSlice []interface{}

var strSlice, err := slices.Each(
	func(e interface{}, i int, s interface{}) error {
		// doSomethingOn(e)

		// return nil error
		return nil
	}, 
targetSlice) // with target slice
```


### Map 

map slice to another.

```go
import (
	"errors"

	"github.com/pengtaikorea-adtech/go-utils/slices"
)

err := slices.Map(
	func(e interface{}, i int, s interface{}) (interface{}, error) {
		// build map value
		if s, ok := e.(string); ok {
			return s, nil
		} else {
			return nil, errors.New("type assertion fail")
		}
	}, 
// return slice element type required
targetSlice, slices.ElemTypeString) 
```

### Filter

filter slice

```go
import "github.com/pengtaikorea-adtech/go-utils/slices"

filtered, err := slices.Filter(
	func(e interface{}, i int, s interface{}) bool {
		// filter function true/false here
		return true
	},
targetSlice)
```

### Reduce

reduce elements

```go
import "github.com/pengtaikorea-adtech/go-utils/slices"

reduced, err := slices.Reduce(
	func(t interface{}, e interface{}, i int, s interface{}) (interface{}, error) {
		// Do somethings with t
		return t, nil
	},
targetSlice, initValue)
)
