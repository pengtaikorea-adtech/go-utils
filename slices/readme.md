# go-util/slices

slice map/reduce lambda functions

## install: go get 

```bash
$ go get "github.com/pengtaikorea-adtech/go-utils/slices"
```

## Types/Constants

- **Any** represents "any" element type. 
  ```go
  type Any interface{}
  ```
- **ASlice** represents "any|anonymous|a-|a " slices.
  ```go
  type ASlice interface{}
  ```
- **ElemType...** Slice Element Type const. to be used for "Map". 
type reflect.Type

  - ElemTypeInt
  - ElemTypeBool
  - ElemTypeString
  

## Snippet

refered by javascript map, forEach, filter, reduce

### Each

runs forEach

```go
import "github.com/pengtaikorea-adtech/go-utils/slices"

var targetSlice []interface{}

var strSlice, err := slices.Each(
	func(e slices.Any, i int, s slices.ASlice) error {
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
	func(e slices.Any, i int, s slices.ASlice) (slices.Any, error) {
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

filter slice.

```go
import "github.com/pengtaikorea-adtech/go-utils/slices"

filtered, err := slices.Filter(
	func(e slices.Any, i int, s slices.ASlice) bool {
		// filter function true/false here
		return true
	},
targetSlice)
```

### Reduce
