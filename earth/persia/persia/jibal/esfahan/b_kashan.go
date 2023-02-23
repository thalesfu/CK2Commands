package esfahan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尚KashanBarony struct {
	feud.BaseBarony
}

var BKashan卡尚 feud.Barony = &卡尚KashanBarony{}

func init() {
	f := BKashan卡尚.(*卡尚KashanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kashan",
		TitleName: "卡尚",
		TitleCode: "b_kashan",
	}
}
