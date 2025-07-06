# protoc-gen-enumlower

A `protoc` plugin that generates Go helper methods for enum types, allowing conversion between lower-case string representations and enum values.

## Features

- Adds `Parse<Type>Lower(s string) (<Type>, error)` to parse lower-case strings into enum values.
- Adds `LowerString() string` method to enum types for lower-case string output.

## Usage

1. Install the plugin:

   ```sh
   go install github.com/thejezzi/protoc-gen-enumlower@latest
   ```

2. Use with `protoc`:

   ```sh
   protoc --enumlower_out=. --enumlower_opt=paths=source_relative your.proto
   ```

3. Import the generated code in your Go project.

## Example

Given a proto enum:

```proto
enum Status {
  STATUS_UNKNOWN = 0;
  STATUS_ACTIVE = 1;
  STATUS_INACTIVE = 2;
}
```

The plugin generates:

```go
func ParseStatusLower(s string) (Status, error) { ... }
func (x Status) LowerString() string { ... }
```

Look under `examples` for more information and run `go generate ./...`

## Requirements

- Go 1.24.3+
- `google.golang.org/protobuf` v1.36.6

## License

MIT
