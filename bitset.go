// Package bitset implments a memory efficient bit array of booleans

package bitset

import "fmt"

type BitSet struct {
	bits []uint8
	size int
}

const (
	bitMaskZero = uint8(0)
	bitMaskOnes = uint8((1 << 8) - 1)
)

var (
	bitMasks = [...]uint8{0x1, 0x2, 0x4, 0x8, 0x10, 0x20, 0x40, 0x80}
)

func (b *BitSet) getPositionAndMask(index int) (int, uint8) {
	if index < 0 || index >= b.size {
		panic(fmt.Errorf("BitSet index (%d) out of bounds (size: %d)", index, b.size))
	}
	position := index >> 3
	mask := bitMasks[index%8]
	return position, mask
}

func (b *BitSet) Init(size int) {
	b.bits = make([]uint8, (size+7)/8)
	b.size = size
}

func (b *BitSet) Size() int {
	return b.size
}

func (b *BitSet) Get(index int) bool {
	position, mask := b.getPositionAndMask(index)
	return (b.bits[position] & mask) != 0
}

func (b *BitSet) Set(index int) {
	position, mask := b.getPositionAndMask(index)
	b.bits[position] |= mask
}

func (b *BitSet) SetAll() {
	for i := 0; i < len(b.bits); i++ {
		b.bits[i] = bitMaskOnes
	}
}

func (b *BitSet) Invert(index int) {
	position, mask := b.getPositionAndMask(index)
	b.bits[position] ^= mask
}

func (b *BitSet) InvertAll() {
	for i := 0; i < len(b.bits); i++ {
		b.bits[i] = ^b.bits[i]
	}
}

func (b *BitSet) Clear(index int) {
	position, mask := b.getPositionAndMask(index)
	b.bits[position] &^= mask
}

func (b *BitSet) ClearAll() {
	for i := 0; i < len(b.bits); i++ {
		b.bits[i] = bitMaskZero
	}
}

func (b *BitSet) String() string {
	value := make([]byte, b.size)
	for i := 0; i < b.size; i++ {
		if b.Get(i) {
			value[i] = '1'
		} else {
			value[i] = '0'
		}
	}
	return string(value)
}

func New(size int) *BitSet {
	b := &BitSet{}
	b.Init(size)
	return b
}

func Parse(s string) (*BitSet, error) {
	b := New(len(s))
	for i, c := range s {
		if c == '0' {
			b.Clear(i)
		} else if c == '1' {
			b.Set(i)
		} else {
			return nil, fmt.Errorf("Invalid binary character: %c", c)
		}
	}
	return b, nil
}
