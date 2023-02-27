package kara_bogaz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉博加兹Kara_bogazBarony struct {
	feud.BaseBarony
}

var BKara_bogaz卡拉博加兹 feud.Barony = &卡拉博加兹Kara_bogazBarony{}

func init() {
    f := BKara_bogaz卡拉博加兹.(*卡拉博加兹Kara_bogazBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kara_bogaz",
		TitleName: "卡拉博加兹",
		TitleCode: "b_kara_bogaz",
	}
}
