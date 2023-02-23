package mema

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡博拉KaboraBarony struct {
	feud.BaseBarony
}

var BKabora卡博拉 feud.Barony = &卡博拉KaboraBarony{}

func init() {
	f := BKabora卡博拉.(*卡博拉KaboraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kabora",
		TitleName: "卡博拉",
		TitleCode: "b_kabora",
	}
}
