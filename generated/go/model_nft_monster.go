/*
 * test
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type NftMonster struct {

	Scope GggDeveloper `json:"scope,omitempty"`
}

// AssertNftMonsterRequired checks if the required fields are not zero-ed
func AssertNftMonsterRequired(obj NftMonster) error {
	if err := AssertGggDeveloperRequired(obj.Scope); err != nil {
		return err
	}
	return nil
}

// AssertRecurseNftMonsterRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of NftMonster (e.g. [][]NftMonster), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseNftMonsterRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aNftMonster, ok := obj.(NftMonster)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertNftMonsterRequired(aNftMonster)
	})
}