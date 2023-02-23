package pecs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希克洛什SiklosBarony struct {
	feud.BaseBarony
}

var BSiklos希克洛什 feud.Barony = &希克洛什SiklosBarony{}

func init() {
	f := BSiklos希克洛什.(*希克洛什SiklosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "siklos",
		TitleName: "希克洛什",
		TitleCode: "b_siklos",
	}
}
