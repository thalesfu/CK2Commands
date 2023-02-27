package syrj

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 辛多尔SindorBarony struct {
	feud.BaseBarony
}

var BSindor辛多尔 feud.Barony = &辛多尔SindorBarony{}

func init() {
    f := BSindor辛多尔.(*辛多尔SindorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sindor",
		TitleName: "辛多尔",
		TitleCode: "b_sindor",
	}
}
