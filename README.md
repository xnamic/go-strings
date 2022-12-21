# Strings
Evaluate string as math and return the math operation

# Supported operators
- `*`
- `/`
- `+`
- `-`
- `^`

# Instalation
go get github.com/xnamic/strings

# Usage and Example

```
var res float64
var err error

// return -150
res, err = strings.Eval("10 (10 - 25)")

// return -0.4
res, err = strings.Eval("-10 / 25")

// return -407
res, err = strings.Eval("1+(1+3) * ((4 - 100) - (3+3))")


```


# Licence
The MIT License (MIT) - see LICENCE for more detail



