package temes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 恰克CsakBarony struct {
	feud.BaseBarony
}

var BCsak恰克 feud.Barony = &恰克CsakBarony{}

func init() {
	f := BCsak恰克.(*恰克CsakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "csak",
		TitleName: "恰克",
		TitleCode: "b_csak",
	}
}
