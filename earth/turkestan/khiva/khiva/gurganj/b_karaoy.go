package gurganj

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉奥伊KaraoyBarony struct {
	feud.BaseBarony
}

var BKaraoy卡拉奥伊 feud.Barony = &卡拉奥伊KaraoyBarony{}

func init() {
    f := BKaraoy卡拉奥伊.(*卡拉奥伊KaraoyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karaoy",
		TitleName: "卡拉奥伊",
		TitleCode: "b_karaoy",
	}
}
