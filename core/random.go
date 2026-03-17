package core

import (
	"crypto/md5"
	"encoding/hex"
	"math"
	"math/big"
	"math/rand/v2"
	"strconv"
)

type Random struct {
	rng_pcg *rand.PCG
	rng     *rand.Rand
	seed    string
}

func NewRandom() *Random {
	seed := strconv.Itoa(rand.Int())
	rng := Random{}
	rng.Seed(seed)

	return &rng
}

func (r *Random) Seed(seed string) {
	bi := big.NewInt(0)
	h := md5.New()
	h.Write([]byte(seed))
	hexstr := hex.EncodeToString(h.Sum(nil))
	bi.SetString(hexstr, 16)

	r.rng_pcg = rand.NewPCG(uint64(bi.Int64()), 1024)
	r.rng = rand.New(r.rng_pcg)
	r.seed = seed
}

func (r *Random) LoadState(seed string, state []byte) {
	r.Seed(seed)
	r.rng_pcg.UnmarshalBinary(state)
}

func (r *Random) SaveState() (string, []byte, error) {
	st, err := r.rng_pcg.MarshalBinary()
	return r.seed, st, err
}

func (r *Random) Float32() float32 {
	return r.rng.Float32()
}

func (r *Random) Float64() float64 {
	return r.rng.Float64()
}

func (r *Random) NormFloat64(stdDev float64, mean float64) float64 {
	return r.rng.NormFloat64()*stdDev + mean
}

func (r *Random) Int() int {
	return r.rng.Int()
}

func (r *Random) IntN(min int, max int) int {
	return r.rng.IntN(max-min) + min
}

func (r *Random) Int32() int32 {
	return r.rng.Int32()
}

func (r *Random) Int32N(min int32, max int32) int32 {
	return r.rng.Int32N(max-min) + min
}

func (r *Random) Int64() int64 {
	return r.rng.Int64()
}

func (r *Random) Int64N(min int64, max int64) int64 {
	return r.rng.Int64N(max-min) + min
}

func (r *Random) Perm(n int) []int {
	return r.rng.Perm(n)
}

func (r *Random) Shuffle(n int, swap func(i, j int)) {
	r.rng.Shuffle(n, swap)
}

func (r *Random) Chances(chances []float64) int {
	sum := 0.0
	for _, v := range chances {
		sum += math.Max(0, v)
	}
	if sum <= 0 {
		return -1
	}
	value := sum
	sum = 0
	for i, v := range chances {
		sum += math.Max(0, v)
		if value < sum {
			return i
		}
	}

	return -1
}
