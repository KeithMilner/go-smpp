// Copyright 2015 go-smpp authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package pdutext

// Binary Class 2 codec.
type BinaryClass2 []byte

// Type implements the Codec interface.
func (s BinaryClass2) Type() DataCoding {
	return BinaryClass2Type
}

// Encode binary.
func (s BinaryClass2) Encode() []byte {
	return s
}

// Decode binary.
func (s BinaryClass2) Decode() []byte {
	return s
}
