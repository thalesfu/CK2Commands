package kara_bogaz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉贾里KaradzhariBarony struct {
	feud.BaseBarony
}

var BKaradzhari卡拉贾里 feud.Barony = &卡拉贾里KaradzhariBarony{}

func init() {
    f := BKaradzhari卡拉贾里.(*卡拉贾里KaradzhariBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karadzhari",
		TitleName: "卡拉贾里",
		TitleCode: "b_karadzhari",
	}
}
