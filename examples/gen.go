//go:generate go install github.com/thejezzi/protoc-gen-enumlower/cmd/protoc-gen-enumlower@latest
//go:generate go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
//go:generate protoc --go_out=. --go_opt=paths=source_relative --enumlower_out=. --enumlower_opt=paths=source_relative example.proto

package examples
