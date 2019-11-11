# Backoff

Backoff provides backoff policies for use in rate limiting.

## Usage

A backoff policy can be created using the [NewConstant](https://godoc.org/github.com/tradyfinance/backoff#NewConstant), [NewLinear](https://godoc.org/github.com/tradyfinance/backoff#NewLinear), or [NewExponential](https://godoc.org/github.com/tradyfinance/backoff#NewExponential) functions. Jitter can be applied using the [NewJitter](https://godoc.org/github.com/tradyfinance/backoff#NewJitter) function. The [Default](https://godoc.org/github.com/tradyfinance/backoff#Default) function returns a default "safe" backoff policy, however it is better to configure a backoff policy for a specific use case.

## Example

```go
p := backoff.Default()
for {
    _, err := http.Get("http://example.com")
    if err != nil {
        break
    }
    p.Sleep()
}
p.Decrease()
```

## Policies

- [Default](https://godoc.org/github.com/tradyfinance/backoff#Default)
- [Nil](https://godoc.org/github.com/tradyfinance/backoff#Nil)
- [Constant](https://godoc.org/github.com/tradyfinance/backoff#Constant)
- [Linear](https://godoc.org/github.com/tradyfinance/backoff#Linear)
- [Exponential](https://godoc.org/github.com/tradyfinance/backoff#Exponential)
- [Jitter](https://godoc.org/github.com/tradyfinance/backoff#Jitter)

## Documentation

Documentation is available [here](https://godoc.org/github.com/tradyfinance/backoff).

## License

This project is released under the [Apache License, Version 2.0](LICENSE).
