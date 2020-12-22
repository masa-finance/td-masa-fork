// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"
	"strings"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}

// InputEncryptedFileEmpty represents TL type `inputEncryptedFileEmpty#1837c364`.
// Empty constructor.
//
// See https://core.telegram.org/constructor/inputEncryptedFileEmpty for reference.
type InputEncryptedFileEmpty struct {
}

// InputEncryptedFileEmptyTypeID is TL type id of InputEncryptedFileEmpty.
const InputEncryptedFileEmptyTypeID = 0x1837c364

// String implements fmt.Stringer.
func (i *InputEncryptedFileEmpty) String() string {
	if i == nil {
		return "InputEncryptedFileEmpty(nil)"
	}
	var sb strings.Builder
	sb.WriteString("InputEncryptedFileEmpty")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (i *InputEncryptedFileEmpty) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputEncryptedFileEmpty#1837c364 as nil")
	}
	b.PutID(InputEncryptedFileEmptyTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputEncryptedFileEmpty) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputEncryptedFileEmpty#1837c364 to nil")
	}
	if err := b.ConsumeID(InputEncryptedFileEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode inputEncryptedFileEmpty#1837c364: %w", err)
	}
	return nil
}

// construct implements constructor of InputEncryptedFileClass.
func (i InputEncryptedFileEmpty) construct() InputEncryptedFileClass { return &i }

// Ensuring interfaces in compile-time for InputEncryptedFileEmpty.
var (
	_ bin.Encoder = &InputEncryptedFileEmpty{}
	_ bin.Decoder = &InputEncryptedFileEmpty{}

	_ InputEncryptedFileClass = &InputEncryptedFileEmpty{}
)

// InputEncryptedFileUploaded represents TL type `inputEncryptedFileUploaded#64bd0306`.
// Sets new encrypted file saved by parts using upload.saveFilePart method.
//
// See https://core.telegram.org/constructor/inputEncryptedFileUploaded for reference.
type InputEncryptedFileUploaded struct {
	// Random file ID created by clien
	ID int64
	// Number of saved parts
	Parts int
	// In case md5-HASH of the (already encrypted) file was transmitted, file content will be checked prior to use
	Md5Checksum string
	// 32-bit fingerprint of the key used to encrypt a file
	KeyFingerprint int
}

// InputEncryptedFileUploadedTypeID is TL type id of InputEncryptedFileUploaded.
const InputEncryptedFileUploadedTypeID = 0x64bd0306

// String implements fmt.Stringer.
func (i *InputEncryptedFileUploaded) String() string {
	if i == nil {
		return "InputEncryptedFileUploaded(nil)"
	}
	var sb strings.Builder
	sb.WriteString("InputEncryptedFileUploaded")
	sb.WriteString("{\n")
	sb.WriteString("\tID: ")
	sb.WriteString(fmt.Sprint(i.ID))
	sb.WriteString(",\n")
	sb.WriteString("\tParts: ")
	sb.WriteString(fmt.Sprint(i.Parts))
	sb.WriteString(",\n")
	sb.WriteString("\tMd5Checksum: ")
	sb.WriteString(fmt.Sprint(i.Md5Checksum))
	sb.WriteString(",\n")
	sb.WriteString("\tKeyFingerprint: ")
	sb.WriteString(fmt.Sprint(i.KeyFingerprint))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (i *InputEncryptedFileUploaded) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputEncryptedFileUploaded#64bd0306 as nil")
	}
	b.PutID(InputEncryptedFileUploadedTypeID)
	b.PutLong(i.ID)
	b.PutInt(i.Parts)
	b.PutString(i.Md5Checksum)
	b.PutInt(i.KeyFingerprint)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputEncryptedFileUploaded) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputEncryptedFileUploaded#64bd0306 to nil")
	}
	if err := b.ConsumeID(InputEncryptedFileUploadedTypeID); err != nil {
		return fmt.Errorf("unable to decode inputEncryptedFileUploaded#64bd0306: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputEncryptedFileUploaded#64bd0306: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputEncryptedFileUploaded#64bd0306: field parts: %w", err)
		}
		i.Parts = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputEncryptedFileUploaded#64bd0306: field md5_checksum: %w", err)
		}
		i.Md5Checksum = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputEncryptedFileUploaded#64bd0306: field key_fingerprint: %w", err)
		}
		i.KeyFingerprint = value
	}
	return nil
}

