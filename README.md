# Algorithms

Collection of algorithm implementations for Go.

## Error detection

[Damm](https://github.com/petoc/algo/tree/master/damm)<br>
[Damm64](https://github.com/petoc/algo/tree/master/damm64) (using quasigroups for character base from 3 to 64 characters)<br>
[Luhn](https://github.com/petoc/algo/tree/master/luhn)<br>


## Usage

### Damm

Implementation of error detection algorithm for numeric codes from H. Michael Damm.

```go
import "github.com/petoc/algo/damm"
```

```go
damm.Calculate("123456789") // 4
damm.Validate("1234567894") // true
```

### Damm64

Based on error detection algorithm from H. Michael Damm. Uses pre-generated quasigroups for character base from 3 to 64 characters.

```go
import "github.com/petoc/algo/damm64"
```

```go
base := "0123456789ABCDEFGHIJKLMNOPQRSTUV"
damm64.Calculate(base, "G12Q") // F
damm64.Validate(base, "G12QF") // true
```
### Luhn

Implementation of error detection algorithm for numeric codes from Hans Peter Luhn.

```go
import "github.com/petoc/algo/luhn"
```

```go
luhn.Calculate("123456789") // 7
luhn.Validate("1234567897") // true
```

## Sources

Damm Quasigroups (http://www.md-software.de/math/DAMM_Quasigruppen.txt)

## License

Licensed under MIT license.
