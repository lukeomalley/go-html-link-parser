# Go HTML Link Parser

This package contains a simple link parser that takes in HTML as a string and returns all the <a> tags
within the document

## Usage

This package exposes one method **Parse** which takes in a reader and returns a slice of links

A link is structured as the following:

```go
  type Link struct {
    Href string
    Text string
  }
```

Check out the tests for an example usage.
