package roslavl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 坎科尔KankolBarony struct {
	feud.BaseBarony
}

var BKankol坎科尔 feud.Barony = &坎科尔KankolBarony{}

func init() {
    f := BKankol坎科尔.(*坎科尔KankolBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kankol",
		TitleName: "坎科尔",
		TitleCode: "b_kankol",
	}
}
