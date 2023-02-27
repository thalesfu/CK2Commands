package tura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基什克涅科尔KishkenekolBarony struct {
	feud.BaseBarony
}

var BKishkenekol基什克涅科尔 feud.Barony = &基什克涅科尔KishkenekolBarony{}

func init() {
    f := BKishkenekol基什克涅科尔.(*基什克涅科尔KishkenekolBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kishkenekol",
		TitleName: "基什克涅科尔",
		TitleCode: "b_kishkenekol",
	}
}
