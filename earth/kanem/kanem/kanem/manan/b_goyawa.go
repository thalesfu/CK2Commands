package manan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 戈亚瓦GoyawaBarony struct {
	feud.BaseBarony
}

var BGoyawa戈亚瓦 feud.Barony = &戈亚瓦GoyawaBarony{}

func init() {
	f := BGoyawa戈亚瓦.(*戈亚瓦GoyawaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "goyawa",
		TitleName: "戈亚瓦",
		TitleCode: "b_goyawa",
	}
}
