package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

func PrivateKey(p *big.Int) *big.Int {
	l := big.NewInt(1)
	l.Sub(p, big.NewInt(2))
	n, err := rand.Int(rand.Reader, l)
	if err != nil {
		panic(err)
	}
	n.Add(n, big.NewInt(2))
	return n
}

func PublicKey(a *big.Int, p *big.Int, g int64) *big.Int {
	return big.NewInt(0).Exp(big.NewInt(g), a, p)
}

func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	priv := PrivateKey(p)
	pub := PublicKey(priv, p, g)
	return priv, pub
}

func SecretKey(a *big.Int, b *big.Int, p *big.Int) *big.Int {
	return big.NewInt(1).Exp(b, a, p)
}
