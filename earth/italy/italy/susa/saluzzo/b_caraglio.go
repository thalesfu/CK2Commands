package saluzzo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉廖CaraglioBarony struct {
	feud.BaseBarony
}

var BCaraglio卡拉廖 feud.Barony = &卡拉廖CaraglioBarony{}

func init() {
	f := BCaraglio卡拉廖.(*卡拉廖CaraglioBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "caraglio",
		TitleName: "卡拉廖",
		TitleCode: "b_caraglio",
	}
}
