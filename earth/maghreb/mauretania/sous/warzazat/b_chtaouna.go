package warzazat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 什塔乌纳ChtaounaBarony struct {
	feud.BaseBarony
}

var BChtaouna什塔乌纳 feud.Barony = &什塔乌纳ChtaounaBarony{}

func init() {
    f := BChtaouna什塔乌纳.(*什塔乌纳ChtaounaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chtaouna",
		TitleName: "什塔乌纳",
		TitleCode: "b_chtaouna",
	}
}
