package hashtable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHappy(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   func(table *HashTable)
		out  func(table *HashTable) bool
	}{
		{
			name: "1",
			in: func(table *HashTable) {
				table.Put(1, 2)
				table.Put(1, 2)
			},
			out: func(table *HashTable) bool {
				position := table.hash(1, len(table.slots))
				if table.slots[position] == nil {
					return false
				}
				if table.slots[position].Next != nil {
					return false
				}
				value, ok := table.Get(1)
				if !ok || value != 2 {
					return false
				}
				return true
			},
		},
		{
			name: "2",
			in: func(table *HashTable) {
				table.Put(1, 2)
				table.Put(1, 2)
				table.Put(11, 22)
			},
			out: func(table *HashTable) bool {
				position := table.hash(1, len(table.slots))
				if table.slots[position] == nil {
					return false
				}
				if table.slots[position].Next == nil {
					return false
				}
				if table.slots[position].Next.Next != nil {
					return false
				}
				value, ok := table.Get(1)
				if !ok || value != 2 {
					return false
				}
				value, ok = table.Get(11)
				if !ok || value != 22 {
					return false
				}
				return true
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			table := New()
			tc.in(table)
			assert.True(t, tc.out(table))
		})
	}
}
