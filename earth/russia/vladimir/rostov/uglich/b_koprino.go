package uglich

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科普里诺KoprinoBarony struct {
	feud.BaseBarony
}

var BKoprino科普里诺 feud.Barony = &科普里诺KoprinoBarony{}

func init() {
	f := BKoprino科普里诺.(*科普里诺KoprinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "koprino",
		TitleName: "科普里诺",
		TitleCode: "b_koprino",
	}
}
