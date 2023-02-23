package perigord

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比龙BironBarony struct {
	feud.BaseBarony
}

var BBiron比龙 feud.Barony = &比龙BironBarony{}

func init() {
	f := BBiron比龙.(*比龙BironBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "biron",
		TitleName: "比龙",
		TitleCode: "b_biron",
	}
}
