package northumberland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫珀斯MorpethBarony struct {
	feud.BaseBarony
}

var BMorpeth莫珀斯 feud.Barony = &莫珀斯MorpethBarony{}

func init() {
    f := BMorpeth莫珀斯.(*莫珀斯MorpethBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "morpeth",
		TitleName: "莫珀斯",
		TitleCode: "b_morpeth",
	}
}