// construct implements constructor of InputEncryptedFileClass.
func (i InputEncryptedFileUploaded) construct() InputEncryptedFileClass { return &i }

// Ensuring interfaces in compile-time for InputEncryptedFileUploaded.
var (
	_ bin.Encoder = &InputEncryptedFileUploaded{}
	_ bin.Decoder = &InputEncryptedFileUploaded{}

	_ InputEncryptedFileClass = &InputEncryptedFileUploaded{}
)

// InputEncryptedFile represents TL type `inputEncryptedFile#5a17b5e5`.
// Sets forwarded encrypted file for attachment.
//
// See https://core.telegram.org/constructor/inputEncryptedFile for reference.
type InputEncryptedFile struct {
	// File ID, value of id parameter from encryptedFile
	ID int64
	// Checking sum, value of access_hash parameter from encryptedFile
	AccessHash int64
}

// InputEncryptedFileTypeID is TL type id of InputEncryptedFile.
const InputEncryptedFileTypeID = 0x5a17b5e5

// String implements fmt.Stringer.
func (i *InputEncryptedFile) String() string {
	if i == nil {
		return "InputEncryptedFile(nil)"
	}
	var sb strings.Builder
	sb.WriteString("InputEncryptedFile")
	sb.WriteString("{\n")
	sb.WriteString("\tID: ")
	sb.WriteString(fmt.Sprint(i.ID))
	sb.WriteString(",\n")
	sb.WriteString("\tAccessHash: ")
	sb.WriteString(fmt.Sprint(i.AccessHash))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (i *InputEncryptedFile) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputEncryptedFile#5a17b5e5 as nil")
	}
	b.PutID(InputEncryptedFileTypeID)
	b.PutLong(i.ID)
	b.PutLong(i.AccessHash)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputEncryptedFile) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputEncryptedFile#5a17b5e5 to nil")
	}
	if err := b.ConsumeID(InputEncryptedFileTypeID); err != nil {
		return fmt.Errorf("unable to decode inputEncryptedFile#5a17b5e5: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputEncryptedFile#5a17b5e5: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputEncryptedFile#5a17b5e5: field access_hash: %w", err)
		}
		i.AccessHash = value
	}
	return nil
}

// construct implements constructor of InputEncryptedFileClass.
func (i InputEncryptedFile) construct() InputEncryptedFileClass { return &i }

// Ensuring interfaces in compile-time for InputEncryptedFile.
var (
	_ bin.Encoder = &InputEncryptedFile{}
	_ bin.Decoder = &InputEncryptedFile{}

	_ InputEncryptedFileClass = &InputEncryptedFile{}
)

// InputEncryptedFileBigUploaded represents TL type `inputEncryptedFileBigUploaded#2dc173c8`.
// Assigns a new big encrypted file (over 10Mb in size), saved in parts using the method upload.saveBigFilePart.
//
// See https://core.telegram.org/constructor/inputEncryptedFileBigUploaded for reference.
type InputEncryptedFileBigUploaded struct {
	// Random file id, created by the client
	ID int64
	// Number of saved parts
	Parts int
	// 32-bit imprint of the key used to encrypt the file
	KeyFingerprint int
}

// InputEncryptedFileBigUploadedTypeID is TL type id of InputEncryptedFileBigUploaded.
const InputEncryptedFileBigUploadedTypeID = 0x2dc173c8

