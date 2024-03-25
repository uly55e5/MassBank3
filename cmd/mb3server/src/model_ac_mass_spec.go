/*
 * MassBank3 API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package mb3server

type AcMassSpec struct {
	MsType string `json:"ms_type,omitempty"`

	IonMode string `json:"ion_mode,omitempty"`

	Subtags []AcMassSpecSubtagsInner `json:"subtags,omitempty"`
}

// AssertAcMassSpecRequired checks if the required fields are not zero-ed
func AssertAcMassSpecRequired(obj AcMassSpec) error {
	for _, el := range obj.Subtags {
		if err := AssertAcMassSpecSubtagsInnerRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseAcMassSpecRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of AcMassSpec (e.g. [][]AcMassSpec), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseAcMassSpecRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aAcMassSpec, ok := obj.(AcMassSpec)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertAcMassSpecRequired(aAcMassSpec)
	})
}
