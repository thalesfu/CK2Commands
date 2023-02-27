package murom

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡西莫夫KasimovBarony struct {
	feud.BaseBarony
}

var BKasimov卡西莫夫 feud.Barony = &卡西莫夫KasimovBarony{}

func init() {
    f := BKasimov卡西莫夫.(*卡西莫夫KasimovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kasimov",
		TitleName: "卡西莫夫",
		TitleCode: "b_kasimov",
	}
}
