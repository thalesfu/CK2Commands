package saintois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥尔努瓦OrnoisBarony struct {
	feud.BaseBarony
}

var BOrnois奥尔努瓦 feud.Barony = &奥尔努瓦OrnoisBarony{}

func init() {
    f := BOrnois奥尔努瓦.(*奥尔努瓦OrnoisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ornois",
		TitleName: "奥尔努瓦",
		TitleCode: "b_ornois",
	}
}
