package sieradzka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙代克SzadekBarony struct {
	feud.BaseBarony
}

var BSzadek沙代克 feud.Barony = &沙代克SzadekBarony{}

func init() {
    f := BSzadek沙代克.(*沙代克SzadekBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "szadek",
		TitleName: "沙代克",
		TitleCode: "b_szadek",
	}
}
