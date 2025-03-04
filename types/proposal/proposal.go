package proposal

import (
	"fmt"

	"github.com/fxamacker/cbor/v2"
	"github.com/pactus-project/pactus/crypto"
	"github.com/pactus-project/pactus/crypto/bls"
	"github.com/pactus-project/pactus/crypto/hash"
	"github.com/pactus-project/pactus/types/block"
	"github.com/pactus-project/pactus/util"
	"github.com/pactus-project/pactus/util/errors"
)

type Proposal struct {
	data proposalData
}
type proposalData struct {
	Height    uint32         `cbor:"1,keyasint"`
	Round     int16          `cbor:"2,keyasint"`
	Block     *block.Block   `cbor:"3,keyasint"`
	Signature *bls.Signature `cbor:"4,keyasint"`
}

func NewProposal(height uint32, round int16, block *block.Block) *Proposal {
	return &Proposal{
		data: proposalData{
			Height: height,
			Round:  round,
			Block:  block,
		},
	}
}
func (p *Proposal) Height() uint32              { return p.data.Height }
func (p *Proposal) Round() int16                { return p.data.Round }
func (p *Proposal) Block() *block.Block         { return p.data.Block }
func (p *Proposal) Signature() crypto.Signature { return p.data.Signature }

func (p *Proposal) SanityCheck() error {
	if p.data.Block == nil {
		return errors.Errorf(errors.ErrInvalidSignature, "no block")
	}
	if p.data.Signature == nil {
		return errors.Errorf(errors.ErrInvalidSignature, "no signature")
	}
	if err := p.data.Block.SanityCheck(); err != nil {
		return err
	}
	if p.data.Height <= 0 {
		return errors.Error(errors.ErrInvalidHeight)
	}
	if p.data.Round < 0 {
		return errors.Error(errors.ErrInvalidRound)
	}
	return nil
}

func (p *Proposal) SetSignature(sig crypto.Signature) {
	p.data.Signature = sig.(*bls.Signature)
}

// SetPublicKey is doing nothing and just satisfies SignableMsg interface.
func (p *Proposal) SetPublicKey(crypto.PublicKey) {}

func (p *Proposal) SignBytes() []byte {
	sb := p.Block().Hash().Bytes()
	sb = append(sb, util.Uint32ToSlice(p.Height())...)
	sb = append(sb, util.Int16ToSlice(p.Round())...)

	return sb
}

func (p *Proposal) MarshalCBOR() ([]byte, error) {
	return cbor.Marshal(p.data)
}

func (p *Proposal) UnmarshalCBOR(bs []byte) error {
	return cbor.Unmarshal(bs, &p.data)
}

func (p *Proposal) Verify(pubKey crypto.PublicKey) error {
	if p.data.Signature == nil {
		return errors.Errorf(errors.ErrInvalidProposal, "no signature")
	}
	if err := pubKey.VerifyAddress(p.data.Block.Header().ProposerAddress()); err != nil {
		return err
	}
	return pubKey.Verify(p.SignBytes(), p.data.Signature)
}
func (p *Proposal) Hash() hash.Hash {
	return hash.CalcHash(p.SignBytes())
}

func (p *Proposal) IsForBlock(hash hash.Hash) bool {
	return p.Block().Hash().EqualsTo(hash)
}

func (p Proposal) Fingerprint() string {
	b := p.Block()
	return fmt.Sprintf("{%v/%v 🗃 %v}", p.data.Height, p.data.Round, b.Fingerprint())
}
