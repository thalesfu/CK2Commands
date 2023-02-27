package gottingen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡塞尔KasselBarony struct {
	feud.BaseBarony
}

var BKassel卡塞尔 feud.Barony = &卡塞尔KasselBarony{}

func init() {
    f := BKassel卡塞尔.(*卡塞尔KasselBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kassel",
		TitleName: "卡塞尔",
		TitleCode: "b_kassel",
	}
}
