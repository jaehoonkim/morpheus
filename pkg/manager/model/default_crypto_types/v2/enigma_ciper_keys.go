//go:generate go run github.com/abice/go-enum --file=enigma_ciper_keys.go --names --nocase=false
package v2

/*
	ENUM(

sentinel.default.crypto
)
*/
type CiperKey int
