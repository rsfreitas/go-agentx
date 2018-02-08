/*
go-agentx
Copyright (C) 2018 Vadim Ipatov <vadim.ipatov@zabbix.com>

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

// Notify defines the pdu notify packet.
type Notify struct {
	Variables Variables
}

// Type returns the pdu packet type.
func (n *Notify) Type() Type {
	return TypeNotify
}

// MarshalBinary returns the pdu packet as a slice of bytes.
func (n *Notify) MarshalBinary() ([]byte, error) {
	data, err := n.Variables.MarshalBinary()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// UnmarshalBinary sets the packet structure from the provided slice of bytes.
func (n *Notify) UnmarshalBinary(data []byte) error {
	return nil
}
