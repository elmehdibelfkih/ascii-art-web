# ascii-art-web

# Go's `encoding/json` Package

The `encoding/json` package in Go provides functions for encoding data to JSON (JavaScript Object Notation) and decoding JSON into Go data structures. JSON is a lightweight format for data exchange, widely used in APIs and web applications.

---

## Key Features

- **Encoding**: Convert Go data types (structs, maps, slices, etc.) into JSON format.
- **Decoding**: Parse JSON data into Go data types.
- **Custom Marshaling/Unmarshaling**: Define how types are converted to/from JSON.
- **Stream Support**: Encode/decode directly from readers and writers for large data sets.

---

## Commonly Used Functions

### 1. Encoding Functions

#### `json.Marshal`

Converts a Go value into JSON-encoded data (a `[]byte`).

- **Input**: Go data types (structs, maps, slices, arrays, etc.)
- **Output**: JSON as a byte slice.
- **Usage**:
  - Useful for storing JSON or sending JSON in API responses.

#### `json.MarshalIndent`

Similar to `json.Marshal`, but produces human-readable indented JSON.

- **Input**: Go data types and indentation settings.
- **Output**: Pretty-printed JSON as a byte slice.
- **Usage**:
  - Suitable for configuration files or logs.

#### `json.NewEncoder`

Creates a `json.Encoder` to write JSON directly to an `io.Writer`.

- **Input**: An `io.Writer` (e.g., file, network connection).
- **Output**: Encodes JSON directly to the writer.
- **Usage**:
  - Ideal for streaming JSON data to files or over networks.

### 2. Decoding Functions

#### `json.Unmarshal`

Parses JSON-encoded data into a Go variable.

- **Input**: JSON data as a byte slice and a pointer to a Go variable.
- **Output**: Populates the Go variable with the parsed data.
- **Usage**:
  - Commonly used to parse API responses.

#### `json.NewDecoder`

Creates a `json.Decoder` to read and parse JSON directly from an `io.Reader`.

- **Input**: An `io.Reader` (e.g., file, network connection).
- **Output**: Decodes JSON directly from the reader.
- **Usage**:
  - Useful for streaming JSON data from files or over networks.

---

## JSON Tags for Structs

To control JSON encoding/decoding, Go allows struct fields to have tags:

- **Field Name**: Maps a struct field to a JSON key.
- **Omit Empty Fields**: Use `omitempty` to exclude fields with zero values.
- **Ignore Fields**: Use `-` to exclude a field entirely.

Example:

```go
 type User struct {
     Name      string `json:"name"`
     Age       int    `json:"age,omitempty"`
     Password  string `json:"-"`
 }
```

---

## Handling Custom Types

You can implement the `json.Marshaler` and `json.Unmarshaler` interfaces to define custom behavior for encoding and decoding.

- ``: Custom encoding logic.
- ``: Custom decoding logic.

---

## Error Handling

Both encoding and decoding functions return an error if something goes wrong (e.g., invalid JSON, unsupported types).

- **Common Errors**:
  - Invalid JSON syntax.
  - Type mismatches during decoding.
- **Best Practices**:
  - Always check for and handle errors.

---

## Stream Processing

The `Encoder` and `Decoder` types allow efficient processing of large or streaming JSON data.

### Encoder

Use `json.NewEncoder` with an `io.Writer` to write JSON incrementally.

### Decoder

Use `json.NewDecoder` with an `io.Reader` to read JSON incrementally, using methods like `Decode` or `Token` for fine-grained control.

---

## Best Practices

1. **Use Structs for Typed Data**: Define structs with appropriate tags for predictable behavior.
2. **Handle Errors Gracefully**: Always check for errors when using `Marshal`/`Unmarshal` or `Encoder`/`Decoder`.
3. **Omit Empty Fields**: Use `omitempty` to reduce JSON size by skipping zero-value fields.
4. **Pretty Print for Debugging**: Use `MarshalIndent` for logs or configuration files.
5. **Stream Large Data**: Use `Encoder` and `Decoder` for large datasets to avoid high memory usage.

---

## Example Workflow

### Encoding JSON

1. Create a struct or map.
2. Use `json.Marshal` or `json.NewEncoder` to convert it into JSON.
3. Handle the resulting `[]byte` or stream.

### Decoding JSON

1. Obtain JSON data from a file, API, or string.
2. Use `json.Unmarshal` or `json.NewDecoder` to parse it into a Go variable.
3. Access the data via the variable.

---

## References

