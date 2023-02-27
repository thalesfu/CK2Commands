package munda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波豆PathoBarony struct {
	feud.BaseBarony
}

var BPatho波豆 feud.Barony = &波豆PathoBarony{}

func init() {
    f := BPatho波豆.(*波豆PathoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "patho",
		TitleName: "波豆",
		TitleCode: "b_patho",
	}
}
