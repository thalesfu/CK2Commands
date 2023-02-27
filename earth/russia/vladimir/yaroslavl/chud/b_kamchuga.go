package chud

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡姆丘加KamchugaBarony struct {
	feud.BaseBarony
}

var BKamchuga卡姆丘加 feud.Barony = &卡姆丘加KamchugaBarony{}

func init() {
    f := BKamchuga卡姆丘加.(*卡姆丘加KamchugaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kamchuga",
		TitleName: "卡姆丘加",
		TitleCode: "b_kamchuga",
	}
}
