package helper

import "errors"

// ValidateRelayCursor validates cursor params.
// allow: first, first + after, last, last + before
func ValidateRelayCursor(first *int, after *string, last *int, before *string) error {
	switch {
	case first != nil && after == nil && last == nil && before == nil,
		first != nil && after != nil && last == nil && before == nil,
		first == nil && after == nil && last != nil && before == nil,
		first == nil && after == nil && last != nil && before != nil:
		return nil
	}
	return errors.New("invalid cursor params: cursor allows first or first + after or last or last + before")
}
