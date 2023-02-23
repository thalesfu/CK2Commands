package badakhshan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡马尔KamarBarony struct {
	feud.BaseBarony
}

var BKamar卡马尔 feud.Barony = &卡马尔KamarBarony{}

func init() {
	f := BKamar卡马尔.(*卡马尔KamarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kamar",
		TitleName: "卡马尔",
		TitleCode: "b_kamar",
	}
}
