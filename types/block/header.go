package block

import (
	"io"
	"time"

	"github.com/pactus-project/pactus/crypto"
	"github.com/pactus-project/pactus/crypto/hash"
	"github.com/pactus-project/pactus/sortition"
	"github.com/pactus-project/pactus/util/encoding"
	"github.com/pactus-project/pactus/util/errors"
)

type Header struct {
	data headerData
}

type headerData struct {
	Version         uint8
	UnixTime        uint32
	PrevBlockHash   hash.Hash
	StateRoot       hash.Hash
	SortitionSeed   sortition.VerifiableSeed
	ProposerAddress crypto.Address
}

// Version returns the block version
func (h *Header) Version() uint8 {
	return h.data.Version
}

// Time returns the block time
func (h *Header) Time() time.Time {
	return time.Unix(int64(h.data.UnixTime), 0)
}

// UnixTime returns the block time in Unix value
func (h *Header) UnixTime() uint32 {
	return h.data.UnixTime
}

// StateRoot returns the state root hash
func (h *Header) StateRoot() hash.Hash {
	return h.data.StateRoot
}

// PrevBlockHash returns the previous block hash
func (h *Header) PrevBlockHash() hash.Hash {
	return h.data.PrevBlockHash
}

// SortitionSeed returns the sortition seed
func (h *Header) SortitionSeed() sortition.VerifiableSeed {
	return h.data.SortitionSeed
}

// ProposerAddress returns the proposer address
func (h *Header) ProposerAddress() crypto.Address {
	return h.data.ProposerAddress
}

func NewHeader(version uint8, time time.Time, stateRoot, prevBlockHash hash.Hash,
	sortitionSeed sortition.VerifiableSeed, proposerAddress crypto.Address) *Header {
	return &Header{
		data: headerData{
			Version:         version,
			UnixTime:        uint32(time.Unix()),
			PrevBlockHash:   prevBlockHash,
			StateRoot:       stateRoot,
			ProposerAddress: proposerAddress,
			SortitionSeed:   sortitionSeed,
		},
	}
}

func (h *Header) SanityCheck() error {
	if err := h.data.StateRoot.SanityCheck(); err != nil {
		return errors.Errorf(errors.ErrInvalidBlock, "invalid state root")
	}
	if err := h.data.ProposerAddress.SanityCheck(); err != nil {
		return errors.Errorf(errors.ErrInvalidBlock, "invalid proposer address")
	}

	return nil
}

// SerializeSize returns the number of bytes it would take to serialize the header.
func (h *Header) SerializeSize() int {
	return 138 // 5 + (2 * 32) + 48 + 21
}

// Encode encodes the receiver to w.
func (h *Header) Encode(w io.Writer) error {
	return encoding.WriteElements(w,
		h.data.Version,
		h.data.UnixTime,
		h.data.PrevBlockHash,
		h.data.StateRoot,
		h.data.SortitionSeed,
		h.data.ProposerAddress)
}

func (h *Header) Decode(r io.Reader) error {
	return encoding.ReadElements(r,
		&h.data.Version,
		&h.data.UnixTime,
		&h.data.PrevBlockHash,
		&h.data.StateRoot,
		&h.data.SortitionSeed,
		&h.data.ProposerAddress)
}
