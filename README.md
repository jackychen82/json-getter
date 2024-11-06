# json-getter

jsongetter is a Go package that provides functionality for retrieving nested values from JSON data by specifying a dot-separated path.

## Features
- Parse JSON data into a map structure.
- Retrieve nested values in JSON by specifying a path using dot notation (e.g., "parent.child").
- Supports accessing array elements in JSON by using indices (e.g., "array.[0]").

## Installation

Add jsongetter to your Go project:

```bash
go get github.com/jackychen82/json-getter
```

Usage
```go
package main

import (
    "fmt"
    "jsongetter"
)

func main() {
    jsonData := []byte(`{
        "name": "John",
        "address": {
            "city": "New York",
            "zip": "10001"
        },
        "tags": ["developer", "golang"]
    }`)

    // Retrieve top-level field
    name := jsongetter.GetNodeValue("name", jsonData)
    fmt.Println("Name:", name) // Output: Name: John

    // Retrieve nested field
    city := jsongetter.GetNodeValue("address.city", jsonData)
    fmt.Println("City:", city) // Output: City: New York

    // Access an array element
    tag := jsongetter.GetNodeValue("tags.[0]", jsonData)
    fmt.Println("First tag:", tag) // Output: First tag: developer
}
```

## Function Reference

### GetNodeValue(path string, in []byte) any

Retrieves the value at the specified path in the JSON data.

- Parameters:
  - path: Dot-separated path to the value.
  - in: JSON data as a []byte.
- Returns: The value at the specified path, or nil if the path is invalid or the value does not exist.

### jsonToMap(in []byte) (map[string]any, error)
Converts JSON data into a map for easier path-based access.

- Parameters:
  - in: JSON data as a []byte.
  - Returns: A map representing the JSON structure.

## License
This project is licensed under the MIT License.