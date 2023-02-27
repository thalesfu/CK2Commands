package bhumilka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 门罗尔MangrolBarony struct {
	feud.BaseBarony
}

var BMangrol门罗尔 feud.Barony = &门罗尔MangrolBarony{}

func init() {
    f := BMangrol门罗尔.(*门罗尔MangrolBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mangrol",
		TitleName: "门罗尔",
		TitleCode: "b_mangrol",
	}
}
