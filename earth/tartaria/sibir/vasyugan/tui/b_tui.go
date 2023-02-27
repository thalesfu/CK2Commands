package tui

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图伊TuiBarony struct {
	feud.BaseBarony
}

var BTui图伊 feud.Barony = &图伊TuiBarony{}

func init() {
    f := BTui图伊.(*图伊TuiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tui",
		TitleName: "图伊",
		TitleCode: "b_tui",
	}
}
