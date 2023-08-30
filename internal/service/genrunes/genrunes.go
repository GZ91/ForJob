package genrunes

import (
	"github.com/GZ91/linkreduct/internal/service"
	"math/rand"
	"time"
)

type Genrun struct {
	letterRunes []rune
	rander      *rand.Rand
}

func New() *Genrun {
	return &Genrun{
		letterRunes: []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"),
		rander:      rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (g Genrun) RandStringRunes(l int) string {
	var shortlink string
	for {
		b := make([]rune, l)
		for i := range b {
			b[i] = g.letterRunes[g.rander.Intn(len(g.letterRunes))]
		}
		shortlink = string(b)
		if !service.CheckURL(shortlink) {
			continue
		}
		break
	}
	return shortlink
}