- [Official Documentation](https://pkg.go.dev/encoding/json)
- [JSON Specification](https://www.json.org/)

==============================================================================

# Go's `encoding/json` Package

The `encoding/json` package in Go provides functions for encoding data to JSON (JavaScript Object Notation) and decoding JSON into Go data structures. JSON is a lightweight format for data exchange, widely used in APIs and web applications.

---

## Key Features

- **Encoding**: Convert Go data types (structs, maps, slices, etc.) into JSON format.
- **Decoding**: Parse JSON data into Go data types.
- **Custom Marshaling/Unmarshaling**: Define how types are converted to/from JSON.
- **Stream Support**: Encode/decode directly from readers and writers for large data sets.

---

## Commonly Used Functions

### 1. Encoding Functions

#### `json.Marshal`
Converts a Go value into JSON-encoded data (a `[]byte`).

- **Input**: Go data types (structs, maps, slices, arrays, etc.).
- **Output**: JSON as a byte slice.
- **Usage**:
  - Useful for storing JSON or sending JSON in API responses.
- **Key Details**:
  - Supports recursive encoding of nested structs, maps, and slices.
  - Automatically escapes special characters in strings.

#### `json.MarshalIndent`
Similar to `json.Marshal`, but produces human-readable indented JSON.

- **Input**: Go data types and indentation settings.
- **Output**: Pretty-printed JSON as a byte slice.
- **Usage**:
  - Suitable for configuration files or logs.
- **Key Details**:
  - Accepts prefix and indent string parameters to customize formatting.

#### `json.NewEncoder`
Creates a `json.Encoder` to write JSON directly to an `io.Writer`.

- **Input**: An `io.Writer` (e.g., file, network connection).
- **Output**: Encodes JSON directly to the writer.
- **Usage**:
  - Ideal for streaming JSON data to files or over networks.
- **Key Details**:
  - Includes methods like `SetEscapeHTML` to control escaping of HTML characters.

### 2. Decoding Functions

#### `json.Unmarshal`
Parses JSON-encoded data into a Go variable.

- **Input**: JSON data as a byte slice and a pointer to a Go variable.
- **Output**: Populates the Go variable with the parsed data.
- **Usage**:
  - Commonly used to parse API responses.
- **Key Details**:
  - Requires pointers to update variables directly.
  - Fails with an error if types mismatch between JSON and the Go variable.

#### `json.NewDecoder`
Creates a `json.Decoder` to read and parse JSON directly from an `io.Reader`.

- **Input**: An `io.Reader` (e.g., file, network connection).
- **Output**: Decodes JSON directly from the reader.
- **Usage**:
  - Useful for streaming JSON data from files or over networks.
- **Key Details**:
  - Supports incremental decoding with the `Token` method for finer control.
  - Can handle large JSON objects efficiently by processing them in chunks.

---

## JSON Tags for Structs

To control JSON encoding/decoding, Go allows struct fields to have tags:

- **Field Name**: Maps a struct field to a JSON key.
- **Omit Empty Fields**: Use `omitempty` to exclude fields with zero values.
- **Ignore Fields**: Use `-` to exclude a field entirely.

Example:
```go
 type User struct {
     Name      string `json:"name"`
     Age       int    `json:"age,omitempty"`
     Password  string `json:"-"`
 }
```

---

## Handling Custom Types

You can implement the `json.Marshaler` and `json.Unmarshaler` interfaces to define custom behavior for encoding and decoding.

- **`MarshalJSON`**: Custom encoding logic.
- **`UnmarshalJSON`**: Custom decoding logic.
- **Key Details**:
  - Provides full control over the JSON representation of a type.
  - Useful for complex types or non-standard data formats.

---

## Error Handling

Both encoding and decoding functions return an error if something goes wrong (e.g., invalid JSON, unsupported types).

- **Common Errors**:
  - Invalid JSON syntax.
  - Type mismatches during decoding.
- **Best Practices**:
  - Always check for and handle errors.
  - Validate JSON structure before decoding.

---

## Stream Processing

The `Encoder` and `Decoder` types allow efficient processing of large or streaming JSON data.

### Encoder
Use `json.NewEncoder` with an `io.Writer` to write JSON incrementally.

- **Methods**:
  - `Encode`: Encodes and writes a single value as JSON.
  - `SetEscapeHTML`: Controls escaping of HTML-specific characters.

### Decoder
Use `json.NewDecoder` with an `io.Reader` to read JSON incrementally.

- **Methods**:
  - `Decode`: Reads and decodes a single JSON value into a Go variable.
  - `Token`: Reads the next JSON token for fine-grained control.
  - `More`: Checks if there is more data to decode.

---

## Best Practices

1. **Use Structs for Typed Data**: Define structs with appropriate tags for predictable behavior.
2. **Handle Errors Gracefully**: Always check for errors when using `Marshal`/`Unmarshal` or `Encoder`/`Decoder`.
3. **Omit Empty Fields**: Use `omitempty` to reduce JSON size by skipping zero-value fields.
4. **Pretty Print for Debugging**: Use `MarshalIndent` for logs or configuration files.
5. **Stream Large Data**: Use `Encoder` and `Decoder` for large datasets to avoid high memory usage.

---

## Example Workflow

### Encoding JSON
1. Create a struct or map.
2. Use `json.Marshal` or `json.NewEncoder` to convert it into JSON.
3. Handle the resulting `[]byte` or stream.

### Decoding JSON
1. Obtain JSON data from a file, API, or string.
2. Use `json.Unmarshal` or `json.NewDecoder` to parse it into a Go variable.
3. Access the data via the variable.

---

## References

- [Official Documentation](https://pkg.go.dev/encoding/json)
- [JSON Specification](https://www.json.org/)

