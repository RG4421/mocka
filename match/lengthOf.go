package match

import "reflect"

// LengthOf returns a new matcher that will match the length
// of strings, slices, arrays, and maps
func LengthOf(length int) SupportedKindsMatcher {
	return &lengthOf{length}
}

type lengthOf struct {
	length int
}

// SupportedKinds returns all the kinds the length of matcher supports
func (lengthOf) SupportedKinds() map[reflect.Kind]struct{} {
	return map[reflect.Kind]struct{}{
		reflect.Array:  {},
		reflect.Map:    {},
		reflect.Slice:  {},
		reflect.String: {},
	}
}

// Match return true if the length matches the provided length; otherwise false
func (m *lengthOf) Match(value interface{}) bool {
	if value == nil {
		return false
	}

	switch v := reflect.ValueOf(value); v.Kind() {
	case reflect.Array, reflect.Slice:
		return v.Len() == m.length
	case reflect.Map:
		return len(v.MapKeys()) == m.length
	case reflect.String:
		return len(value.(string)) == m.length
	default:
		return false
	}
}
