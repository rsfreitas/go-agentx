/*
go-agentx
Copyright (C) 2015 Philipp Brüll <bruell@simia.tech>

This library is free software; you can redistribute it and/or
modify it under the terms of the GNU Lesser General Public
License as published by the Free Software Foundation; either
version 2.1 of the License, or (at your option) any later version.

This library is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public
License along with this library; if not, write to the Free Software
Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301
USA
*/

package pdu

import (
	"github.com/rsfreitas/go-agentx/marshaler"
	"gopkg.in/errgo.v1"
)

// Register defines the pdu register packet.
type Register struct {
	Timeout Timeout
	Context *OctetString
	Subtree ObjectIdentifier
}

// Type returns the pdu packet type.
func (r *Register) Type() Type {
	return TypeRegister
}

// MarshalBinary returns the pdu packet as a slice of bytes.
func (r *Register) MarshalBinary() ([]byte, error) {
	var combined marshaler.Multi

	if r.Context == nil {
		combined = marshaler.NewMulti(&r.Timeout, &r.Subtree)
	} else {
		combined = marshaler.NewMulti(r.Context, &r.Timeout, &r.Subtree)
	}

	combinedBytes, err := combined.MarshalBinary()
	if err != nil {
		return nil, errgo.Mask(err)
	}

	return combinedBytes, nil
}

// UnmarshalBinary sets the packet structure from the provided slice of bytes.
func (r *Register) UnmarshalBinary(data []byte) error {
	return nil
}
