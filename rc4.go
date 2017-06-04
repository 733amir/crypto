package crypto

// RC4 is a struct that hold information about status of RC4 cipher machine.
type RC4 struct {
	frame [256]byte
}

// Init will take 1 to 256 bytes and initialize the structure for RC4 pseudorandom stream.
func (r *RC4) Init(key []byte) {
	// Set frame and initialKey.
	for i := 0; i < 256; i++ {
		r.frame[i] = byte(i)
	}

	// Generate initial frame permutation.
	for i, j := 0, 0; i < 256; i++ {
		j = (j + int(r.frame[i]) + int(key[i%len(key)])) % 256
		r.frame[i], r.frame[j] = r.frame[j], r.frame[i]
	}
}

// KeyStream return a function that by each call to that function it returns a byte of key stream.
func (r *RC4) KeyStream() func() byte {
	i, j := 0, 0
	return func() byte {
		i = (i + 1) % 256
		j = (j + int(r.frame[i])) % 256
		r.frame[i], r.frame[j] = r.frame[j], r.frame[i]
		key := (int(r.frame[i]) + int(r.frame[j])) % 256
		return byte(r.frame[key])
	}
}
