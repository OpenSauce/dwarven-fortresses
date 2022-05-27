package worldgen

import (
	"math/rand"
	"time"

	"github.com/ojrac/opensimplex-go"
)

type WorldGen struct {
	sampler opensimplex.Noise
}

func New() WorldGen {
	rand.Seed(time.Now().UTC().UnixNano())
	return WorldGen{
		sampler: opensimplex.New(rand.Int63()),
	}
}

func GenerateWorld() {

}

func (w *WorldGen) GenerateTile(x, y, z, octaves int, scale, persistance, lacunarity float64) float64 {
	amplitude := 1.0
	frequency := 1.0
	noiseHeight := 1.0

	for i := 0; i < octaves; i++ {
		sampleX := float64(x) / scale * frequency
		sampleY := float64(y) / scale * frequency
		sampleZ := float64(z) / scale * frequency

		perlinValue := w.sampler.Eval3(sampleX, sampleY, sampleZ)
		noiseHeight += perlinValue * amplitude

		amplitude *= persistance
		frequency *= lacunarity
	}

	return noiseHeight
}
