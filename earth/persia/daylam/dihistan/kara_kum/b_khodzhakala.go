package kara_kum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍贾卡拉KhodzhakalaBarony struct {
	feud.BaseBarony
}

var BKhodzhakala霍贾卡拉 feud.Barony = &霍贾卡拉KhodzhakalaBarony{}

func init() {
    f := BKhodzhakala霍贾卡拉.(*霍贾卡拉KhodzhakalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khodzhakala",
		TitleName: "霍贾卡拉",
		TitleCode: "b_khodzhakala",
	}
}
