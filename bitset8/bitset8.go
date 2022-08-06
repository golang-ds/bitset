package bitset

type Bitset8 uint8

// New returns an empty 8-bit bitset.
func New() (b Bitset8) {
	return
}

// Add adds the given number to bitset (changes the bit in that position to 1). It returns ErrNumberNotFit8 error
// if the given number doesn't fit 8-bit bitset.
// time-complexity: O(1)
func (b *Bitset8) Add(n uint8) error {
	if n > 7 {
		return ErrNumberNotFit8
	}

	*b = *b | 1<<n

	return nil
}

// Contains returns true if the bitset contains the given number (if the bit in that position is 1). It returns ErrNumberNotFit8 error
// if the given number doesn't fit 8-bit bitset.
// time-complexity: O(1)
func (b *Bitset8) Contains(n uint8) (bool, error) {
	if n > 7 {
		return false, ErrNumberNotFit8
	}

	return *b&(1<<n) != 0, nil
}

// Delete removes the given number from the bitset (makes the bit in that position 0). It returns ErrNumberNotFit8 error
// if the given number doesn't fit 8-bit bitset.
// time-complexity: O(1)
func (b *Bitset8) Delete(n uint8) error {
	if n > 7 {
		return ErrNumberNotFit8
	}

	*b = ^(^*b | 1<<n)

	return nil
}

// Intersection gets another 8-bit bitset and returns a bitset containing the intersection of both.
func (b *Bitset8) Intersection(b2 Bitset8) Bitset8 {
	return *b & b2
}

// Union gets another 8-bit bitset and returns a bitset containing the union of both.
func (b *Bitset8) Union(b2 Bitset8) Bitset8 {
	return *b | b2
}

// SymmetricDiff gets another 8-bit bitset and returns a bitset containing the symmetric difference of both.
func (b *Bitset8) SymmetricDiff(b2 Bitset8) Bitset8 {
	return *b ^ b2
}

// Diff gets another 8-bit bitset and returns a bitset containing the difference of both.
func (b *Bitset8) Diff(b2 Bitset8) Bitset8 {
	return *b &^ b2
}

// IncAll increments elements one unit.
func (b *Bitset8) IncAll() {
	*b = *b << 1
}

// DecAll decrements elements one unit.
func (b *Bitset8) DecAll() {
	*b = *b >> 1
}
