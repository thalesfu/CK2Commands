package calatrava

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉特拉瓦CalatravaBarony struct {
	feud.BaseBarony
}

var BCalatrava卡拉特拉瓦 feud.Barony = &卡拉特拉瓦CalatravaBarony{}

func init() {
	f := BCalatrava卡拉特拉瓦.(*卡拉特拉瓦CalatravaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "calatrava",
		TitleName: "卡拉特拉瓦",
		TitleCode: "b_calatrava",
	}
}
