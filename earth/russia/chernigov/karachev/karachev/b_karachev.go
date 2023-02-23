package karachev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉切夫KarachevBarony struct {
	feud.BaseBarony
}

var BKarachev卡拉切夫 feud.Barony = &卡拉切夫KarachevBarony{}

func init() {
	f := BKarachev卡拉切夫.(*卡拉切夫KarachevBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karachev",
		TitleName: "卡拉切夫",
		TitleCode: "b_karachev",
	}
}
