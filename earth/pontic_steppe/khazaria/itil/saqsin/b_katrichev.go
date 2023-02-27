package saqsin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡特里切夫KatrichevBarony struct {
	feud.BaseBarony
}

var BKatrichev卡特里切夫 feud.Barony = &卡特里切夫KatrichevBarony{}

func init() {
    f := BKatrichev卡特里切夫.(*卡特里切夫KatrichevBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "katrichev",
		TitleName: "卡特里切夫",
		TitleCode: "b_katrichev",
	}
}
