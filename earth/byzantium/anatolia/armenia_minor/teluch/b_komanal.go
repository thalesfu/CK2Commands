package teluch

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科曼阿拉KomanalBarony struct {
	feud.BaseBarony
}

var BKomanal科曼阿拉 feud.Barony = &科曼阿拉KomanalBarony{}

func init() {
    f := BKomanal科曼阿拉.(*科曼阿拉KomanalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "komanal",
		TitleName: "科曼阿拉",
		TitleCode: "b_komanal",
	}
}