// String implements fmt.Stringer.
func (i *InputEncryptedFileBigUploaded) String() string {
	if i == nil {
		return "InputEncryptedFileBigUploaded(nil)"
	}
	var sb strings.Builder
	sb.WriteString("InputEncryptedFileBigUploaded")
	sb.WriteString("{\n")
	sb.WriteString("\tID: ")
	sb.WriteString(fmt.Sprint(i.ID))
	sb.WriteString(",\n")
	sb.WriteString("\tParts: ")
	sb.WriteString(fmt.Sprint(i.Parts))
	sb.WriteString(",\n")
	sb.WriteString("\tKeyFingerprint: ")
	sb.WriteString(fmt.Sprint(i.KeyFingerprint))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (i *InputEncryptedFileBigUploaded) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputEncryptedFileBigUploaded#2dc173c8 as nil")
	}
	b.PutID(InputEncryptedFileBigUploadedTypeID)
	b.PutLong(i.ID)
	b.PutInt(i.Parts)
	b.PutInt(i.KeyFingerprint)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputEncryptedFileBigUploaded) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputEncryptedFileBigUploaded#2dc173c8 to nil")
	}
	if err := b.ConsumeID(InputEncryptedFileBigUploadedTypeID); err != nil {
		return fmt.Errorf("unable to decode inputEncryptedFileBigUploaded#2dc173c8: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputEncryptedFileBigUploaded#2dc173c8: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputEncryptedFileBigUploaded#2dc173c8: field parts: %w", err)
		}
		i.Parts = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputEncryptedFileBigUploaded#2dc173c8: field key_fingerprint: %w", err)
		}
		i.KeyFingerprint = value
	}
	return nil
}

// construct implements constructor of InputEncryptedFileClass.
func (i InputEncryptedFileBigUploaded) construct() InputEncryptedFileClass { return &i }

// Ensuring interfaces in compile-time for InputEncryptedFileBigUploaded.
var (
	_ bin.Encoder = &InputEncryptedFileBigUploaded{}
	_ bin.Decoder = &InputEncryptedFileBigUploaded{}

	_ InputEncryptedFileClass = &InputEncryptedFileBigUploaded{}
)

// InputEncryptedFileClass represents InputEncryptedFile generic type.
//
// See https://core.telegram.org/type/InputEncryptedFile for reference.
//
// Example:
//  g, err := DecodeInputEncryptedFile(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *InputEncryptedFileEmpty: // inputEncryptedFileEmpty#1837c364
//  case *InputEncryptedFileUploaded: // inputEncryptedFileUploaded#64bd0306
//  case *InputEncryptedFile: // inputEncryptedFile#5a17b5e5
//  case *InputEncryptedFileBigUploaded: // inputEncryptedFileBigUploaded#2dc173c8
//  default: panic(v)
//  }
type InputEncryptedFileClass interface {
	bin.Encoder
	bin.Decoder
	construct() InputEncryptedFileClass
	fmt.Stringer
}

// DecodeInputEncryptedFile implements binary de-serialization for InputEncryptedFileClass.
func DecodeInputEncryptedFile(buf *bin.Buffer) (InputEncryptedFileClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputEncryptedFileEmptyTypeID:
		// Decoding inputEncryptedFileEmpty#1837c364.
		v := InputEncryptedFileEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputEncryptedFileClass: %w", err)
		}
		return &v, nil
	case InputEncryptedFileUploadedTypeID:
		// Decoding inputEncryptedFileUploaded#64bd0306.
		v := InputEncryptedFileUploaded{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputEncryptedFileClass: %w", err)
		}
		return &v, nil
	case InputEncryptedFileTypeID:
		// Decoding inputEncryptedFile#5a17b5e5.
		v := InputEncryptedFile{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputEncryptedFileClass: %w", err)
		}
		return &v, nil
	case InputEncryptedFileBigUploadedTypeID:
		// Decoding inputEncryptedFileBigUploaded#2dc173c8.
		v := InputEncryptedFileBigUploaded{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputEncryptedFileClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputEncryptedFileClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputEncryptedFile boxes the InputEncryptedFileClass providing a helper.
type InputEncryptedFileBox struct {
	InputEncryptedFile InputEncryptedFileClass
}

// Decode implements bin.Decoder for InputEncryptedFileBox.
func (b *InputEncryptedFileBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputEncryptedFileBox to nil")
	}
	v, err := DecodeInputEncryptedFile(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputEncryptedFile = v
	return nil
}

// Encode implements bin.Encode for InputEncryptedFileBox.
func (b *InputEncryptedFileBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputEncryptedFile == nil {
		return fmt.Errorf("unable to encode InputEncryptedFileClass as nil")
	}
	return b.InputEncryptedFile.Encode(buf)
}
