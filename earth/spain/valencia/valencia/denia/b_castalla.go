package denia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡斯塔利亚CastallaBarony struct {
	feud.BaseBarony
}

var BCastalla卡斯塔利亚 feud.Barony = &卡斯塔利亚CastallaBarony{}

func init() {
    f := BCastalla卡斯塔利亚.(*卡斯塔利亚CastallaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "castalla",
		TitleName: "卡斯塔利亚",
		TitleCode: "b_castalla",
	}
}
