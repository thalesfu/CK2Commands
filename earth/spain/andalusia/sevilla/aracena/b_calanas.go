package aracena

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉尼亚斯CalanasBarony struct {
	feud.BaseBarony
}

var BCalanas卡拉尼亚斯 feud.Barony = &卡拉尼亚斯CalanasBarony{}

func init() {
	f := BCalanas卡拉尼亚斯.(*卡拉尼亚斯CalanasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "calanas",
		TitleName: "卡拉尼亚斯",
		TitleCode: "b_calanas",
	}
}
