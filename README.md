# go-repeat

Package repeat provides tools for repeat something, for the Go programming language.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-repeat

[![GoDoc](https://godoc.org/github.com/reiver/go-repeat?status.svg)](https://godoc.org/github.com/reiver/go-repeat)

## Example

Here is an example:

```golang
import "github.com/reiver/go-repeat"

// ...

// Do this forever, but wait between 4 hours to 5 hours between each time.
repeat.Forever(4 * time.Hour, 5 * time.Hour,

	// Do this until it succeeds, but only try doing it 3 times, and then give up.
	repeat.RedoUntilOKFunc(3,

		// Delay doing this between 0 milliseconds and 50 milliseconds, and then do it.
		repeat.DelayFunc(0, 50 * time.Millisecond,
			fn,
		),
	),
)
```

## Import

To import package **repeat** use `import` code like the follownig:
```
import "github.com/reiver/go-repeat"
```

## Installation

To install package **repeat** do the following:
```
GOPROXY=direct go get https://github.com/reiver/go-repeat
```

## Author

Package **repeat** was written by [Charles Iliya Krempeaux](http://reiver.link)
