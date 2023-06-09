/*
 * test
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type GggScopeEnum string

// List of GggScopeEnum
const (
	GGGSCOPEENUM_FOO GggScopeEnum = "foo"
	GGGSCOPEENUM_BAR GggScopeEnum = "bar"
)

// AssertGggScopeEnumRequired checks if the required fields are not zero-ed
func AssertGggScopeEnumRequired(obj GggScopeEnum) error {
	return nil
}

// AssertRecurseGggScopeEnumRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of GggScopeEnum (e.g. [][]GggScopeEnum), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseGggScopeEnumRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aGggScopeEnum, ok := obj.(GggScopeEnum)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertGggScopeEnumRequired(aGggScopeEnum)
	})
}
