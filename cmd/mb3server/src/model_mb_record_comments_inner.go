/*
 * MassBank3 API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package mb3server

type MbRecordCommentsInner struct {

	// The subtag to describe the comment
	Subtag string `json:"subtag,omitempty"`

	// The comment
	Value string `json:"value,omitempty"`
}

// AssertMbRecordCommentsInnerRequired checks if the required fields are not zero-ed
func AssertMbRecordCommentsInnerRequired(obj MbRecordCommentsInner) error {
	return nil
}

// AssertRecurseMbRecordCommentsInnerRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of MbRecordCommentsInner (e.g. [][]MbRecordCommentsInner), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseMbRecordCommentsInnerRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aMbRecordCommentsInner, ok := obj.(MbRecordCommentsInner)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertMbRecordCommentsInnerRequired(aMbRecordCommentsInner)
	})
}
