//go:build !no_gotd_slices
// +build !no_gotd_slices

// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
)

// SecureValueErrorClassArray is adapter for slice of SecureValueErrorClass.
type SecureValueErrorClassArray []SecureValueErrorClass

// Sort sorts slice of SecureValueErrorClass.
func (s SecureValueErrorClassArray) Sort(less func(a, b SecureValueErrorClass) bool) SecureValueErrorClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of SecureValueErrorClass.
func (s SecureValueErrorClassArray) SortStable(less func(a, b SecureValueErrorClass) bool) SecureValueErrorClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of SecureValueErrorClass.
func (s SecureValueErrorClassArray) Retain(keep func(x SecureValueErrorClass) bool) SecureValueErrorClassArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s SecureValueErrorClassArray) First() (v SecureValueErrorClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s SecureValueErrorClassArray) Last() (v SecureValueErrorClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *SecureValueErrorClassArray) PopFirst() (v SecureValueErrorClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero SecureValueErrorClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *SecureValueErrorClassArray) Pop() (v SecureValueErrorClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsSecureValueErrorData returns copy with only SecureValueErrorData constructors.
func (s SecureValueErrorClassArray) AsSecureValueErrorData() (to SecureValueErrorDataArray) {
	for _, elem := range s {
		value, ok := elem.(*SecureValueErrorData)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsSecureValueErrorFrontSide returns copy with only SecureValueErrorFrontSide constructors.
func (s SecureValueErrorClassArray) AsSecureValueErrorFrontSide() (to SecureValueErrorFrontSideArray) {
	for _, elem := range s {
		value, ok := elem.(*SecureValueErrorFrontSide)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsSecureValueErrorReverseSide returns copy with only SecureValueErrorReverseSide constructors.
func (s SecureValueErrorClassArray) AsSecureValueErrorReverseSide() (to SecureValueErrorReverseSideArray) {
	for _, elem := range s {
		value, ok := elem.(*SecureValueErrorReverseSide)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsSecureValueErrorSelfie returns copy with only SecureValueErrorSelfie constructors.
func (s SecureValueErrorClassArray) AsSecureValueErrorSelfie() (to SecureValueErrorSelfieArray) {
	for _, elem := range s {
		value, ok := elem.(*SecureValueErrorSelfie)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsSecureValueErrorFile returns copy with only SecureValueErrorFile constructors.
func (s SecureValueErrorClassArray) AsSecureValueErrorFile() (to SecureValueErrorFileArray) {
	for _, elem := range s {
		value, ok := elem.(*SecureValueErrorFile)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsSecureValueErrorFiles returns copy with only SecureValueErrorFiles constructors.
func (s SecureValueErrorClassArray) AsSecureValueErrorFiles() (to SecureValueErrorFilesArray) {
	for _, elem := range s {
		value, ok := elem.(*SecureValueErrorFiles)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsSecureValueError returns copy with only SecureValueError constructors.
func (s SecureValueErrorClassArray) AsSecureValueError() (to SecureValueErrorArray) {
	for _, elem := range s {
		value, ok := elem.(*SecureValueError)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsSecureValueErrorTranslationFile returns copy with only SecureValueErrorTranslationFile constructors.
func (s SecureValueErrorClassArray) AsSecureValueErrorTranslationFile() (to SecureValueErrorTranslationFileArray) {
	for _, elem := range s {
		value, ok := elem.(*SecureValueErrorTranslationFile)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsSecureValueErrorTranslationFiles returns copy with only SecureValueErrorTranslationFiles constructors.
func (s SecureValueErrorClassArray) AsSecureValueErrorTranslationFiles() (to SecureValueErrorTranslationFilesArray) {
	for _, elem := range s {
		value, ok := elem.(*SecureValueErrorTranslationFiles)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// SecureValueErrorDataArray is adapter for slice of SecureValueErrorData.
type SecureValueErrorDataArray []SecureValueErrorData

// Sort sorts slice of SecureValueErrorData.
func (s SecureValueErrorDataArray) Sort(less func(a, b SecureValueErrorData) bool) SecureValueErrorDataArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of SecureValueErrorData.
func (s SecureValueErrorDataArray) SortStable(less func(a, b SecureValueErrorData) bool) SecureValueErrorDataArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of SecureValueErrorData.
func (s SecureValueErrorDataArray) Retain(keep func(x SecureValueErrorData) bool) SecureValueErrorDataArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s SecureValueErrorDataArray) First() (v SecureValueErrorData, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s SecureValueErrorDataArray) Last() (v SecureValueErrorData, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *SecureValueErrorDataArray) PopFirst() (v SecureValueErrorData, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero SecureValueErrorData
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *SecureValueErrorDataArray) Pop() (v SecureValueErrorData, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SecureValueErrorFrontSideArray is adapter for slice of SecureValueErrorFrontSide.
type SecureValueErrorFrontSideArray []SecureValueErrorFrontSide

// Sort sorts slice of SecureValueErrorFrontSide.
func (s SecureValueErrorFrontSideArray) Sort(less func(a, b SecureValueErrorFrontSide) bool) SecureValueErrorFrontSideArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of SecureValueErrorFrontSide.
func (s SecureValueErrorFrontSideArray) SortStable(less func(a, b SecureValueErrorFrontSide) bool) SecureValueErrorFrontSideArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of SecureValueErrorFrontSide.
func (s SecureValueErrorFrontSideArray) Retain(keep func(x SecureValueErrorFrontSide) bool) SecureValueErrorFrontSideArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s SecureValueErrorFrontSideArray) First() (v SecureValueErrorFrontSide, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s SecureValueErrorFrontSideArray) Last() (v SecureValueErrorFrontSide, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *SecureValueErrorFrontSideArray) PopFirst() (v SecureValueErrorFrontSide, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero SecureValueErrorFrontSide
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *SecureValueErrorFrontSideArray) Pop() (v SecureValueErrorFrontSide, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SecureValueErrorReverseSideArray is adapter for slice of SecureValueErrorReverseSide.
type SecureValueErrorReverseSideArray []SecureValueErrorReverseSide

// Sort sorts slice of SecureValueErrorReverseSide.
func (s SecureValueErrorReverseSideArray) Sort(less func(a, b SecureValueErrorReverseSide) bool) SecureValueErrorReverseSideArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of SecureValueErrorReverseSide.
func (s SecureValueErrorReverseSideArray) SortStable(less func(a, b SecureValueErrorReverseSide) bool) SecureValueErrorReverseSideArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of SecureValueErrorReverseSide.
func (s SecureValueErrorReverseSideArray) Retain(keep func(x SecureValueErrorReverseSide) bool) SecureValueErrorReverseSideArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s SecureValueErrorReverseSideArray) First() (v SecureValueErrorReverseSide, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s SecureValueErrorReverseSideArray) Last() (v SecureValueErrorReverseSide, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *SecureValueErrorReverseSideArray) PopFirst() (v SecureValueErrorReverseSide, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero SecureValueErrorReverseSide
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *SecureValueErrorReverseSideArray) Pop() (v SecureValueErrorReverseSide, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SecureValueErrorSelfieArray is adapter for slice of SecureValueErrorSelfie.
type SecureValueErrorSelfieArray []SecureValueErrorSelfie

// Sort sorts slice of SecureValueErrorSelfie.
func (s SecureValueErrorSelfieArray) Sort(less func(a, b SecureValueErrorSelfie) bool) SecureValueErrorSelfieArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of SecureValueErrorSelfie.
func (s SecureValueErrorSelfieArray) SortStable(less func(a, b SecureValueErrorSelfie) bool) SecureValueErrorSelfieArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of SecureValueErrorSelfie.
func (s SecureValueErrorSelfieArray) Retain(keep func(x SecureValueErrorSelfie) bool) SecureValueErrorSelfieArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s SecureValueErrorSelfieArray) First() (v SecureValueErrorSelfie, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s SecureValueErrorSelfieArray) Last() (v SecureValueErrorSelfie, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *SecureValueErrorSelfieArray) PopFirst() (v SecureValueErrorSelfie, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero SecureValueErrorSelfie
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *SecureValueErrorSelfieArray) Pop() (v SecureValueErrorSelfie, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SecureValueErrorFileArray is adapter for slice of SecureValueErrorFile.
type SecureValueErrorFileArray []SecureValueErrorFile

// Sort sorts slice of SecureValueErrorFile.
func (s SecureValueErrorFileArray) Sort(less func(a, b SecureValueErrorFile) bool) SecureValueErrorFileArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of SecureValueErrorFile.
func (s SecureValueErrorFileArray) SortStable(less func(a, b SecureValueErrorFile) bool) SecureValueErrorFileArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of SecureValueErrorFile.
func (s SecureValueErrorFileArray) Retain(keep func(x SecureValueErrorFile) bool) SecureValueErrorFileArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s SecureValueErrorFileArray) First() (v SecureValueErrorFile, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s SecureValueErrorFileArray) Last() (v SecureValueErrorFile, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *SecureValueErrorFileArray) PopFirst() (v SecureValueErrorFile, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero SecureValueErrorFile
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *SecureValueErrorFileArray) Pop() (v SecureValueErrorFile, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SecureValueErrorFilesArray is adapter for slice of SecureValueErrorFiles.
type SecureValueErrorFilesArray []SecureValueErrorFiles

// Sort sorts slice of SecureValueErrorFiles.
func (s SecureValueErrorFilesArray) Sort(less func(a, b SecureValueErrorFiles) bool) SecureValueErrorFilesArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of SecureValueErrorFiles.
func (s SecureValueErrorFilesArray) SortStable(less func(a, b SecureValueErrorFiles) bool) SecureValueErrorFilesArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of SecureValueErrorFiles.
func (s SecureValueErrorFilesArray) Retain(keep func(x SecureValueErrorFiles) bool) SecureValueErrorFilesArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s SecureValueErrorFilesArray) First() (v SecureValueErrorFiles, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s SecureValueErrorFilesArray) Last() (v SecureValueErrorFiles, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *SecureValueErrorFilesArray) PopFirst() (v SecureValueErrorFiles, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero SecureValueErrorFiles
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *SecureValueErrorFilesArray) Pop() (v SecureValueErrorFiles, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SecureValueErrorArray is adapter for slice of SecureValueError.
type SecureValueErrorArray []SecureValueError

// Sort sorts slice of SecureValueError.
func (s SecureValueErrorArray) Sort(less func(a, b SecureValueError) bool) SecureValueErrorArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of SecureValueError.
func (s SecureValueErrorArray) SortStable(less func(a, b SecureValueError) bool) SecureValueErrorArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of SecureValueError.
func (s SecureValueErrorArray) Retain(keep func(x SecureValueError) bool) SecureValueErrorArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s SecureValueErrorArray) First() (v SecureValueError, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s SecureValueErrorArray) Last() (v SecureValueError, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *SecureValueErrorArray) PopFirst() (v SecureValueError, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero SecureValueError
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *SecureValueErrorArray) Pop() (v SecureValueError, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SecureValueErrorTranslationFileArray is adapter for slice of SecureValueErrorTranslationFile.
type SecureValueErrorTranslationFileArray []SecureValueErrorTranslationFile

// Sort sorts slice of SecureValueErrorTranslationFile.
func (s SecureValueErrorTranslationFileArray) Sort(less func(a, b SecureValueErrorTranslationFile) bool) SecureValueErrorTranslationFileArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of SecureValueErrorTranslationFile.
func (s SecureValueErrorTranslationFileArray) SortStable(less func(a, b SecureValueErrorTranslationFile) bool) SecureValueErrorTranslationFileArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of SecureValueErrorTranslationFile.
func (s SecureValueErrorTranslationFileArray) Retain(keep func(x SecureValueErrorTranslationFile) bool) SecureValueErrorTranslationFileArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s SecureValueErrorTranslationFileArray) First() (v SecureValueErrorTranslationFile, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s SecureValueErrorTranslationFileArray) Last() (v SecureValueErrorTranslationFile, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *SecureValueErrorTranslationFileArray) PopFirst() (v SecureValueErrorTranslationFile, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero SecureValueErrorTranslationFile
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *SecureValueErrorTranslationFileArray) Pop() (v SecureValueErrorTranslationFile, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SecureValueErrorTranslationFilesArray is adapter for slice of SecureValueErrorTranslationFiles.
type SecureValueErrorTranslationFilesArray []SecureValueErrorTranslationFiles

// Sort sorts slice of SecureValueErrorTranslationFiles.
func (s SecureValueErrorTranslationFilesArray) Sort(less func(a, b SecureValueErrorTranslationFiles) bool) SecureValueErrorTranslationFilesArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of SecureValueErrorTranslationFiles.
func (s SecureValueErrorTranslationFilesArray) SortStable(less func(a, b SecureValueErrorTranslationFiles) bool) SecureValueErrorTranslationFilesArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of SecureValueErrorTranslationFiles.
func (s SecureValueErrorTranslationFilesArray) Retain(keep func(x SecureValueErrorTranslationFiles) bool) SecureValueErrorTranslationFilesArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s SecureValueErrorTranslationFilesArray) First() (v SecureValueErrorTranslationFiles, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s SecureValueErrorTranslationFilesArray) Last() (v SecureValueErrorTranslationFiles, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *SecureValueErrorTranslationFilesArray) PopFirst() (v SecureValueErrorTranslationFiles, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero SecureValueErrorTranslationFiles
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *SecureValueErrorTranslationFilesArray) Pop() (v SecureValueErrorTranslationFiles, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}