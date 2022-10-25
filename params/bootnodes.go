// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import "github.com/ethereum/go-ethereum/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// Ethereum Foundation Go Bootnodes//TODO
	"enode://6361df9418eece9f434599a6b210739d81440e52b2706f605f304de80c6cf7ea1c292e96376cdaac8bcf7d06c6a9125dfe4db3ca8aa4c8ecedfa55f77e2476fc@18.163.19.88:33668",
	"enode://65e8916cc0184acb14bd3b545b1d60296b822afe2bb3ee5476da292613221c18a70eae2db10f9c386e31668258c3225a7339dbda58db865060da1e69c8a0375f@18.167.78.35:33668",
	"enode://23b98693149c5488b14d7214e3261c02e728c800f8540bb5fcd67d2fcfe1e5b1647c89fc0be90d3f8e4033366eb816d52b02688d202111baa27594224ee75f8e@43.198.34.98:33668",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
var TestnetBootnodes = []string{
	"enode://fcfd4d35045391606c98fefefeed793977ccc8b50788d5a1dfb93116aa14370453d0bd12890c29b596906e35ddb0d8d1956f9329fc2d0a80f5c87b604069584c@16.162.167.156:33668",
	"enode://bd857e298d8434e71182765d8c2cd23bea5d745bb31ae0d94f1c991423c4a11f9ec9710a465008bbd788e1c6b98e41805744f5dc5ac8bfebd6156d99be441e57@43.198.46.101:33668",
}

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	return ""
}
