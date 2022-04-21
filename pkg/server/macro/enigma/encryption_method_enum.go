// Code generated by go-enum DO NOT EDIT.
// Version:
// Revision:
// Build Date:
// Built By:

package enigma

import (
	"fmt"
	"strings"
)

const (
	// EncryptionMethodNONE is a EncryptionMethod of type NONE.
	EncryptionMethodNONE EncryptionMethod = iota
	// EncryptionMethodAES is a EncryptionMethod of type AES.
	EncryptionMethodAES
	// EncryptionMethodDES is a EncryptionMethod of type DES.
	EncryptionMethodDES
)

const _EncryptionMethodName = "NONEAESDES"

var _EncryptionMethodNames = []string{
	_EncryptionMethodName[0:4],
	_EncryptionMethodName[4:7],
	_EncryptionMethodName[7:10],
}

// EncryptionMethodNames returns a list of possible string values of EncryptionMethod.
func EncryptionMethodNames() []string {
	tmp := make([]string, len(_EncryptionMethodNames))
	copy(tmp, _EncryptionMethodNames)
	return tmp
}

var _EncryptionMethodMap = map[EncryptionMethod]string{
	EncryptionMethodNONE: _EncryptionMethodName[0:4],
	EncryptionMethodAES:  _EncryptionMethodName[4:7],
	EncryptionMethodDES:  _EncryptionMethodName[7:10],
}

// String implements the Stringer interface.
func (x EncryptionMethod) String() string {
	if str, ok := _EncryptionMethodMap[x]; ok {
		return str
	}
	return fmt.Sprintf("EncryptionMethod(%d)", x)
}

var _EncryptionMethodValue = map[string]EncryptionMethod{
	_EncryptionMethodName[0:4]:                   EncryptionMethodNONE,
	strings.ToLower(_EncryptionMethodName[0:4]):  EncryptionMethodNONE,
	_EncryptionMethodName[4:7]:                   EncryptionMethodAES,
	strings.ToLower(_EncryptionMethodName[4:7]):  EncryptionMethodAES,
	_EncryptionMethodName[7:10]:                  EncryptionMethodDES,
	strings.ToLower(_EncryptionMethodName[7:10]): EncryptionMethodDES,
}

// ParseEncryptionMethod attempts to convert a string to a EncryptionMethod.
func ParseEncryptionMethod(name string) (EncryptionMethod, error) {
	if x, ok := _EncryptionMethodValue[name]; ok {
		return x, nil
	}
	// Case insensitive parse, do a separate lookup to prevent unnecessary cost of lowercasing a string if we don't need to.
	if x, ok := _EncryptionMethodValue[strings.ToLower(name)]; ok {
		return x, nil
	}
	return EncryptionMethod(0), fmt.Errorf("%s is not a valid EncryptionMethod, try [%s]", name, strings.Join(_EncryptionMethodNames, ", "))
}